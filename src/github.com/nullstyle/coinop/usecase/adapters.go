package usecase

import (
	"fmt"
	"github.com/nullstyle/coinop/entity"
)

// RepoID represents an identity stored within an implementation of this packages repository port.  It implements `entity.ID`.
type RepoID struct {
	T string
	V int64
}

// IsNew returns true if the identity has not been persisted.
func (id *RepoID) IsNew() bool {
	return id.V == 0
}

// Type returns the type of the identity
func (id *RepoID) Type() string {
	return id.T
}

// ID returns a string representation of the identity suitable for comparison
func (id *RepoID) ID() string {
	return fmt.Sprintf("%s/%d", id.T, id.V)
}

func (id *RepoID) ShortID() string {
	return fmt.Sprintf("%d", id.V)
}

var _ entity.ID = &RepoID{}
