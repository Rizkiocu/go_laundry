package usecase

import (
	"fmt"
	"go_laundry/model"
	"go_laundry/repository"
)

type BillUseCase interface {
	CreateNew(payload model.Bill) error
	FindById(id string) (model.Bill, error)
	FindAll() ([]model.Bill, error)
}

type billUseCase struct {
	repo       repository.BillRepository
	productUC  ProductUseCase
	employeeUC EmployeeUseCase
	customerUC CustomerUseCase
}

// FindAll implements BillUseCase.
func (b *billUseCase) FindAll() ([]model.Bill, error) {
	// Panggil repository untuk mencari semua data tagihan
	bills, err := b.repo.FindAll()
	if err != nil {
		return nil, err
	}

	// Untuk setiap tagihan, ambil data produk, karyawan, dan pelanggan terkait
	for i, bill := range bills {
		for j, detail := range bill.BillDetails {
			product, err := b.productUC.FindById(detail.Product.Id)
			if err != nil {
				return nil, fmt.Errorf("gagal mencari produk untuk BillDetail ke-%d: %v", j+1, err)
			}
			bills[i].BillDetails[j].Product = product
		}

		employee, err := b.employeeUC.FindById(bill.EmployeeId.Id)
		if err != nil {
			return nil, fmt.Errorf("Gagal mencari karyawan: %v", err)
		}
		bills[i].EmployeeId = employee

		customer, err := b.customerUC.FindById(bill.CustomerId.Id)
		if err != nil {
			return nil, fmt.Errorf("Gagal mencari pelanggan: %v", err)
		}
		bills[i].CustomerId = customer
	}

	return bills, nil
}

// FindById implements BillUseCase.
func (b *billUseCase) FindById(id string) (model.Bill, error) {
	// Panggil repository untuk mencari Bill berdasarkan ID
	bill, err := b.repo.FindById(id)
	if err != nil {
		return model.Bill{}, err
	}

	// Mengambil data produk untuk setiap BillDetail
	for i, detail := range bill.BillDetails {
		product, err := b.productUC.FindById(detail.Product.Id)
		if err != nil {
			return model.Bill{}, fmt.Errorf("Gagal mencari produk untuk BillDetail ke-%d: %v", i+1, err)
		}
		bill.BillDetails[i].Product = product
	}

	// Mengambil data karyawan
	employee, err := b.employeeUC.FindById(bill.EmployeeId.Id)
	if err != nil {
		return model.Bill{}, fmt.Errorf("Gagal mencari karyawan: %v", err)
	}
	bill.EmployeeId = employee

	// Mengambil data pelanggan
	customer, err := b.customerUC.FindById(bill.CustomerId.Id)
	if err != nil {
		return model.Bill{}, fmt.Errorf("Gagal mencari pelanggan: %v", err)
	}
	bill.CustomerId = customer

	return bill, nil
}

// CreateNew implements BillUseCase.
func (b *billUseCase) CreateNew(payload model.Bill) error {
	// Validasi data payload di sini jika diperlukan
	if payload.Id == "" {
		return fmt.Errorf("ID Bill tidak boleh kosong")
	}
	// Validasi ketersediaan produk
	for _, detail := range payload.BillDetails {
		_, err := b.productUC.FindById(detail.Product.Id)
		if err != nil {
			return fmt.Errorf("Gagal mencari produk: %v", err)
		}
		// if detail.Product.Price != detail.ProductPrice {
		// 	return fmt.Errorf("Harga produk '%s' tidak sesuai dengan harga yang diberikan di BillDetails", detail.Product.Name)
		// }

		if detail.Qty <= 0 {
			return fmt.Errorf("Qty tidak boleh kosong")
		}
	}

	// Validasi ketersediaan karyawan
	_, err := b.employeeUC.FindById(payload.EmployeeId.Id)
	if err != nil {
		return fmt.Errorf("Gagal mencari karyawan: %v", err)
	}

	// Validasi ketersediaan pelanggan
	_, err = b.customerUC.FindById(payload.CustomerId.Id)
	if err != nil {
		return fmt.Errorf("Gagal mencari pelanggan: %v", err)
	}

	// Panggil repository untuk menyimpan data Bill
	err = b.repo.Save(payload)
	if err != nil {
		return err
	}

	return nil

}

func NewBillUseCase(repo repository.BillRepository, employeeUC EmployeeUseCase, customerUC CustomerUseCase, productUC ProductUseCase) BillUseCase {
	return &billUseCase{
		repo:       repo,
		employeeUC: employeeUC,
		customerUC: customerUC,
		productUC:  productUC,
	}
}
