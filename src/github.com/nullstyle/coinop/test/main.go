package test

import (
	"github.com/onsi/ginkgo"
	//
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

// VerifyMock asserts the expectations on each of `ms`.
func VerifyMock(ms ...mock.Mock) {
	for _, m := range ms {
		m.AssertExpectations(ginkgo.GinkgoT())
	}
}

// ItSucceeds is a shortcut for checking error status on an operation.
func ItSucceeds(err *error) {
	It("succeeds", func() {
		Expect(*err).To(BeNil())
	})
}
