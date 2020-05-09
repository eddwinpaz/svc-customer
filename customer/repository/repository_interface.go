package repository

import "svc-customer/customer/entity"

// Repository interface to repository
type Repository interface {
	HealthCheck() error
	GetByUUID(customerUUID string) (*entity.Customer, error)
	Fetch(page int, limit int) ([]*entity.Customer, int, int, int, error)
	Store(customer entity.Customer) error
	UpdateByUUID(customer entity.Customer, customerUUID string) error
	DeleteByUUID(customerUUID string) error
	// RequestCustomerToken(email string, password string) (entity.Customer, error)
}
