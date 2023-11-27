package delivery

import (
	"fmt"
	"go_laundry/config"
	controller "go_laundry/delivery/controller/cli"
	"go_laundry/manager"
	"go_laundry/repository"
	"go_laundry/usecase"
	"os"
)

type Console struct {
	// Tempat untuk menaruh semua usecase yg dibutuhkan
	uomUC      usecase.UomUseCase
	productUC  usecase.ProductUseCase
	customerUC usecase.CustomerUseCase
	employeeUC usecase.EmployeeUseCase
	billUC     usecase.BillUseCase
}

func (c *Console) showMainMenu() {
	fmt.Println(`
	|+++++ Enigma Laundry Menu +++++|
	| 1. Master UOM                 |
	| 2. Master Product             |
	| 3. Master Customer            |
	| 4. Master Employee             |
	| 5. Transaksi                  |
	| 6. Keluar                     |
	`)
	fmt.Println("Pilih Menu (1-6): ")
}

func (c *Console) Run() {
	for {
		c.showMainMenu()
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			controller.NewUomController(c.uomUC).UomMenuForm()
		case "2":
			controller.NewProductController(c.productUC).ProductMenuForm()
		case "3":
			controller.NewCustomerController(c.customerUC).CustomerMenuForm()
		case "4":
			controller.NewEmployeeController(c.employeeUC).EmployeeMenuForm()
		case "5":
			controller.NewBillController(c.billUC).BillMenuForm()
		case "6":
			os.Exit(0)
		}
	}
}

func NewConsole() *Console {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	con, err := manager.NewInfraManager(cfg)
	if err != nil {
		fmt.Println(err)
	}
	db := con.Conn()
	//repositort
	uomRepo := repository.NewUomRepository(db)
	productRepo := repository.NewProductRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	billRepo := repository.NewBillRepository(db)

	//usecase
	uomUC := usecase.NewUomUseCase(uomRepo)
	productUC := usecase.NewProductUseCase(productRepo, uomUC)
	customerUC := usecase.NewCustomerUseCase(customerRepo)
	employeeUC := usecase.NewEmployeeUseCase(employeeRepo)
	billUC := usecase.NewBillUseCase(billRepo, employeeUC, customerUC, productUC)
	return &Console{
		uomUC:      uomUC,
		productUC:  productUC,
		customerUC: customerUC,
		employeeUC: employeeUC,
		billUC:     billUC,
	}
}
