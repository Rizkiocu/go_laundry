package controller

import (
	"fmt"
	"go_laundry/model"
	"go_laundry/usecase"
)

type ProductController struct {
	productUC usecase.ProductUseCase
}

func (p *ProductController) ProductMenuForm() {
	fmt.Print(`
	|		+++++ Master Product +++++	|
	| 1. Tambah Data				|
	| 2. Lihat Data					|
	| 3. Update Data				|
	| 4. Hapus Data					|
	| 5. Cari Data Berdasarkan ID	|
	| 6. Cari Data Berdasarkan Name	|
	| 7. Keluar                     |
	`)
	fmt.Println("Pilih Menu (1-7): ")
	var selectMenuProduct string
	fmt.Scanln(&selectMenuProduct)
	switch selectMenuProduct {
	case "1":
		p.insertFormProduct()
	case "2":
		p.showListFormProduct()
	case "3":
		p.updateFormProduct()
	case "4":
		p.deleteFormProduct()
	case "5":
		p.getByIdFormProduct()
	case "6":
		p.getByNameFormProduct()
	case "7":
		return
	}
}

func (p *ProductController) insertFormProduct() {
	var product model.Product
	fmt.Println("Inputkan Id")
	fmt.Scanln(&product.Id)
	fmt.Println("Inputkan Name")
	fmt.Scanln(&product.Name)
	fmt.Println("Inputkan Price")
	fmt.Scanln(&product.Price)
	fmt.Println("Inputkan id Uom")
	fmt.Scanln(&product.Uom.Id)
	err := p.productUC.CreateNew(product)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data Berhasil ditambah")
}

func (p *ProductController) showListFormProduct() {
	products, err := p.productUC.FindAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(products) == 0 {
		fmt.Println("Product is empty")
		return
	}

	for _, product := range products {
		fmt.Printf("ID: %s, Name: %s, Price: %d, Id: %s, Name Uom : %s\n", product.Id, product.Name, product.Price, product.Uom.Id, product.Uom.Name)
	}
}

func (p *ProductController) updateFormProduct() {
	var product model.Product
	fmt.Println("Inputkan Id")
	fmt.Scanln(&product.Id)
	fmt.Println("Inputkan Name")
	fmt.Scanln(&product.Name)
	fmt.Println("Inputkan Price")
	fmt.Scanln(&product.Price)
	fmt.Println("Inputkan Id Uom")
	fmt.Scanln(&product.Uom.Id)
	err := p.productUC.Update(product)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data Berhasil diUpdate")
}

func (p *ProductController) deleteFormProduct() {
	var product model.Product
	fmt.Println("Inputkan Id")
	fmt.Scanln(&product.Id)
	err := p.productUC.Delete(product.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data Berhasil dihapus!")
}

func (p *ProductController) getByIdFormProduct() {
	var product model.Product
	fmt.Println("Inputkan Id")
	fmt.Scanln(&product.Id)
	products, err := p.productUC.FindById(product.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ID: %s, Name: %s, Price: %d, Uom Id: %s, Name Uom : %s\n", products.Id, products.Name, products.Price, products.Uom.Id, products.Uom.Name)

}
func (p *ProductController) getByNameFormProduct() {
	var product model.Product
	fmt.Println("Inputkan Name")
	fmt.Scanln(&product.Name)
	products, err := p.productUC.GetByName(product.Name)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, product := range products {
		fmt.Printf("ID: %s, Name: %s, Price: %d, Uom Id: %s, Name Uom: %s\n", product.Id, product.Name, product.Price, product.Uom.Id, product.Uom.Name)
	}

}

func NewProductController(productUC usecase.ProductUseCase) *ProductController {
	return &ProductController{productUC: productUC}
}
