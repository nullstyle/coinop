package entity

import (
	"time"
)

// Account represents a single named account.
type Account struct {
	ID   AccountID
	Name string

	CreatedAt time.Time
}

// AccountID represents a durable identifier for an account
type AccountID int64

// Client represents a single integrator with the coinop service, for example an
// API client, a gui client or the cli client.
type Client struct {
	ID string
}

// User represents a single user of a system.
type User struct {
	ID UserID

	CreatedAt time.Time
}

// IsNew returns if this user is considered "new" by the system or not. User's
// who have a zero value ID are considered "new".
func (u *User) IsNew() bool {
	return u.ID == 0
}

// UserID represents a durable identifier for an user
type UserID int64
