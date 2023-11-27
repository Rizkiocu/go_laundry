package usecase

import (
	"fmt"
	"go_laundry/model"
	"go_laundry/model/dto"
	"go_laundry/repository"
)

type CustomerUseCase interface {
	CreateNew(payload model.Customer) error
	FindById(id string) (model.Customer, error)
	FindAll() ([]model.Customer, error)
	GetByPhone(Phone_number string) ([]model.Customer, error)
	Update(payload model.Customer) error
	Delete(id string) error
	Paging(payload dto.PageRequest) ([]model.Customer, dto.Paging, error)
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

// Paging implements CustomerUseCase.
func (c *customerUseCase) Paging(payload dto.PageRequest) ([]model.Customer, dto.Paging, error) {
	return c.repo.Paging(payload)
}

// CreateNew implements CustomeUseCase.
func (c *customerUseCase) CreateNew(payload model.Customer) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	if payload.Phone_number == "" {
		return fmt.Errorf("phone is required")
	}
	if payload.Address == "" {
		return fmt.Errorf("address is required")
	}

	// Check if the phone number is already in use
	existingCustomer, err := c.repo.FindByPhone(payload.Phone_number)
	if err != nil {
		return fmt.Errorf("error checking phone number uniqueness: %v", err)
	}
	if existingCustomer != nil {
		return fmt.Errorf("phone number is already in use")
	}

	err = c.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to save new customer: %v", err)
	}
	return nil
}

// Delete implements CustomeUseCase.
func (c *customerUseCase) Delete(id string) error {
	customer, err := c.FindById(id)
	if err != nil {
		return err
	}
	err = c.repo.DeleteById(customer.Id)
	if err != nil {
		return fmt.Errorf("failed to delete customer: %v", err)
	}
	return nil
}

// FindAll implements CustomeUseCase.
func (c *customerUseCase) FindAll() ([]model.Customer, error) {
	customers, err := c.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch customer: %v", err)
	}
	return customers, nil
}

// FindById implements CustomeUseCase.
func (c *customerUseCase) FindById(id string) (model.Customer, error) {
	customer, err := c.repo.FindById(id)
	if err != nil {
		return model.Customer{}, fmt.Errorf("id customer not found: %v", err)
	}

	return customer, nil
}
func (c *customerUseCase) GetByPhone(phone_number string) ([]model.Customer, error) {
	customers, err := c.repo.FindByPhone(phone_number)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch customer by Phone: %v", err)
	}
	return customers, nil
}

// Update implements CustomeUseCase.
func (c *customerUseCase) Update(payload model.Customer) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	if payload.Phone_number == "" {
		return fmt.Errorf("phone is required")
	}
	if payload.Address == "" {
		return fmt.Errorf("address is required")
	}

	existingCustomer, err := c.repo.FindByPhone(payload.Phone_number)
	if err != nil {
		return fmt.Errorf("error checking phone number uniqueness: %v", err)
	}
	if existingCustomer != nil {
		return fmt.Errorf("phone number is already in use")
	}

	err = c.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update customer: %v", err)
	}
	return nil
}
func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{repo: repo}
}
