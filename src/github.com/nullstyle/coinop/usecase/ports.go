package usecase

import (
	"github.com/nullstyle/coinop/entity"
)

// UserRepository implementations can create new users, usually by recording a
// new row in a database or
type UserRepository interface {
	CreateUser(*entity.User) error
}
