package usecase_test

import (
	"github.com/nullstyle/coinop/entity"
	. "github.com/nullstyle/coinop/usecase"
)

type mockUserRepository struct {
	Users  []entity.User
	Err    error
	NoSave bool
}

func (repo *mockUserRepository) CreateUser(user *entity.User) error {
	if repo.NoSave {
		return nil
	}

	if repo.Err != nil {
		return repo.Err
	}

	user.ID = entity.UserID(len(repo.Users) + 1)
	repo.Users = append(repo.Users, *user)
	return nil
}

// ensure MockUserRepository conforms to UserRepository port
var _ UserRepository = &mockUserRepository{}
