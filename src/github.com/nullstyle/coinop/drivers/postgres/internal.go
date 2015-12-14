package postgres

import (
	"github.com/lib/pq"
)

// isUniqueErr returns true if the error represents a postgres
// "unique_violation" error (code 23505)
func isUniqueErr(err error) bool {
	if err, ok := err.(*pq.Error); ok {
		return err.Code == "23505"
	}

	return false
}
