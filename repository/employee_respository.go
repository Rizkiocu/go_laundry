package repository

import (
	"database/sql"
	"fmt"
	"go_laundry/model"
	"go_laundry/model/dto"
	"math"
)

type EmployeeRepository interface {
	Save(employee model.Employee) error
	FindById(id string) (model.Employee, error)
	FindAll() ([]model.Employee, error)
	FindByPhone(phone_number string) ([]model.Employee, error)
	Update(employee model.Employee) error
	DeleteById(id string) error
	Paging(payload dto.PageRequest) ([]model.Employee, dto.Paging, error)
}
type employeeRespository struct {
	db *sql.DB
}

// Paging implements EmployeeRepository.
func (e *employeeRespository) Paging(payload dto.PageRequest) ([]model.Employee, dto.Paging, error) {
	if payload.Page <= 0 {
		payload.Page = 1
	}
	q := `SELECT id, name, phone_number, address FROM employee LIMIT $2 OFFSET $1`
	rows, err := e.db.Query(q, (payload.Page-1)*payload.Size, payload.Size)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Phone_number, &employee.Address)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		employees = append(employees, employee)
	}
	var count int
	row := e.db.QueryRow("SELECT COUNT(id) FROM employee")
	if err := row.Scan(&count); err != nil {
		return nil, dto.Paging{}, err
	}

	paging := dto.Paging{
		Page:       payload.Page,
		Size:       payload.Size,
		TotalRows:  count,
		TotalPages: int(math.Ceil(float64(count) / float64(payload.Size))), // (totalrow / size)
	}

	return employees, paging, nil

}

// DeleteById implements CustomerRepository.
func (e *employeeRespository) DeleteById(id string) error {
	_, err := e.db.Exec("DELETE FROM employee WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements CustomerRepository.
func (e *employeeRespository) FindAll() ([]model.Employee, error) {
	rows, err := e.db.Query("SELECT id,name,phone_number,address name FROM employee")
	if err != nil {
		return nil, err
	}

	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Phone_number, &employee.Address)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

// FindById implements CustomerRepository.
func (e *employeeRespository) FindById(id string) (model.Employee, error) {
	row := e.db.QueryRow("SELECT id, name, phone_number, address FROM employee WHERE id=$1", id)
	employee := model.Employee{}
	err := row.Scan(&employee.Id, &employee.Name, &employee.Phone_number, &employee.Address)
	if err != nil {
		fmt.Println("id not found", employee, err)
	}

	return employee, nil
}

// FindByPhone implements CustomerRepository.
func (e *employeeRespository) FindByPhone(phone_number string) ([]model.Employee, error) {
	rows, err := e.db.Query(`SELECT id, name, phone_number, address FROM employee WHERE phone_number ILIKE $1`, "%"+phone_number+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Jangan lupa menutup rows setelah digunakan
	var employees []model.Employee
	for rows.Next() {
		employee := model.Employee{}
		err := rows.Scan(
			&employee.Id,
			&employee.Name,
			&employee.Phone_number,
			&employee.Address,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee) // Menambahkan pelanggan ke slice
	}
	return employees, nil
}

// Save implements CustomerRepository.
func (e *employeeRespository) Save(employee model.Employee) error {
	_, err := e.db.Exec("INSERT INTO employee (id, name, phone_number, address) VALUES ($1, $2, $3, $4)", employee.Id, employee.Name, employee.Phone_number, employee.Address)
	if err != nil {
		return err
	}

	return nil
}

// Update implements CustomerRepository.
func (e *employeeRespository) Update(employee model.Employee) error {
	_, err := e.db.Exec("UPDATE employee SET name=$2, phone_number=$3, address=$4  WHERE id=$1", employee.Id, employee.Name, employee.Phone_number, employee.Address)
	if err != nil {
		return err
	}
	return nil
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRespository{db: db}
}
