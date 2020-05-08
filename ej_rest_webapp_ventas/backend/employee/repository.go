package employee

import (
	"database/sql"

	"github.com/zeroidentidad/backend/helper"
)

type Repository interface {
	GetEmployees(params *getEmployeesRequest) ([]*Employee, error)
	GetTotalEmployees() (int64, error)
	GetEmployeeById(param *getEmployeeByIDRequest) (*Employee, error)
	GetBestEmployee() (*BestEmployee, error)
	InsertEmployee(params *addEmployeeRequest) (int64, error)
	UpdateEmployee(params *updateEmployeeRequest) (int64, error)
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

func (r *repository) GetBestEmployee() (*BestEmployee, error) {
	const sql = `SELECT e.id, COUNT(e.id) as totalVentas, 
				e.first_name, e.last_name
				FROM  orders o INNER JOIN employees e ON o.employee_id=e.id
				GROUP BY o.employee_id
				ORDER BY totalVentas DESC limit 1`

	row := r.db.QueryRow(sql)
	employee := &BestEmployee{}

	err := row.Scan(&employee.ID, &employee.TotalVentas, &employee.LastName, &employee.FirstName)
	helper.Catch(err)

	return employee, nil
}

func (r *repository) InsertEmployee(params *addEmployeeRequest) (int64, error) {
	const sql = `INSERT INTO employees 
				(first_name, last_name, company, address, business_phone,
				email_address, fax_number, home_phone, job_title, mobile_phone)
				VALUES(?,?,?,?,?,?,?,?,?,?)`

	insert, err := r.db.Prepare(sql)
	helper.Catch(err)

	result, _err := insert.Exec(params.FirstName, params.LastName, params.Company, params.Address, params.BusinessPhone, params.EmailAddress, params.FaxNumber, params.HomePhone, params.JobTitle, params.MobilePhone)

	helper.Catch(_err)

	id, _ := result.LastInsertId()

	return id, nil
}

func (r *repository) UpdateEmployee(params *updateEmployeeRequest) (int64, error) {
	const sql = `UPDATE employees SET 
				first_name=?, last_name=?, company=?, address=?, business_phone=?,
				email_address=?, fax_number=?, home_phone=?, job_title=?, mobile_phone=?
				WHERE id=?`

	update, err := r.db.Prepare(sql)
	helper.Catch(err)

	_, _err := update.Exec(params.FirstName, params.LastName, params.Company, params.Address, params.BusinessPhone, params.EmailAddress, params.FaxNumber, params.HomePhone, params.JobTitle, params.MobilePhone, params.ID)

	helper.Catch(_err)

	return params.ID, nil
}
