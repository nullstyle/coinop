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

	JustBeforeEach(func() {
		subject = CreateWebhook{repo}
		output, err = subject.Exec(input)
	})

	Context("a working repo", func() {
		BeforeEach(func() {
			input = fake.WebhookEntity()
			repo = &mockWebhookRepository{}
		})

		It("outputs 1", func() {
			Expect(output.V).To(Equal(int64(1)))
			Expect(err).To(BeNil())
		})

		It("records the user on the repo", func() {
			repo := repo.(*mockWebhookRepository)
			Expect(repo.Items).To(HaveLen(1))
			id := repo.Items[0].ID.(*RepoID)
			Expect(*id).To(Equal(output))
		})

		Context("an invalid webhook", func() {
			BeforeEach(func() {
				input.DestinationFilter = ""
			})

			It("errors", func() {
				msg := "failed to create webhook (validate): " +
					"invalid webhook: bad destination filter"
				Expect(err).To(MatchError(msg))
			})
		})
	})

	Context("an erroring repo", func() {
		BeforeEach(func() {
			input = fake.WebhookEntity()
			repo = &mockWebhookRepository{
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
		BeforeEach(func() {
			input = fake.WebhookEntity()
			repo = &mockWebhookRepository{NoSave: true}
		})

		It("errors", func() {
			Expect(err).To(MatchError("failed to create webhook (repo)"))
		})
	})
})
