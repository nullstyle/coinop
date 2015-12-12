package usecase_test

import (
	// "github.com/nullstyle/coinop/entity"
	. "github.com/nullstyle/coinop/usecase"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("WatchPayments", func() {
	var (
		subject WatchPayments
		network PaymentProvider
		// err     error
	)

	JustBeforeEach(func() {
		subject.Payments = network
		subject.Processor.Hooks = &MockWebhookRepository{}
		subject.Processor.Deliveries = &MockDeliveryRepository{}
		subject.Processor.Sender = &MockDeliverySender{}

		// err = subject.Exec()
	})

	// Context("a working repo", func() { })
	// Context("an erroring repo", func() {})

	PIt("saves the last seen paging token into the ")
})
