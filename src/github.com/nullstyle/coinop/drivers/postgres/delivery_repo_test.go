package postgres_test

import (
	. "github.com/nullstyle/coinop/drivers/postgres"
	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/fake"
	// "github.com/nullstyle/coinop/usecase"
	. "github.com/nullstyle/coinop/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("*postgres.Driver as DeliveryRepository", func() {
	var subject Driver

	Describe("LoadCursor", func() {
		var (
			cursor string
			err    error
		)
		JustBeforeEach(func() {
			cursor, err = subject.LoadCursor()
		})

		Context("a working db", func() {
			BeforeEach(func() {
				subject = Driver{DB: db}
			})

			Context("no saved cursor", func() {
				It("returns 'now'", func() {
					Expect(cursor).To(Equal("now"))
				})

				ItSucceeds(&err)
			})

			Context("previously saved cursor", func() {
				BeforeEach(func() {
					err := subject.SaveKV("cursor", "123")
					Expect(err).To(BeNil())
				})

				It("returns the previously saved value", func() {
					Expect(cursor).To(Equal("123"))
				})

				ItSucceeds(&err)
			})

			Context("corrupted cursor", func() {
				BeforeEach(func() {
					err := subject.SaveKV("cursor", 213)
					Expect(err).To(BeNil())
				})

				It("fails", func() {
					Expect(err).To(MatchError(ContainSubstring("cannot unmarshal")))
				})
			})
		})

		Context("an missing db", func() {
			BeforeEach(func() {
				subject = Driver{DB: badDB}
			})

			It("fails", func() {
				Expect(err).To(MatchError(ContainSubstring("does not exist")))
			})
		})

	})

	Describe("SaveDeliveries", func() {
		var (
			cursor string
			ds     []entity.Delivery
			err    error
		)

		BeforeEach(func() {
			cursor = "123"
			ds = []entity.Delivery{
				fake.DeliveryEntity(),
				fake.DeliveryEntity(),
			}
			subject = Driver{DB: db}
		})

		JustBeforeEach(func() {
			err = subject.SaveDeliveries(cursor, ds)
		})

		It("saves the cursor", func() {
			saved, err := subject.LoadCursor()
			Expect(err).To(BeNil())
			Expect(saved).To(Equal(cursor))
		})

		It("saves both deliveries to the db", func() {
			var dc int
			err := subject.DB.Get(&dc, "SELECT COUNT(*) FROM coinop.deliveries")
			Expect(err).To(BeNil())
			Expect(dc).To(Equal(2))
		})

		It("populates the ids", func() {
			Expect(ds[0].ID).ToNot(BeNil())
			Expect(ds[1].ID).ToNot(BeNil())
		})

		ItSucceeds(&err)

	})
})
