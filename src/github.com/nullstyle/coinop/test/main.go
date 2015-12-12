package test

import (
	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/mock"
)

// VerifyMock asserts the expectations on each of `ms`.
func VerifyMock(ms ...mock.Mock) {
	for _, m := range ms {
		m.AssertExpectations(ginkgo.GinkgoT())
	}
}
