package usecase_test

import (
	"errors"
	"fmt"

	"github.com/nullstyle/coinop/entity"
	. "github.com/nullstyle/coinop/usecase"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateUser", func() {
	var (
		subject CreateUser
		repo    UserRepository
		output  int64
		err     error
	)

	JustBeforeEach(func() {
		subject = CreateUser{repo}
		output, err = subject.Exec()
	})

	Context("an empty repo", func() {
		BeforeEach(func() {
			repo = &mockUserRepository{}
		})

		It("outputs 1", func() {
			Expect(output).To(Equal(int64(1)))
			Expect(err).To(BeNil())
		})

		It("records the user on the repo", func() {
			repo := repo.(*mockUserRepository)
			Expect(repo.Users).To(HaveLen(1))
			Expect(repo.Users[0].ID).To(Equal(output))
		})
	})

	Context("a repo with existing users", func() {
		BeforeEach(func() {
			repo = &mockUserRepository{
				Users: []entity.User{
					{ID: 1},
					{ID: 2},
				},
			}
		})
		It("outputs  the next available id", func() {
			Expect(output).To(Equal(int64(3)))
			Expect(err).To(BeNil())
		})

		It("records the user on the repo", func() {
			repo := repo.(*mockUserRepository)
			Expect(repo.Users).To(HaveLen(3))
			Expect(repo.Users[2].ID).To(Equal(output))
		})
	})

	Context("an erroring repo", func() {
		BeforeEach(func() {
			repo = &mockUserRepository{
				Err: errors.New("cannot save user"),
			}
		})

		It("errors with a *CreateUserError", func() {
			fmt.Fprintf(GinkgoWriter, "here: %s", err.Error())
			Expect(err).
				To(BeAssignableToTypeOf(&CreateUserError{}))
			Expect(err).
				To(MatchError("failed to create user (repo): cannot save user"))
		})
	})

	Context("a repo that silently does not save the user", func() {
		BeforeEach(func() {
			repo = &mockUserRepository{NoSave: true}
		})

		It("errors", func() {
			Expect(err).To(MatchError("failed to create user (repo)"))
		})
	})
})
