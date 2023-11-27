package controller

import (
	"fmt"
	"go_laundry/model"
	"go_laundry/usecase"
)

type EmployeeController struct {
	employeeUC usecase.EmployeeUseCase
}

func (e *EmployeeController) EmployeeMenuForm() {
	fmt.Print(`
	|		+++++ Master Employee +++++	|
	| 1. Tambah Data				|
	| 2. Lihat Data					|
	| 3. Update Data				|
	| 4. Hapus Data					|
	| 5. Cari Data Berdasarkan ID	|
	| 6. Cari Data Berdasarkan Phone|
	| 7. Keluar                     |
	`)
	fmt.Println("Pilih Menu (1-7): ")
	var selectMenuEmployee string
	fmt.Scanln(&selectMenuEmployee)
	switch selectMenuEmployee {
	case "1":
		e.insertFormEmployee()
	case "2":
		e.showListFormEmployee()
	case "3":
		e.updateFormEmployee()
	case "4":
		e.deleteFormEmployee()
	case "5":
		e.getByIdFormEmployee()
	case "6":
		e.getByPhoneFormEmployee()
	case "7":
		return
	}
}

func (e *EmployeeController) insertFormEmployee() {
	var employee model.Employee
	fmt.Println("Inputkan Id")
	fmt.Scanln(&employee.Id)
	fmt.Println("Inputkan Name")
	fmt.Scanln(&employee.Name)
	fmt.Println("Inputkan Phone Number")
	fmt.Scanln(&employee.Phone_number)
	fmt.Println("Inputkan Address")
	fmt.Scanln(&employee.Address)
	err := e.employeeUC.CreateNew(employee)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (e *EmployeeController) showListFormEmployee() {
	employees, err := e.employeeUC.FindAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(employees) == 0 {
		fmt.Println("Employee is empty")
		return
	}

	for _, employee := range employees {
		fmt.Printf("ID: %s, Name: %s, Phone: %s, Address: %s\n", employee.Id, employee.Name, employee.Phone_number, employee.Address)
	}
}

func (e *EmployeeController) updateFormEmployee() {
	var employee model.Employee
	fmt.Println("Inputkan Id")
	fmt.Scanln(&employee.Id)
	fmt.Println("Inputkan Name")
	fmt.Scanln(&employee.Name)
	fmt.Println("Inputkan Phone Number")
	fmt.Scanln(&employee.Phone_number)
	fmt.Println("Inputkan Address")
	fmt.Scanln(&employee.Address)
	err := e.employeeUC.Update(employee)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data Berhasil diUpdate")
}

func (e *EmployeeController) deleteFormEmployee() {
	var employee model.Employee
	fmt.Println("Inputkan Id")
	fmt.Scanln(&employee.Id)
	err := e.employeeUC.Delete(employee.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data Berhasil dihapus!")
}

func (e *EmployeeController) getByIdFormEmployee() {
	var employee model.Employee
	fmt.Println("Inputkan Id")
	fmt.Scanln(&employee.Id)
	employees, err := e.employeeUC.FindById(employee.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ID: %s, Name: %s, Phone: %s, Address: %s\n", employees.Id, employees.Name, employees.Phone_number, employees.Address)

}
func (e *EmployeeController) getByPhoneFormEmployee() {
	var employee model.Employee
	fmt.Println("Inputkan Phone Number")
	fmt.Scanln(&employee.Phone_number)
	employees, err := e.employeeUC.GetByPhone(employee.Phone_number)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, employee := range employees {
		fmt.Printf("ID: %s, Name: %s, Phone: %s, Address: %s\n", employee.Id, employee.Name, employee.Phone_number, employee.Address)
	}

}

func NewEmployeeController(employeeUC usecase.EmployeeUseCase) *EmployeeController {
	return &EmployeeController{employeeUC: employeeUC}
}
