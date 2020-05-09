package usecase

import (
	"owner/owner/repository"
	"svc-customer/customer/entity"
)

// CustomerDTO Data Transfer Object structure for entity
type CustomerDTO struct {
	entity.Customer
}

// GetCustomerImpl implementation
type GetCustomerImpl struct {
	Repo repository.Repository
}

// HealthCheck UseCase check database health check
func (uc *GetCustomerImpl) HealthCheck() error {
	err := uc.Repo.HealthCheck()
	if err != nil {
		return entity.ErrDatabaseError
	}
	return nil
}

// Store Create Customer inside Database by customer entity
func (uc *GetCustomerImpl) Store(customer entity.Customer) error {
	err := uc.Repo.Store(customer)
	if err != nil {
		return err
	}
	return nil
}

// GetByUUID Get Customer Personal information from Database by ID usecase
func (uc *GetCustomerImpl) GetByUUID(customerUUID string) (*entity.Customer, error) {
	customer, err := uc.Repo.GetByUUID(customerUUID)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

// Fetch Get Customer Personal information from Database by ID usecase
func (uc *GetCustomerImpl) Fetch(page int, limit int) ([]*entity.Customer, int, int, int, error) {
	customers, total, pages, page, err := uc.Repo.Fetch(page, limit)
	if err != nil {
		return nil, total, pages, page, err
	}
	return customers, total, pages, page, nil
}

// UpdateByUUID Get Customer Personal information from Database by ID usecase
func (uc *GetCustomerImpl) UpdateByUUID(customer entity.Customer, customerUUID string) error {
	err := uc.Repo.UpdateByUUID(customer, customerUUID)
	if err != nil {
		if err == entity.ErrorNotFoundOnDB {
			return entity.ErrNotFound
		}
		return err
	}
	return nil
}

// DeleteByUUID Delete Customer from Database by ID usecase
func (uc *GetCustomerImpl) DeleteByUUID(customerUUID string) error {
	err := uc.Repo.DeleteByUUID(customerUUID)
	if err != nil {
		return err
	}
	return nil
}
