package manager

import (
	"go_laundry/repository"
)

type RepoManager interface {
	CustomerRepo() repository.CustomerRepository
	EmployeeRepo() repository.EmployeeRepository
	ProductRepo() repository.ProductRepository
	UomRepo() repository.UomRepository
	BillRepo() repository.BillRepository
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infraManager InfraManager
}

// UserRepo implements RepoManager.
func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infraManager.Conn())
}

// BillRepo implements RepoManager.
func (r *repoManager) BillRepo() repository.BillRepository {
	return repository.NewBillRepository(r.infraManager.Conn())
}

// CustomerRepo implements RepoManager.
func (r *repoManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infraManager.Conn())
}

// EmployeeRepo implements RepoManager.
func (r *repoManager) EmployeeRepo() repository.EmployeeRepository {
	return repository.NewEmployeeRepository(r.infraManager.Conn())
}

// ProductRepo implements RepoManager.
func (r *repoManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infraManager.Conn())
}

// UomRepo implements RepoManager.
func (r *repoManager) UomRepo() repository.UomRepository {
	return repository.NewUomRepository(r.infraManager.Conn())
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
