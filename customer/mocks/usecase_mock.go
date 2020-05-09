package mocks

import (
	"customer/customer/entity"

	"github.com/stretchr/testify/mock"
)

// MockedRepository is a mocked object that implements an interface
// that describes an object that the code I am testing relies on.
type MockedRepository struct {
	mock.Mock
}

// DeleteByUUID is a method on MyMockedObject that implements some interface
// and just records the activity, and returns what the Mock object tells it to.
//
// In the real object, this method would do something useful, but since this
// is a mocked object - we're just going to stub it out.
//
// NOTE: This method is not being tested here, code that uses this object is.
func (m *MockedRepository) DeleteByUUID(customerUUID string) error {
	args := m.Called(customerUUID)
	return args.Error(0)
}

// HealthCheck Check Database Connection
func (m *MockedRepository) HealthCheck() error {
	args := m.Called()
	return args.Error(0)
}

// GetByUUID a
func (m *MockedRepository) GetByUUID(customerUUID string) (*entity.Customer, error) {

	var customer *entity.Customer

	args := m.Called(customerUUID)

	if rf, ok := args.Get(0).(func(string) *entity.Customer); ok {
		customer = rf(customerUUID)
	} else {
		if args.Get(0) != nil {
			customer = args.Get(0).(*entity.Customer)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(string) error); ok {
		r1 = rf(customerUUID)
	} else {
		r1 = args.Error(1)
	}

	return customer, r1
}

// Fetch a
func (m *MockedRepository) Fetch(page int, limit int) ([]*entity.Customer, int, int, int, error) {

	mockCustomer := &entity.Customer{
		CustomerUUID: "039d69ee-f9cb-4a3d-87e4-6eb63c302579",
		Name:         "Eddwin",
		LastName:     "Paz",
		Dni:          "264573076",
		DniType:      "DNI",
		Email:        "ep@latamig.com",
		Phone:        "56933375029",
		Country:      "CL",
	}

	var rows []*entity.Customer
	rows = append(rows, mockCustomer)

	var err error
	var total = 1
	var pages = 1

	return rows, total, pages, page, err
}

// Store s
func (m *MockedRepository) Store(customer entity.Customer) error {
	args := m.Called(customer)
	var r0 error
	if rf, ok := args.Get(0).(func(entity.Customer) error); ok {
		r0 = rf(customer)
	} else {
		r0 = args.Error(0)
	}
	return r0
}

// UpdateByUUID a
func (m *MockedRepository) UpdateByUUID(customer entity.Customer, customerUUID string) error {

	args := m.Called(customer, customerUUID)

	var r0 error

	if rf, ok := args.Get(0).(func(entity.Customer, string) error); ok {
		r0 = rf(customer, customerUUID)
	} else {
		r0 = args.Error(0)
	}
	return r0
}
