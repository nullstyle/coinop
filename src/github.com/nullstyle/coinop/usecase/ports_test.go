package usecase_test

import . "github.com/nullstyle/coinop/usecase"
import "github.com/nullstyle/coinop/entity"
import "github.com/stretchr/testify/mock"
import "time"

type MockDeliveryRepository struct {
	mock.Mock
}

func (_m *MockDeliveryRepository) LoadCursor() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockDeliveryRepository) SaveDeliveries(cursor string, d []entity.Delivery) error {
	ret := _m.Called(cursor, d)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []entity.Delivery) error); ok {
		r0 = rf(cursor, d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockDeliveryRepository) StartDelivery(_a0 entity.Delivery) (int64, time.Time, error) {
	ret := _m.Called(_a0)

	var r0 int64
	if rf, ok := ret.Get(0).(func(entity.Delivery) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 time.Time
	if rf, ok := ret.Get(1).(func(entity.Delivery) time.Time); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(time.Time)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(entity.Delivery) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
func (_m *MockDeliveryRepository) MarkDeliverySuccess(token int64, delivery entity.Delivery) error {
	ret := _m.Called(token, delivery)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, entity.Delivery) error); ok {
		r0 = rf(token, delivery)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockDeliveryRepository) MarkDeliveryFailed(token int64, delivery entity.Delivery) error {
	ret := _m.Called(token, delivery)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, entity.Delivery) error); ok {
		r0 = rf(token, delivery)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockDeliveryRepository) FailedDeliveries() ([]entity.Delivery, error) {
	ret := _m.Called()

	var r0 []entity.Delivery
	if rf, ok := ret.Get(0).(func() []entity.Delivery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Delivery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type MockDeliverySender struct {
	mock.Mock
}

func (_m *MockDeliverySender) SendDelivery(_a0 entity.Delivery) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Delivery) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type MockPaymentProvider struct {
	mock.Mock
}

func (_m *MockPaymentProvider) StreamPayments(cursor string, fn PaymentHandler) error {
	ret := _m.Called(cursor, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, PaymentHandler) error); ok {
		r0 = rf(cursor, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type MockWebhookPresenter struct {
	mock.Mock
}

func (_m *MockWebhookPresenter) List(_a0 []entity.Webhook) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]entity.Webhook) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type MockWebhookRepository struct {
	mock.Mock
}

func (_m *MockWebhookRepository) SaveWebhook(_a0 *entity.Webhook) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Webhook) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockWebhookRepository) DestroyWebhook(ID RepoID) error {
	ret := _m.Called(ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(RepoID) error); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockWebhookRepository) ListWebhooks() ([]entity.Webhook, error) {
	ret := _m.Called()

	var r0 []entity.Webhook
	if rf, ok := ret.Get(0).(func() []entity.Webhook); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Webhook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *MockWebhookRepository) ForDestination(dest entity.AccountID) ([]entity.Webhook, error) {
	ret := _m.Called(dest)

	var r0 []entity.Webhook
	if rf, ok := ret.Get(0).(func(entity.AccountID) []entity.Webhook); ok {
		r0 = rf(dest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Webhook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.AccountID) error); ok {
		r1 = rf(dest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockPaymentProcessor struct {
	mock.Mock
}

func (_m *mockPaymentProcessor) Exec(p entity.Payment) error {
	ret := _m.Called(p)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Payment) error); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
