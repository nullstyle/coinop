package test

import (
	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/mock"
)

// VerifyMock asserts the expectations on `m`.
func VerifyMock(m mock.Mock) {
	m.AssertExpectations(ginkgo.GinkgoT())
}
