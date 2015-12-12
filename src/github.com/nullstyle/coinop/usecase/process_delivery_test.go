package usecase_test

import (
	"time"
	// "errors"
	//
	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/fake"
	"github.com/nullstyle/coinop/test"
	. "github.com/nullstyle/coinop/usecase"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

var _ = Describe("ProcessDelivery", func() {
	var (
		repo     *MockDeliveryRepository
		send     *MockDeliverySender
		delivery entity.Delivery
		err      error
	)

	BeforeEach(func() {
		repo = &MockDeliveryRepository{}
		send = &MockDeliverySender{}
	})

	AfterEach(func() {
		test.VerifyMock(repo.Mock, send.Mock)
	})

	JustBeforeEach(func() {
		subject := &ProcessDelivery{repo, send}
		err = subject.Exec(context.Background(), delivery)
	})

	Context("working ports", func() {
		BeforeEach(func() {
			delivery = fake.DeliveryEntity()
			repo.
				On("StartDelivery", delivery).
				Return(int64(1), time.Now().Add(1*time.Second), nil)
			repo.On("MarkDeliverySuccess", int64(1), delivery).Return(nil)
			send.On("SendDelivery", delivery).Return(nil)
		})

		It("succeeds", func() {
			Expect(err).To(BeNil())
		})

	})

	Context("invalid delivery", func() {
		BeforeEach(func() {
			delivery = fake.DeliveryEntity()
			delivery.URL = nil
		})
		It("fails", func() {
			Expect(err).To(MatchError("invalid delivery: empty url"))
		})
	})

	Context("sender stalls for longer than deadline", func() {
		BeforeEach(func() {
			delivery = fake.DeliveryEntity()
			deadline := time.Now().Add(10 * time.Millisecond)
			repo.On("StartDelivery", delivery).Return(int64(1), deadline, nil)
			send.On("SendDelivery", delivery).
				WaitUntil(time.After(15 * time.Millisecond)).
				Return(nil)
		})

		It("fails", func() {
			Expect(err).To(MatchError("deadline exceeded"))
		})
	})

	Context("sender fails to deliver", func() {
		Context("repo fails to mark delivery failed", func() {})
	})

	Context("repo fails to mark delivery finished", func() {})
	Context("repo fails to mark delivery started", func() {})

	Context("a previously started delivery", func() {
		Context("after the delivery window has expired", func() {

		})
	})
	Context("a failed delivery", func() {})
	Context("a finished delivery", func() {})

})
