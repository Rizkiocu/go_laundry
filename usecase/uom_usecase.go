package usecase

import (
	"fmt"
	"go_laundry/model"
	"go_laundry/model/dto"
	"go_laundry/repository"
	"go_laundry/util/common"
)

type UomUseCase interface {
	CreateNew(payload model.Uom) error
	FindById(id string) (model.Uom, error)
	FindAll() ([]model.Uom, error)
	Update(payload model.Uom) error
	Delete(id string) error
	Paging(payload dto.PageRequest) ([]model.Uom, dto.Paging, error)
}

type uomUseCase struct {
	repo repository.UomRepository
}

// Paging implements UomUseCase.
func (u *uomUseCase) Paging(payload dto.PageRequest) ([]model.Uom, dto.Paging, error) {
	return u.repo.Paging(payload)
}

// CreateNew implements UomUseCase.
func (u *uomUseCase) CreateNew(payload model.Uom) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	payload.Id = common.GenerateID()
	err := u.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to create new uom: %v", err)
	}
	return nil
}

// Delete implements UomUseCase.
func (u *uomUseCase) Delete(id string) error {
	uom, err := u.FindById(id)
	if err != nil {
		return err
	}

	err = u.repo.DeleteById(uom.Id)
	if err != nil {
		return fmt.Errorf("Failed to delete uom: %v", err)
	}
	return nil
}

// FindAll implements UomUseCase.
func (u *uomUseCase) FindAll() ([]model.Uom, error) {
	uom, err := u.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find all uom: %v", err)
	}
	return uom, nil
}

// FindById implements UomUseCase.
func (u *uomUseCase) FindById(id string) (model.Uom, error) {
	uom, err := u.repo.FindById(id)
	if err != nil {
		return model.Uom{}, fmt.Errorf("uom not found")
	}
	return uom, nil
}

// Update implements UomUseCase.
func (u *uomUseCase) Update(payload model.Uom) error {

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	_, err := u.FindById(payload.Id)
	if err != nil {
		return err
	}

	err = u.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update uom: %v", err)
	}
	return nil
}

func NewUomUseCase(repo repository.UomRepository) UomUseCase {
	return &uomUseCase{repo: repo}
}
