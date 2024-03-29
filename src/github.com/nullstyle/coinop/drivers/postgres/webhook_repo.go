package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/usecase"
)

// DestroyWebhook deletes a webhook from the postgres db
func (db *Driver) DestroyWebhook(ID usecase.RepoID) error {
	if ID.T != "webhook" {
		return errors.New("bad id")
	}

	_, err := db.DB.Exec(Queries.Webhook.Delete, ID.V)
	return err
}

// ForDestination reads all webhooks that could be triggered by `dest`
func (db *Driver) ForDestination(
	dest entity.AccountID,
) ([]entity.Webhook, error) {
	return db.selectWebhooks(Queries.Webhook.ForDestination, string(dest))
}

// ListWebhooks reads all webhooks from the database
func (db *Driver) ListWebhooks() ([]entity.Webhook, error) {
	return db.selectWebhooks(Queries.Webhook.All)
}

// SaveWebhook writes the webhook to the postgres database, updating an existing
// row if it exists.
func (db *Driver) SaveWebhook(hook *entity.Webhook) error {
	if hook.ID == nil || hook.IsNew() {
		return db.insertWebhook(hook)
	}

	panic("TODO: update existing webhook")
}

func (db *Driver) insertWebhook(hook *entity.Webhook) error {
	var (
		id  int64
		mf  sql.NullString
		mtf sql.NullString
	)

	if hook.MemoFilter != nil {
		mtf = sql.NullString{String: hook.MemoFilter.Type, Valid: true}
		mf = sql.NullString{String: hook.MemoFilter.Value, Valid: true}
	}

	err := db.DB.Get(&id, Queries.Webhook.Insert,
		hook.URL.String(),
		string(hook.DestinationFilter),
		mtf,
		mf,
		time.Now().UTC(),
	)
	if err != nil {
		return err
	}

	hook.ID = &usecase.RepoID{T: "webhook", V: id}
	return nil
}

func (db *Driver) selectWebhooks(
	sql string,
	args ...interface{},
) (result []entity.Webhook, err error) {
	var rows []Webhook
	err = db.DB.Select(&rows, sql, args...)
	if err != nil {
		return
	}
	result = make([]entity.Webhook, len(rows))
	for i, row := range rows {
		result[i], err = row.Entity()
		if err != nil {
			err = errors.New("corrupt webhook row: " + err.Error())
			return
		}
	}
	return
}

var _ usecase.WebhookRepository = &Driver{}
