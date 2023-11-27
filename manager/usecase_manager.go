package manager

import "go_laundry/usecase"

type UsecaseManager interface {
	CustomerUC() usecase.CustomerUseCase
	EmployeeUC() usecase.EmployeeUseCase
	ProductUC() usecase.ProductUseCase
	UomUC() usecase.UomUseCase
	BillUC() usecase.BillUseCase
	AuthUC() usecase.AuthUseCase
	UserUC() usecase.UserUseCase
}

type usecaseManager struct {
	repoManager RepoManager
}

// AuthUC implements UsecaseManager.
func (u *usecaseManager) AuthUC() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.repoManager.UserRepo())
}

// UserUC implements UsecaseManager.
func (u *usecaseManager) UserUC() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepo())
}

// BillUC implements UsecaseManager.
func (u *usecaseManager) BillUC() usecase.BillUseCase {
	return usecase.NewBillUseCase(u.repoManager.BillRepo(), u.EmployeeUC(), u.CustomerUC(), u.ProductUC())
}

// CustomerUC implements UsecaseManager.
func (u *usecaseManager) CustomerUC() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManager.CustomerRepo())
}

// EmployeeUC implements UsecaseManager.
func (u *usecaseManager) EmployeeUC() usecase.EmployeeUseCase {
	return usecase.NewEmployeeUseCase(u.repoManager.EmployeeRepo())
}

// ProductUC implements UsecaseManager.
func (u *usecaseManager) ProductUC() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repoManager.ProductRepo(), u.UomUC())
}

// UomUC implements UsecaseManager.
func (u *usecaseManager) UomUC() usecase.UomUseCase {
	return usecase.NewUomUseCase(u.repoManager.UomRepo())
}

func NewUseCaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
