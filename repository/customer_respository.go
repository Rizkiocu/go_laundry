package repository

import (
	"database/sql"
	"fmt"
	"go_laundry/model"
	"go_laundry/model/dto"
	"math"
)

type CustomerRepository interface {
	Save(customer model.Customer) error
	FindById(id string) (model.Customer, error)
	FindAll() ([]model.Customer, error)
	FindByPhone(phone_number string) ([]model.Customer, error)
	Update(customer model.Customer) error
	DeleteById(id string) error
	Paging(payload dto.PageRequest) ([]model.Customer, dto.Paging, error)
}
type customerRespository struct {
	db *sql.DB
}

// Paging implements CustomerRepository.
func (c *customerRespository) Paging(payload dto.PageRequest) ([]model.Customer, dto.Paging, error) {
	if payload.Page <= 0 {
		payload.Page = 1
	}
	q := `SELECT id, name, phone_number, address FROM customer LIMIT $2 OFFSET $1`
	rows, err := c.db.Query(q, (payload.Page-1)*payload.Size, payload.Size)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Phone_number, &customer.Address)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		customers = append(customers, customer)
	}
	var count int
	row := c.db.QueryRow("SELECT COUNT(id) FROM customer")
	if err := row.Scan(&count); err != nil {
		return nil, dto.Paging{}, err
	}

	paging := dto.Paging{
		Page:       payload.Page,
		Size:       payload.Size,
		TotalRows:  count,
		TotalPages: int(math.Ceil(float64(count) / float64(payload.Size))), // (totalrow / size)
	}

	return customers, paging, nil

}

// DeleteById implements CustomerRepository.
func (c *customerRespository) DeleteById(id string) error {
	_, err := c.db.Exec("DELETE FROM customer WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements CustomerRepository.
func (c *customerRespository) FindAll() ([]model.Customer, error) {
	rows, err := c.db.Query("SELECT id,name,phone_number,address name FROM customer")
	if err != nil {
		return nil, err
	}

	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Phone_number, &customer.Address)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

// FindById implements CustomerRepository.
func (c *customerRespository) FindById(id string) (model.Customer, error) {
	row := c.db.QueryRow("SELECT id, name, phone_number, address FROM customer WHERE id=$1", id)
	customer := model.Customer{}
	err := row.Scan(&customer.Id, &customer.Name, &customer.Phone_number, &customer.Address)
	if err != nil {
		fmt.Println("id yg dicari tidak ada", customer, err)
	}

	return customer, nil
}

// FindByPhone implements CustomerRepository.
func (c *customerRespository) FindByPhone(phone_number string) ([]model.Customer, error) {
	rows, err := c.db.Query(`SELECT id, name, phone_number, address FROM customer WHERE phone_number ILIKE $1`, "%"+phone_number+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Jangan lupa menutup rows setelah digunakan
	var customers []model.Customer
	for rows.Next() {
		customer := model.Customer{}
		err := rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.Phone_number,
			&customer.Address,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer) // Menambahkan pelanggan ke slice
	}
	return customers, nil
}

// Save implements CustomerRepository.
func (c *customerRespository) Save(customer model.Customer) error {
	_, err := c.db.Exec("INSERT INTO customer (id, name, phone_number, address) VALUES ($1, $2, $3, $4)", customer.Id, customer.Name, customer.Phone_number, customer.Address)
	if err != nil {
		return err
	}

	return nil
}

// Update implements CustomerRepository.
func (c *customerRespository) Update(customer model.Customer) error {
	_, err := c.db.Exec("UPDATE customer SET name=$2, phone_number=$3, address=$4  WHERE id=$1", customer.Id, customer.Name, customer.Phone_number, customer.Address)
	if err != nil {
		return err
	}
	return nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRespository{db: db}
}
