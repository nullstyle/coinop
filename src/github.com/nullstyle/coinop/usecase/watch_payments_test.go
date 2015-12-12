package usecase_test

import (
	"errors"

	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/fake"
	. "github.com/nullstyle/coinop/usecase"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("WatchPayments", func() {
	var (
		stream *MockPaymentProvider
		proc   *mockPaymentProcessor
		cursor string
		err    error
	)

	BeforeEach(func() {
		cursor = "now"
		stream = &MockPaymentProvider{}
		proc = &mockPaymentProcessor{}
	})

	JustBeforeEach(func() {
		subject := &WatchPayments{stream, proc}
		err = subject.Exec(cursor)
	})

	Context("no payments seen", func() {
		BeforeEach(func() {
			stream.On("StreamPayments", cursor, mock.Anything).Return(nil)
		})

		It("succeeds", func() {
			Expect(err).To(BeNil())
		})
	})

	Context("an error occurs prior to any emitted payments", func() {
		BeforeEach(func() {
			stream.
				On("StreamPayments", cursor, mock.Anything).
				Return(errors.New("boom"))
		})

		It("fails", func() {
			Expect(err).To(MatchError("boom"))
		})
	})

	Context("several payments are seen", func() {
		BeforeEach(func() {
			payments := []entity.Payment{fake.PaymentEntity(), fake.PaymentEntity()}

			stream.
				On("StreamPayments", cursor, mock.Anything).
				Return(nil).
				Run(func(
				args mock.Arguments,
			) {
				fn := args.Get(1).(PaymentHandler)
				fn(payments[0])
				fn(payments[1])
			})

			proc.On("Exec", payments[0]).Return(nil)
			proc.On("Exec", payments[1]).Return(nil)
		})

		It("succeeds", func() {
			Expect(err).To(BeNil())
		})
	})

})
