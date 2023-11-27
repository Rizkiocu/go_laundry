package repository

import (
	"database/sql"
	"fmt"
	"go_laundry/model"
)

type ProductRepository interface {
	Save(product model.Product) error
	FindById(id string) (model.Product, error)
	FindAll() ([]model.Product, error)
	FindByName(name string) ([]model.Product, error)
	Update(product model.Product) error
	DeleteById(id string) error
}

type productRepository struct {
	db *sql.DB
}

// DeleteById implements ProductRepository.
func (p *productRepository) DeleteById(id string) error {
	_, err := p.db.Exec("DELETE FROM product WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements ProductRepository.
func (p *productRepository) FindAll() ([]model.Product, error) {
	rows, err := p.db.Query("SELECT p.id, p.name, p.price, u.id, u.name FROM product p JOIN uom u ON u.id = p.uom_id")
	if err != nil {
		return nil, err
	}

	var products []model.Product
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Uom.Id, &product.Uom.Name)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// FindById implements ProductRepository.
func (p *productRepository) FindById(id string) (model.Product, error) {
	row := p.db.QueryRow("SELECT p.id,p.name, p.price, u.id, u.name FROM product p JOIN uom u ON u.id = p.uom_id WHERE p.id=$1", id)
	product := model.Product{}
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Uom.Id, &product.Uom.Name)
	if err != nil {
		fmt.Println("id yg dicari tidak ada", product, err)
	}

	return product, nil
}

func (p *productRepository) FindByName(name string) ([]model.Product, error) {
	rows, err := p.db.Query(`SELECT p.id, p.name, p.price, u.id, u.name FROM product p
	JOIN uom u ON u.id = p.uom_id WHERE name ILIKE '$1'`, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	var products []model.Product
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Uom.Id,
			&product.Uom.Name,
		)
		if err != nil {
			return nil, err
		}
	}
	return products, nil
}

// Save implements ProductRepository.
func (p *productRepository) Save(product model.Product) error {
	_, err := p.db.Exec("INSERT INTO product (id, name, price, uom_id) VALUES ($1, $2, $3, $4)", product.Id, product.Name, product.Price, product.Uom.Id)
	if err != nil {
		return err
	}
	return nil
}

// Update implements ProductRepository.
func (p *productRepository) Update(product model.Product) error {
	_, err := p.db.Exec("UPDATE product SET name=$2, price=$3, uom_id=$4 WHERE id=$1", product.Id, product.Name, product.Price, product.Uom.Id)
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}
