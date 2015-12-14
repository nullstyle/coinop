package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/usecase"
)

// LoadCursor loads the last saved payment cursor.
func (db *Driver) LoadCursor() (result string, err error) {
	err = db.GetKV("cursor", &result)

	if err == sql.ErrNoRows {
		result = "now"
		err = nil
	}
	return
}

// SaveDeliveries saves a batch of deliveries into the db.
func (db *Driver) SaveDeliveries(cursor string, d []entity.Delivery) error {
	tx, err := db.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = db.saveCursor(tx, cursor)
	if err != nil {
		return err
	}

	err = db.saveDeliveries(tx, d)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// StartDelivery indicates to the repository that the caller wants to perform
// the provided delivery.  An implementation should ensure that the token
// returned will authorize an update to the delivery during a window of
// exclusive access to the delivery bounded by the returned deadline.
func (db *Driver) StartDelivery(ent entity.Delivery) (
	token int64,
	deadline time.Time,
	err error,
) {
	id := ent.ID.(*usecase.RepoID).V
	start := time.Now()
	deadline = start.Add(1 * time.Minute)

	err = db.DB.Get(&token, Queries.Delivery.Start, start, id)

	if err == sql.ErrNoRows {
		//TODO: turn this into a well-known error of the usecase layer
		err = errors.New("could not acquire lock")
	}

	return
}

// MarkDeliverySuccess marks the provided delivery as successfully completed.
func (db *Driver) MarkDeliverySuccess(
	token int64,
	delivery entity.Delivery,
) error {

	var version int
	err := db.DB.Get(&version, Queries.Delivery.MarkSuccessful,
		time.Now().UTC(),
		delivery.ID.(*usecase.RepoID).V,
		token,
	)

	if err == sql.ErrNoRows {
		return errors.New("lock expired")
	}

	return err
}

// MarkDeliveryFailed marks the provided delivery as a failure, provided it
// has not been marked as successful or failed in the past.
func (db *Driver) MarkDeliveryFailed(
	token int64,
	delivery entity.Delivery,
) error {

	var version int
	err := db.DB.Get(&version, Queries.Delivery.MarkFailed,
		time.Now().UTC(),
		delivery.ID.(*usecase.RepoID).V,
		token,
	)

	if err == sql.ErrNoRows {
		return errors.New("lock expired")
	}

	return err
}

// FailedDeliveries returns a slice of deliveries that have been marked as
// failed.
func (db *Driver) FailedDeliveries() ([]entity.Delivery, error) {
	return nil, nil
}

var _ usecase.DeliveryRepository = &Driver{}

func (db *Driver) saveCursor(tx *sqlx.Tx, cursor string) error {
	return db.saveKV(tx, "cursor", cursor)
}

func (db *Driver) saveDeliveries(tx *sqlx.Tx, ds []entity.Delivery) error {
	for i := range ds {
		err := db.saveDelivery(tx, &ds[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Driver) saveDelivery(tx *sqlx.Tx, d *entity.Delivery) error {
	var row Delivery
	err := row.ScanInsert(*d)
	if err != nil {
		return err
	}

	sql, args, err := db.DB.BindNamed(Queries.Delivery.Insert, &row)
	if err != nil {
		return err
	}

	err = db.DB.Get(&row, sql, args...)
	if err != nil {
		return err
	}

	d.ID = &usecase.RepoID{T: "delivery", V: row.ID}
	return nil
}
