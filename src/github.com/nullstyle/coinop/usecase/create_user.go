package usecase

import (
	"fmt"
	"time"

	"github.com/nullstyle/coinop/entity"
)

// CreateUser Use Case
//
// A client should be able to create a new account, so that future payments can
// be recorded for the user.
//
// In: nothing
// Output: the new created account id
//
//
// Steps:
// 1. A coinop client requests a new user to be created
// 2. System forms new user account with default values
// 3. System records user account in repo
// 4. System delivers newly created account's id to client
//
type CreateUser struct {
	DB UserRepository
}

// Exec runs the use case.
func (kase *CreateUser) Exec() (uid int64, err error) {
	u := entity.User{
		CreatedAt: time.Now(),
	}

	err = kase.DB.CreateUser(&u)
	if err != nil {
		err = &CreateUserError{Step: "repo", Child: err}
		return
	}

	if u.IsNew() {
		// TODO: add some further explanation that the repo failed to assign
		// and ID.
		err = &CreateUserError{Step: "repo"}
		return
	}

	uid = u.ID
	return
}

// CreateUserError is emitted when
type CreateUserError struct {
	Step  string
	Child error
}

func (err *CreateUserError) Error() string {
	base := fmt.Sprintf("failed to create user (%s)", err.Step)

	if err.Child != nil {
		return fmt.Sprintf("%s: %s", base, err.Child)
	}

	return base
}
