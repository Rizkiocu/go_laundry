package controller

import (
	"fmt"
	"go_laundry/model"
	"go_laundry/usecase"
)

type UomController struct {
	uomUC usecase.UomUseCase
}

func (u *UomController) UomMenuForm() {
	fmt.Print(`
	|		+++++ Master UOM +++++	|
	| 1. Tambah Data				|
	| 2. Lihat Data					|
	| 3. Update Data				|
	| 4. Hapus Data					|
	| 5. Cari Data Berdasarkan ID	|
	| 6. Keluar                     |
	`)
	fmt.Println("Pilih Menu (1-6): ")
	var selectMenuUom string
	fmt.Scanln(&selectMenuUom)
	switch selectMenuUom {
	case "1":
		u.insertFormUom()
	case "2":
		u.showListFormUom()
	case "3":
		u.updateFormUom()
	case "4":
		u.deleteFormUom()
	case "5":
		u.getByIdFormUom()
	case "6":
		return
	}
}

func (u *UomController) insertFormUom() {
	var uom model.Uom
	fmt.Println("Inputkan Id")
	fmt.Scanln(&uom.Id)
	fmt.Println("Inputkan Name")
	fmt.Scanln(&uom.Name)
	err := u.uomUC.CreateNew(uom)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (u *UomController) showListFormUom() {
	uoms, err := u.uomUC.FindAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(uoms) == 0 {
		fmt.Println("Uom is empty")
		return
	}

	for _, uom := range uoms {
		fmt.Printf("ID: %s, Name: %s\n", uom.Id, uom.Name)
	}
}

func (u *UomController) updateFormUom() {
	var uom model.Uom
	fmt.Println("Inputkan Id")
	fmt.Scanln(&uom.Id)
	fmt.Println("Inputkan Name")
	fmt.Scanln(&uom.Name)
	err := u.uomUC.Update(uom)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data Berhasil diUpdate")
}

func (u *UomController) deleteFormUom() {
	var uom model.Uom
	fmt.Println("Inputkan Id")
	fmt.Scanln(&uom.Id)
	err := u.uomUC.Delete(uom.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data Berhasil dihapus!")
}

func (u *UomController) getByIdFormUom() {
	var uom model.Uom
	fmt.Println("Inputkan Id")
	fmt.Scanln(&uom.Id)
	uoms, err := u.uomUC.FindById(uom.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data Yang dicari : ", uoms.Id, uoms.Name)
}

func NewUomController(uomUC usecase.UomUseCase) *UomController {
	return &UomController{uomUC: uomUC}
}
