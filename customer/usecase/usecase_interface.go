package usecase

import "svc-customer/customer/entity"

// Usecase interface
type Usecase interface {
	HealthCheck() error
	GetByUUID(customerUUID string) (*entity.Customer, error) // CustomerDTO
	Fetch(page int, limit int) ([]*entity.Customer, int, int, int, error)
	Store(customer entity.Customer) error
	UpdateByUUID(customer entity.Customer, customerUUID string) error
	DeleteByUUID(customerUUID string) error
	// RequestCustomerToken(email string, password string) (CustomerDTO, error)
}
