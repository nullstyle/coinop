package postgres_test

import (
	. "github.com/nullstyle/coinop/drivers/postgres"
	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/fake"
	"github.com/nullstyle/coinop/usecase"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("*postgres.Driver as WebrookRepository", func() {

	Describe("SaveWebhook", func() {
		var (
			subject *Driver
			input   entity.Webhook
			err     error
		)
		JustBeforeEach(func() {
			input = fake.WebhookEntity()
			subject = &Driver{DB: db}
			err = subject.SaveWebhook(&input)
		})

		Context("a new webhook", func() {
			It("sets an id", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(input.ID).To(BeAssignableToTypeOf(&usecase.RepoID{}))
				id := input.ID.(*usecase.RepoID)
				Expect(id.T).To(Equal("webhook"))
				Expect(id.V).To(Equal(int64(1)))
			})

			PIt("inserts a row")
		})
		Context("an existing webhook", func() {
			PIt("updates the row")
		})

		Context("an erroring db", func() {
			PIt("errors")
		})

	})
})
