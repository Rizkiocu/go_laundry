package usecase

import (
	"fmt"
	"go_laundry/model"
	"go_laundry/model/dto"
	"go_laundry/repository"
)

type EmployeeUseCase interface {
	CreateNew(payload model.Employee) error
	FindById(id string) (model.Employee, error)
	FindAll() ([]model.Employee, error)
	GetByPhone(Phone_number string) ([]model.Employee, error)
	Update(payload model.Employee) error
	Delete(id string) error
	Paging(payload dto.PageRequest) ([]model.Employee, dto.Paging, error)
}

type employeeUseCase struct {
	repo repository.EmployeeRepository
}

// Paging implements EmployeeUseCase.
func (e *employeeUseCase) Paging(payload dto.PageRequest) ([]model.Employee, dto.Paging, error) {
	return e.repo.Paging(payload)
}

// CreateNew implements EmployeeUseCase.
func (e *employeeUseCase) CreateNew(payload model.Employee) error {
	if payload.Id == "" {
		return fmt.Errorf("Id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("Name is required")
	}
	if payload.Phone_number == "" {
		return fmt.Errorf("Phone is required")
	}
	if payload.Address == "" {
		return fmt.Errorf("Adress is required")
	}

	// Check if the phone number is already in use
	existingEmployee, err := e.repo.FindByPhone(payload.Phone_number)
	if err != nil {
		return fmt.Errorf("error checking phone number uniqueness: %v", err)
	}
	if existingEmployee != nil {
		return fmt.Errorf("phone number is already in use")
	}

	err = e.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to save new Employee: %v", err)
	}
	return nil
}

// Delete implements EmployeeUseCase.
func (e *employeeUseCase) Delete(id string) error {
	employee, err := e.FindById(id)
	if err != nil {
		return err
	}
	err = e.repo.DeleteById(employee.Id)
	if err != nil {
		return fmt.Errorf("failed to delete employee: %v", err)
	}
	return nil
}

// FindAll implements EmployeeUseCase.
func (e *employeeUseCase) FindAll() ([]model.Employee, error) {
	employees, err := e.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch employee: %v", err)
	}
	return employees, nil
}

// FindById implements EmployeeUseCase.
func (e *employeeUseCase) FindById(id string) (model.Employee, error) {
	employee, err := e.repo.FindById(id)
	if err != nil {
		return model.Employee{}, fmt.Errorf("id employee not found: %v", err)
	}

	return employee, nil
}
func (e *employeeUseCase) GetByPhone(phone_number string) ([]model.Employee, error) {
	employees, err := e.repo.FindByPhone(phone_number)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch employee by Phone: %v", err)
	}
	return employees, nil
}

// Update implements EmployeeUseCase.
func (e *employeeUseCase) Update(payload model.Employee) error {
	if payload.Id == "" {
		return fmt.Errorf("Id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("Name is required")
	}
	if payload.Phone_number == "" {
		return fmt.Errorf("Phone is required")
	}
	if payload.Address == "" {
		return fmt.Errorf("Address is required")
	}
	existingEmployee, err := e.repo.FindByPhone(payload.Phone_number)
	if err != nil {
		return fmt.Errorf("error checking phone number uniqueness: %v", err)
	}
	if existingEmployee != nil {
		return fmt.Errorf("phone number is already in use")
	}

	err = e.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update Employee: %v", err)
	}
	return nil
}
func NewEmployeeUseCase(repo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{repo: repo}
}
