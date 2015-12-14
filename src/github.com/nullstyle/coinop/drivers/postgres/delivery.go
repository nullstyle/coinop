package postgres

import (
	"time"

	"github.com/nullstyle/coinop/entity"
)

// Scan populates `row` from `ent`.
func (row *Delivery) Scan(ent entity.Delivery) (err error) {
	row.URL = ent.URL.String()

	err = row.Payment.Scan(ent.Payment)
	if err != nil {
		return
	}

	if ent.StartedAt != (time.Time{}) {
		t := ent.StartedAt
		row.StartedAt = &t
	}

	if ent.LastFailedAt != (time.Time{}) {
		t := ent.LastFailedAt
		row.LastFailedAt = &t
	}

	if ent.SucceededAt != (time.Time{}) {
		t := ent.SucceededAt
		row.SucceededAt = &t
	}

	return
}

//ScanInsert scans `ent` into `row` and prepares `row` such that it represents
//a new, insertable row.
func (row *Delivery) ScanInsert(ent entity.Delivery) error {
	row.ID = 0
	row.CreatedAt = time.Now().UTC()
	row.UpdatedAt = &row.CreatedAt

	return row.Scan(ent)
}
