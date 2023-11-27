package controller

import (
	"fmt"
	"go_laundry/model"
	"go_laundry/usecase"
	"time"
)

type BillController struct {
	billUC usecase.BillUseCase
}

func (b *BillController) BillMenuForm() {
	fmt.Print(`
	|		+++++ Master Product +++++	|
	| 1. Tambah Data					|
	| 2. Lihat Data						|
	| 3. lihat data by id				|
	| 4. Keluar                     	|
	`)
	fmt.Println("Pilih Menu (1-4): ")
	var selectMenuBill string
	fmt.Scanln(&selectMenuBill)
	switch selectMenuBill {
	case "1":
		b.insertFormBill()
	case "2":
		b.showListFormBill()
	case "3":
		b.getByIdFormBill()
	case "4":
		return
	}
}

func (b *BillController) insertFormBill() {
	var bill model.Bill
	fmt.Println("Inputkan bill Id")
	fmt.Scanln(&bill.Id)
	bill.BillDate = time.Now()
	bill.EntryDate = time.Now()
	bill.FinishDate = bill.EntryDate.AddDate(0, 0, 2)
	fmt.Println("Inputkan Employee Id")
	fmt.Scanln(&bill.EmployeeId.Id)
	fmt.Println("Inputkan Customer Id ")
	fmt.Scanln(&bill.CustomerId.Id)
	//perulangan untuk berapa kali bill detailnya

	var numBillDetails int
	fmt.Println("Inputkan jumlah Bill Details:")
	fmt.Scanln(&numBillDetails)

	bill.BillDetails = make([]model.BillDetail, numBillDetails)

	for i := 0; i < numBillDetails; i++ {
		var billDetail model.BillDetail
		fmt.Printf("Inputkan Id untuk Bill Detail #%d: ", i+1)
		fmt.Scanln(&billDetail.Id)
		fmt.Printf("Inputkan Product Id untuk Bill Detail #%d: ", i+1)
		fmt.Scanln(&billDetail.Product.Id)
		fmt.Printf("Inputkan Product Price untuk Bill Detail #%d: ", i+1)
		fmt.Scanln(&billDetail.ProductPrice)
		fmt.Printf("Inputkan Qty untuk Bill Detail #%d: ", i+1)
		fmt.Scanln(&billDetail.Qty)

		// Tambahkan Bill Detail ke dalam slice Bill Details
		bill.BillDetails[i] = billDetail
	}
	err := b.billUC.CreateNew(bill)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data Berhasil ditambah")
}
func (b *BillController) getByIdFormBill() {
	var bill model.Bill
	fmt.Println("Inputkan Id")
	fmt.Scanln(&bill.Id)
	// Panggil use case untuk mencari Bill berdasarkan ID
	bill, err := b.billUC.FindById(bill.Id)
	if err != nil {
		fmt.Printf("Gagal mencari Bill: %v\n", err)
		return
	}

	// Tampilkan data Bill yang ditemukan
	fmt.Println("Data Bill:")
	fmt.Printf("ID: %s\n", bill.Id)
	fmt.Printf("Bill Date: %s\n", bill.BillDate.Truncate(24*time.Hour).Format("02-01-2006"))
	fmt.Printf("Entry Date: %s\n", bill.EntryDate.Truncate(24*time.Hour).Format("02-01-2006"))
	fmt.Printf("Finish Date: %s\n", bill.FinishDate.Truncate(24*time.Hour).Format("02-01-2006"))
	fmt.Printf("Employee Name: %s\n", bill.EmployeeId.Name)
	fmt.Printf("Customer Name: %s\n", bill.CustomerId.Name)

	// Tampilkan detail Bill
	fmt.Println("Bill Details:")
	for i, detail := range bill.BillDetails {
		fmt.Printf("Detail %d:\n", i+1)
		fmt.Printf("  ID: %s\n", detail.Id)
		fmt.Printf("  Product Name: %s\n", detail.Product.Name)
		fmt.Printf("  Product Price: %d\n", detail.ProductPrice)
		fmt.Printf("  Qty: %d\n", detail.Qty)
	}

}

func (b *BillController) showListFormBill() {
	bills, err := b.billUC.FindAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(bills) == 0 {
		fmt.Println("Product is empty")
		return
	}
	for _, bill := range bills {
		fmt.Printf("Bill ID: %s\n", bill.Id)
		fmt.Printf("Bill Date: %s\n", bill.BillDate)
		fmt.Printf("Entry Date: %s\n", bill.EntryDate)
		fmt.Printf("Finish Date: %s\n", bill.FinishDate)
		fmt.Printf("Employee Name: %s\n", bill.EmployeeId.Name)
		fmt.Printf("Customer Name: %s\n", bill.CustomerId.Name)

		// Iterasi melalui BillDetails
		for _, detail := range bill.BillDetails {
			fmt.Printf("  Detail ID: %s\n", detail.Id)
			fmt.Printf("  Product ID: %s\n", detail.Product.Id)
			fmt.Printf("  Product Price: %d\n", detail.ProductPrice)
			fmt.Printf("  Qty: %d\n", detail.Qty)
			fmt.Println()
		}

		fmt.Println()
	}

}

func NewBillController(billUC usecase.BillUseCase) *BillController {
	return &BillController{billUC: billUC}
}
