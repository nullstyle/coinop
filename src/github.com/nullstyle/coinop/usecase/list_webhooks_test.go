package usecase_test

import (
	"errors"

	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/fake"
	"github.com/nullstyle/coinop/test"
	. "github.com/nullstyle/coinop/usecase"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListWebhooks", func() {
	var (
		subject ListWebhooks
		repo    *MockWebhookRepository
		pres    *MockWebhookPresenter
		err     error
	)

	BeforeEach(func() {
		repo = &MockWebhookRepository{}
		pres = &MockWebhookPresenter{}
	})

	AfterEach(func() {
		test.VerifyMock(repo.Mock, pres.Mock)
	})

	JustBeforeEach(func() {
		subject = ListWebhooks{repo, pres}
		err = subject.Exec()
	})

	Context("no webhooks", func() {
		BeforeEach(func() {
			repo.On("ListWebhooks").Return([]entity.Webhook{}, nil)
			pres.On("List", []entity.Webhook{}).Return(nil)
		})

		It("succeeds", func() {
			Expect(err).To(BeNil())
		})
	})

	Context("has webhooks", func() {
		BeforeEach(func() {
			es := []entity.Webhook{
				fake.WebhookEntity(),
				fake.WebhookEntity(),
			}
			repo.On("ListWebhooks").Return(es, nil)
			pres.On("List", es).Return(nil)
		})

		It("succeeds", func() {
			Expect(err).To(BeNil())
		})
	})

	Context("erroring repo", func() {
		BeforeEach(func() {
			repo.On("ListWebhooks").Return(nil, errors.New("boom"))
		})

		It("fails", func() {
			Expect(err).To(MatchError("boom"))
		})
	})

	Context("erroring presenter", func() {
		BeforeEach(func() {
			repo.On("ListWebhooks").Return([]entity.Webhook{}, nil)
			pres.On("List", []entity.Webhook{}).Return(errors.New("boom"))
		})

		It("fails", func() {
			Expect(err).To(MatchError("boom"))
		})
	})
})
