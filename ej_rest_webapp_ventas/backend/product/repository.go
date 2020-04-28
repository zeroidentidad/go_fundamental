package product

import "database/sql"

type Repository interface {
	GetProductById(product int) (*Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(dbconn *sql.DB) Repository {
	return &repository{db: dbconn}
}

func (r *repository) GetProductById(productId int) (*Product, error) {
	const sql = `SELECT id, product_code, product_name, COALESCE(description,''),
				standard_cost, list_price, category 
				FROM products WHERE id=?`

	row := r.db.QueryRow(sql, productId)
	product := &Product{}

	err := row.Scan(&product.ID, &product.ProductCode, &product.ProductName, &product.Description, &product.StandarCost, &product.ListPrice, &product.Category)

	if err != nil {
		panic(err)
	}

	return product, err
}
