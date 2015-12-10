package usecase_test

import (
	"errors"

	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/fake"
	. "github.com/nullstyle/coinop/usecase"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateWebhook", func() {
	var (
		subject CreateWebhook
		input   entity.Webhook
		repo    WebhookRepository
		output  RepoID
		err     error
	)

	BeforeEach(func() {
		repo = &MockWebhookRepository{}
		input = fake.WebhookEntity()
	})

	JustBeforeEach(func() {
		subject = CreateWebhook{repo}
		output, err = subject.Exec(input)
	})

	Context("a working repo", func() {
		It("succeeds", func() {
			Expect(err).To(BeNil())
		})
	})

	Context("an erroring repo", func() {
		BeforeEach(func() {
			repo = &MockWebhookRepository{
				Err: errors.New("error"),
			}
		})

		It("errors with a *CreateWebhookError", func() {
			Expect(err).
				To(BeAssignableToTypeOf(&CreateWebhookError{}))
			Expect(err).
				To(MatchError("failed to create webhook (repo): error"))
		})
	})

	Context("a repo that silently does not save the user", func() {
		It("errors", func() {
			Expect(err).To(MatchError("failed to create webhook (repo)"))
		})
	})
})
