package usecase

import (
	"fmt"
	"go_laundry/model"
	"go_laundry/repository"
)

type ProductUseCase interface {
	CreateNew(payload model.Product) error
	FindById(id string) (model.Product, error)
	FindAll() ([]model.Product, error)
	GetByName(name string) ([]model.Product, error)
	Update(payload model.Product) error
	Delete(id string) error
}

type productUseCase struct {
	repo  repository.ProductRepository
	uomUC UomUseCase
}

// CreateNew implements ProductUseCase.
func (p *productUseCase) CreateNew(payload model.Product) error {
	if payload.Id == "" {
		return fmt.Errorf("Id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("Name is required")
	}
	if payload.Price <= 0 {
		return fmt.Errorf("Price must be greater than zero")
	}
	_, err := p.uomUC.FindById(payload.Uom.Id)
	if err != nil {
		return err
	}

	err = p.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to save new product: %v", err)
	}
	return nil
}

// Delete implements ProductUseCase.
func (p *productUseCase) Delete(id string) error {
	uom, err := p.FindById(id)
	if err != nil {
		return err
	}
	err = p.repo.DeleteById(uom.Id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}
	return nil
}

// FindAll implements ProductUseCase.
func (p *productUseCase) FindAll() ([]model.Product, error) {
	products, err := p.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %v", err)
	}
	return products, nil
}

// FindById implements ProductUseCase.
func (p *productUseCase) FindById(id string) (model.Product, error) {
	product, err := p.repo.FindById(id)
	if err != nil {
		return model.Product{}, fmt.Errorf("id product not found: %v", err)
	}
	return product, nil
}
func (p *productUseCase) GetByName(name string) ([]model.Product, error) {
	products, err := p.repo.FindByName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products by name: %v", err)
	}
	return products, nil
}

// Update implements ProductUseCase.
func (p *productUseCase) Update(payload model.Product) error {
	if payload.Id == "" {
		return fmt.Errorf("Id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("Name is required")
	}
	if payload.Price <= 0 {
		return fmt.Errorf("Price must be greater than zero")
	}
	_, err := p.uomUC.FindById(payload.Uom.Id)
	if err != nil {
		return err
	}

	err = p.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update product: %v", err)
	}
	return nil
}
func NewProductUseCase(repo repository.ProductRepository, uomUC UomUseCase) ProductUseCase {
	return &productUseCase{
		repo:  repo,
		uomUC: uomUC,
	}
}
