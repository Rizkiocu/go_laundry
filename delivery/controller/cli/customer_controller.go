package controller

import (
	"fmt"
	"go_laundry/model"
	"go_laundry/usecase"
)

type CustomerController struct {
	customerUC usecase.CustomerUseCase
}

func (c *CustomerController) CustomerMenuForm() {
	fmt.Print(`
	|		+++++ Master Customer +++++	|
	| 1. Tambah Data				|
	| 2. Lihat Data					|
	| 3. Update Data				|
	| 4. Hapus Data					|
	| 5. Cari Data Berdasarkan ID	|
	| 6. Cari Data Berdasarkan Phone|
	| 7. Keluar                     |
	`)
	fmt.Println("Pilih Menu (1-7): ")
	var selectMenuCustomer string
	fmt.Scanln(&selectMenuCustomer)
	switch selectMenuCustomer {
	case "1":
		c.insertFormCustomer()
	case "2":
		c.showListFormCustomer()
	case "3":
		c.updateFormCustomer()
	case "4":
		c.deleteFormCustomer()
	case "5":
		c.getByIdFormCustomer()
	case "6":
		c.getByPhoneFormCustomer()
	case "7":
		return
	}
}

func (c *CustomerController) insertFormCustomer() {
	var customer model.Customer
	fmt.Println("Inputkan Id")
	fmt.Scanln(&customer.Id)
	fmt.Println("Inputkan Name")
	fmt.Scanln(&customer.Name)
	fmt.Println("Inputkan Phone Number")
	fmt.Scanln(&customer.Phone_number)
	fmt.Println("Inputkan Address")
	fmt.Scanln(&customer.Address)
	err := c.customerUC.CreateNew(customer)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (c *CustomerController) showListFormCustomer() {
	customers, err := c.customerUC.FindAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(customers) == 0 {
		fmt.Println("Customer is empty")
		return
	}

	for _, customer := range customers {
		fmt.Printf("ID: %s, Name: %s, Phone: %s, Address: %s\n", customer.Id, customer.Name, customer.Phone_number, customer.Address)
	}
}

func (c *CustomerController) updateFormCustomer() {
	var customer model.Customer
	fmt.Println("Inputkan Id")
	fmt.Scanln(&customer.Id)
	fmt.Println("Inputkan Name")
	fmt.Scanln(&customer.Name)
	fmt.Println("Inputkan Phone Number")
	fmt.Scanln(&customer.Phone_number)
	fmt.Println("Inputkan Address")
	fmt.Scanln(&customer.Address)
	err := c.customerUC.Update(customer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data Berhasil diUpdate")
}

func (c *CustomerController) deleteFormCustomer() {
	var customer model.Customer
	fmt.Println("Inputkan Id")
	fmt.Scanln(&customer.Id)
	err := c.customerUC.Delete(customer.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data Berhasil dihapus!")
}

func (c *CustomerController) getByIdFormCustomer() {
	var customer model.Customer
	fmt.Println("Inputkan Id")
	fmt.Scanln(&customer.Id)
	customers, err := c.customerUC.FindById(customer.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ID: %s, Name: %s, Phone: %s, Address: %s\n", customers.Id, customers.Name, customers.Phone_number, customers.Address)

}
func (c *CustomerController) getByPhoneFormCustomer() {
	var customer model.Customer
	fmt.Println("Inputkan Phone Number")
	fmt.Scanln(&customer.Phone_number)
	customers, err := c.customerUC.GetByPhone(customer.Phone_number)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, customer := range customers {
		fmt.Printf("ID: %s, Name: %s, Phone: %s, Address: %s\n", customer.Id, customer.Name, customer.Phone_number, customer.Address)
	}

}

func NewCustomerController(customeUC usecase.CustomerUseCase) *CustomerController {
	return &CustomerController{customerUC: customeUC}
}
