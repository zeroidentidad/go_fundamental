package employee

import (
	"database/sql"

	"github.com/zeroidentidad/backend/helper"
)

type Repository interface {
	GetEmployees(params *getEmployeesRequest) ([]*Employee, error)
	GetTotalEmployees() (int64, error)
	GetEmployeeById(param *getEmployeeByIDRequest) (*Employee, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetEmployees(params *getEmployeesRequest) ([]*Employee, error) {
	const sql = `SELECT id, first_name, last_name, company, 
				email_address, job_title, business_phone, 
				home_phone, COALESCE(mobile_phone, ''), fax_number, address
				FROM employees
				LIMIT ? OFFSET ?`

	results, err := r.db.Query(sql, params.Limit, params.Offset)
	helper.Catch(err)

	var employees []*Employee

	for results.Next() {
		employee := &Employee{}
		err := results.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Company, &employee.EmailAddress, &employee.JobTitle, &employee.BusinessPhone, &employee.HomePhone, &employee.MobilePhone, &employee.FaxNumber, &employee.Address)
		helper.Catch(err)

		employees = append(employees, employee)
	}

	return employees, nil
}

func (r *repository) GetTotalEmployees() (int64, error) {
	const sql = `select COUNT(*) from employees`
	var total int64

	row := r.db.QueryRow(sql)

	err := row.Scan(&total)
	helper.Catch(err)

	return total, nil
}

func (r *repository) GetEmployeeById(param *getEmployeeByIDRequest) (*Employee, error) {
	const sql = `SELECT id, first_name, last_name, company, 
				email_address, job_title, business_phone, 
				home_phone, COALESCE(mobile_phone, ''), fax_number, address
				FROM employees
				WHERE id=?`

	row := r.db.QueryRow(sql, param.EmployeeID)
	employee := &Employee{}

	err := row.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Company, &employee.EmailAddress, &employee.JobTitle, &employee.BusinessPhone, &employee.HomePhone, &employee.MobilePhone, &employee.FaxNumber, &employee.Address)
	helper.Catch(err)

	return employee, nil
}
