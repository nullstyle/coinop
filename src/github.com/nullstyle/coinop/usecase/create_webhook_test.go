package usecase_test

import (
	"errors"

	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/fake"
	"github.com/nullstyle/coinop/test"
	. "github.com/nullstyle/coinop/usecase"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("CreateWebhook", func() {
	var (
		subject CreateWebhook
		input   entity.Webhook
		repo    *MockWebhookRepository
		output  RepoID
		err     error
	)

	BeforeEach(func() {
		repo = &MockWebhookRepository{}
		input = fake.WebhookEntity()
	})

	AfterEach(func() {
		test.VerifyMock(repo.Mock)
	})

	JustBeforeEach(func() {
		subject = CreateWebhook{repo}
		output, err = subject.Exec(input)
	})

	Context("a working repo", func() {
		BeforeEach(func() {
			repo.On("SaveWebhook", mock.Anything).Return(nil).Run(func(
				args mock.Arguments,
			) {
				hook := args.Get(0).(*entity.Webhook)
				hook.ID = &RepoID{T: "webhook", V: 1}
			})
		})

		It("succeeds", func() {
			Expect(err).To(BeNil())
		})
	})

	Context("an erroring repo", func() {
		BeforeEach(func() {
			repo.On("SaveWebhook", mock.Anything).Return(errors.New("error"))
		})

		It("errors with a *CreateWebhookError", func() {
			Expect(err).
				To(BeAssignableToTypeOf(&CreateWebhookError{}))
			Expect(err).
				To(MatchError("failed to create webhook (repo): error"))
		})
	})

	Context("a repo that silently does not save the user", func() {
		BeforeEach(func() {
			repo.On("SaveWebhook", mock.Anything).Return(nil)
		})

		It("errors", func() {
			Expect(err).To(MatchError("failed to create webhook (repo)"))
		})
	})
})
