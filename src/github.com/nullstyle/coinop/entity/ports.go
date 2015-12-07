package entity

// ID represents a identity, and whether or not that identity
// has been persisted.
type ID interface {
	IsNew() bool
	Type() string
	ID() string
	ShortID() string
}
