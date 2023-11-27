package repository

import (
	"database/sql"
	"go_laundry/model"
)

type BillRepository interface {
	Save(bill model.Bill) error
	FindById(id string) (model.Bill, error)
	FindAll() ([]model.Bill, error)
}
type billRespository struct {
	db *sql.DB
}

// FindAll implements BillRepository.
func (b *billRespository) FindAll() ([]model.Bill, error) {
	// Eksekusi query
	rows, err := b.db.Query(`
	 SELECT b.id, b.bill_date, b.entry_date, b.finish_date, b.employee_id, b.customer_id, e.name as employee_name, c.name as customer_name
	 FROM bill b
	 JOIN employee e ON b.employee_id = e.id
	 JOIN customer c ON b.customer_id = c.id
 `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Membaca hasil query ke slice Bills
	var bills []model.Bill
	for rows.Next() {
		var bill model.Bill
		err := rows.Scan(
			&bill.Id, &bill.BillDate, &bill.EntryDate, &bill.FinishDate, &bill.EmployeeId.Id, &bill.CustomerId.Id,
			&bill.EmployeeId.Name, &bill.CustomerId.Name,
		)
		if err != nil {
			return nil, err
		}

		// Eksekusi query untuk BillDetail
		detailRows, err := b.db.Query(`SELECT id, product_id, product_price, qty FROM bill_detail WHERE bill_id = $1`, bill.Id)
		if err != nil {
			return nil, err
		}
		defer detailRows.Close()

		// Iterasi hasil query untuk mengisi slice BillDetails
		for detailRows.Next() {
			var detail model.BillDetail
			err := detailRows.Scan(&detail.Id, &detail.Product.Id, &detail.ProductPrice, &detail.Qty)
			if err != nil {
				return nil, err
			}
			bill.BillDetails = append(bill.BillDetails, detail)
		}

		bills = append(bills, bill)
	}

	return bills, nil
}

// FindById implements BillRepository.
func (b *billRespository) FindById(id string) (model.Bill, error) {

	// Eksekusi query
	row := b.db.QueryRow(`SELECT b.id, b.bill_date, b.entry_date, b.finish_date, b.employee_id, b.customer_id,e.name as employee_name, c.name as customer_name FROM bill b JOIN employee e ON b.employee_id = e.id JOIN customer c ON b.customer_id = c.id WHERE b.id = $1`, id)

	// Membaca hasil query ke struct Bill
	var bill model.Bill
	err := row.Scan(
		&bill.Id, &bill.BillDate, &bill.EntryDate, &bill.FinishDate, &bill.EmployeeId.Id, &bill.CustomerId.Id,
		&bill.EmployeeId.Name, &bill.CustomerId.Name,
	)

	if err != nil {
		return model.Bill{}, err
	}

	// Eksekusi query untuk BillDetail
	rows, err := b.db.Query(`SELECT id, product_id, product_price, qty FROM bill_detail WHERE bill_id = $1`, id)
	if err != nil {
		return model.Bill{}, err
	}
	defer rows.Close()

	// Iterasi hasil query untuk mengisi slice BillDetails
	for rows.Next() {
		var detail model.BillDetail
		err := rows.Scan(&detail.Id, &detail.Product.Id, &detail.ProductPrice, &detail.Qty)
		if err != nil {
			return model.Bill{}, err
		}
		bill.BillDetails = append(bill.BillDetails, detail)
	}

	return bill, nil
}

// Save implements BillRepository.
func (b *billRespository) Save(bill model.Bill) error {
	// Mulai transaksi
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // Dirollback jika terjadi kesalahan

	// Simpan Bill
	_, err = tx.Exec(`INSERT INTO bill (id, bill_date, entry_date, finish_date, employee_id, customer_id)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, bill.Id, bill.BillDate, bill.EntryDate, bill.FinishDate, bill.EmployeeId.Id, bill.CustomerId.Id)
	if err != nil {
		return err
	}

	// Simpan BillDetails
	for _, detail := range bill.BillDetails {
		_, err = tx.Exec(`
			INSERT INTO bill_detail (id, bill_id, product_id, product_price, qty)
			VALUES ($1, $2, $3, $4, $5)
		`, detail.Id, bill.Id, detail.Product.Id, detail.ProductPrice, detail.Qty)
		if err != nil {
			return err
		}
	}

	// Commit transaksi
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func NewBillRepository(db *sql.DB) BillRepository {
	return &billRespository{db: db}
}
