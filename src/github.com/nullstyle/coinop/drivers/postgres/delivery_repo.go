package postgres

import (
	// "database/sql"
	// "errors"
	"time"

	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/usecase"
)

// LoadCursor loads the last saved payment cursor.
func (db *Driver) LoadCursor() (string, error) {
	//TODO
	return "now", nil
}

// SaveDeliveries saves a batch of deliveries into the db.
func (db *Driver) SaveDeliveries(cursor string, d []entity.Delivery) error {
	//TODO
	return nil
}

// StartDelivery indicates to the repository that the caller wants to perform
// the provided delivery.  An implementation should ensure that the token
// returned will authorize an update to the delivery during a window of
// exclusive access to the delivery bounded by the returned deadline.
func (db *Driver) StartDelivery(entity.Delivery) (
	token int64,
	deadline time.Time,
	err error,
) {
	return int64(0), time.Now(), nil
}

// MarkDeliverySuccess marks the provided delivery as successfully completed.
func (db *Driver) MarkDeliverySuccess(
	token int64,
	delivery entity.Delivery,
) error {
	return nil
}

// MarkDeliveryFailed marks the provided delivery as a failure, provided it
// has not been marked as successful or failed in the past.
func (db *Driver) MarkDeliveryFailed(
	token int64,
	delivery entity.Delivery,
) error {
	return nil
}

// FailedDeliveries returns a slice of deliveries that have been marked as
// failed.
func (db *Driver) FailedDeliveries() ([]entity.Delivery, error) {
	return nil, nil
}

var _ usecase.DeliveryRepository = &Driver{}
