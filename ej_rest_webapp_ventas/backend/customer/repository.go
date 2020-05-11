package customer

import (
	"database/sql"

	"github.com/zeroidentidad/backend/helper"
)

type Repository interface {
	GetCustomers(params *getCustomersRequest) ([]*Customer, error)
	GetTotalCustomers() (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetTotalCustomers() (int64, error) {
	const sql = `SELECT COUNT(*) FROM customers`
	var total int64

	row := r.db.QueryRow(sql)

	err := row.Scan(&total)
	helper.Catch(err)

	return total, nil
}

func (r *repository) GetCustomers(params *getCustomersRequest) ([]*Customer, error) {
	const sql = `SELECT id, first_name, last_name, address, 
				business_phone, city, company 
				FROM customers 
				LIMIT ? OFFSET ? `

	results, err := r.db.Query(sql, &params.Limit, &params.Offset)
	helper.Catch(err)

	var customers []*Customer
	for results.Next() {
		customer := &Customer{}
		err = results.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Address, &customer.BusinessPhone, &customer.City, &customer.Company)
		helper.Catch(err)

		customers = append(customers, customer)
	}

	return customers, nil
}
