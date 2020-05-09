package usecase

import (
	"customer/customer/entity"
	"owner/owner/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsecaseGetByUUIDValid(t *testing.T) {

	mockCustomer := entity.Customer{
		CustomerUUID: "039d69ee-f9cb-4a3d-87e4-6eb63c302579",
		Name:         "Eddwin",
		LastName:     "Paz",
		Email:        "ep@latamig.com",
		Phone:        "56933375029",
		Country:      "CL",
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("GetByUUID", mockCustomer.CustomerUUID).Return(&mockCustomer, nil)

	// Call Usecase Implementation
	u := GetCustomerImpl{
		Repo: mockCase,
	}

	// Execute
	cust, err := u.GetByUUID(mockCustomer.CustomerUUID)

	assert.NoError(t, err)
	assert.NotNil(t, cust)
}

func TestUsecaseGetByUUIDInvalidUUID(t *testing.T) {

	// Bad UUID Parameter
	customerUUID := "039d69ee-4a3d-87e4-6eb63c302579"

	// Mock Return
	mockCase := new(mocks.MockedRepository)
	mockCase.On("GetByUUID", customerUUID).Return(nil, entity.ErrBadParamInput)

	// Call Usecase Implementation
	u := GetCustomerImpl{
		Repo: mockCase,
	}

	// Execute
	cust, err := u.GetByUUID(customerUUID)

	// Validate Expected
	assert.Equal(t, err, entity.ErrBadParamInput)
	assert.Nil(t, cust)

	mockCase.AssertExpectations(t)
}

func TestUsecaseStoreValid(t *testing.T) {

	mockCustomer := entity.Customer{
		Name:     "Eddwin",
		LastName: "Paz",
		Email:    "ep@latamig.com",
		Phone:    "56933375029",
		Country:  "CL",
		Password: "monkey123",
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("Store", mockCustomer).Return(nil)

	// Call Usecase Implementation
	u := GetCustomerImpl{
		Repo: mockCase,
	}

	// Execute
	err := u.Store(mockCustomer)

	assert.NoError(t, err)
}

func TestUsecaseStoreInvalid(t *testing.T) {

	mockCustomer := entity.Customer{}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("Store", mockCustomer).Return(entity.ErrSQLError)

	// Call Usecase Implementation
	u := GetCustomerImpl{
		Repo: mockCase,
	}

	// Execute
	err := u.Store(mockCustomer)

	// Validate Expects
	assert.EqualError(t, entity.ErrSQLError, err.Error())
}

func TestUsecaseUpdateByUUIDValid(t *testing.T) {

	customerUUID := "039d69ee-f9cb-4a3d-87e4-6eb63c302579"

	mockCustomer := entity.Customer{
		CustomerUUID: "",
		Name:         "Eddwin",
		LastName:     "Paz",
		Email:        "ep@latamig.com",
		Phone:        "56933375029",
		Country:      "CL",
		Password:     "",
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("UpdateByUUID", mockCustomer, customerUUID).Return(nil)

	// Call Usecase Implementation
	u := GetCustomerImpl{
		Repo: mockCase,
	}

	// Execute
	err := u.UpdateByUUID(mockCustomer, customerUUID)

	assert.NoError(t, err)
}

func TestUsecaseUpdateByUUIDInvalid(t *testing.T) {

	// customerUUID := "039d69ee-f9cb-4a3d-87e4-6eb63c302579"

	mockCustomer := entity.Customer{
		CustomerUUID: "",
		Name:         "Eddwin",
		LastName:     "Paz",
		Email:        "ep@latamig.com",
		Phone:        "56933375029",
		Country:      "CL",
		Password:     "",
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("UpdateByUUID", mockCustomer, "").Return(entity.ErrNotFound)

	// Call Usecase Implementation
	u := GetCustomerImpl{
		Repo: mockCase,
	}

	// Execute
	err := u.UpdateByUUID(mockCustomer, "")

	assert.Error(t, entity.ErrNotFound, err.Error())
}

func TestUsecaseUpdateByUUIDInvalidDB(t *testing.T) {

	// customerUUID := "039d69ee-f9cb-4a3d-87e4-6eb63c302579"

	mockCustomer := entity.Customer{
		CustomerUUID: "",
		Name:         "Eddwin",
		LastName:     "Paz",
		Email:        "ep@latamig.com",
		Phone:        "56933375029",
		Country:      "CL",
		Password:     "",
	}

	mockCase := new(mocks.MockedRepository)
	mockCase.On("UpdateByUUID", mockCustomer, "").Return(entity.ErrorNotFoundOnDB)

	// Call Usecase Implementation
	u := GetCustomerImpl{
		Repo: mockCase,
	}

	// Execute
	err := u.UpdateByUUID(mockCustomer, "")

	assert.Error(t, entity.ErrNotFound, err.Error())
}

func TestUsecaseDeleteByUUIDValid(t *testing.T) {

	// Bad UUID Parameter
	customerUUID := "039d69ee-4a3d-87e4-6eb63c302579"

	// Mock Return
	mockCase := new(mocks.MockedRepository)
	mockCase.On("DeleteByUUID", customerUUID).Return(nil)

	// Call Usecase Implementation
	u := GetCustomerImpl{
		Repo: mockCase,
	}

	// Execute
	err := u.DeleteByUUID(customerUUID)

	// Validate Expected
	assert.NoError(t, err)

	mockCase.AssertExpectations(t)
}

func TestUsecaseDeleteByUUIDInvalid(t *testing.T) {

	// Mock Return
	mockCase := new(mocks.MockedRepository)
	mockCase.On("DeleteByUUID", "").Return(entity.ErrSQLError)

	// Call Usecase Implementation
	u := GetCustomerImpl{
		Repo: mockCase,
	}

	// Execute
	err := u.DeleteByUUID("")

	// Validate Expected
	assert.EqualError(t, entity.ErrSQLError, err.Error())

	mockCase.AssertExpectations(t)
}

func TestUsecaseFetchValid(t *testing.T) {

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

	mockCase := new(mocks.MockedRepository)
	mockCase.On("Fetch").Return(rows, nil)

	// Call Usecase Implementation
	u := GetCustomerImpl{
		Repo: mockCase,
	}

	// Execute
	cust, _, _, _, err := u.Fetch(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, rows, cust)
}

func TestUsecaseHealthCheckValid(t *testing.T) {

	mockCase := new(mocks.MockedRepository)
	mockCase.On("HealthCheck").Return(nil)

	// Call Usecase Implementation
	u := GetCustomerImpl{
		Repo: mockCase,
	}

	// Execute
	err := u.HealthCheck()

	assert.NoError(t, err)
}

func TestUsecaseHealthCheckError(t *testing.T) {

	mockCase := new(mocks.MockedRepository)
	mockCase.On("HealthCheck").Return(entity.ErrDatabaseError)

	// Call Usecase Implementation
	u := GetCustomerImpl{
		Repo: mockCase,
	}

	// Execute
	err := u.HealthCheck()
	assert.Error(t, entity.ErrDatabaseError, err)
}
