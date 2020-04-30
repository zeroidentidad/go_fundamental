package product

import "database/sql"

type Repository interface {
	GetProductById(productId int) (*Product, error)
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
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

func (r *repository) GetProducts(params *getProductsRequest) ([]*Product, error) {
	const sql = `SELECT id, product_code, product_name, COALESCE(description,''),
				standard_cost, list_price, category 
				FROM products LIMIT ? OFFSET ?`

	results, err := r.db.Query(sql, params.Limit, params.Offset)

	if err != nil {
		panic(err)
	}

	var products []*Product
	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.ProductCode, &product.ProductName, &product.Description, &product.StandarCost, &product.ListPrice, &product.Category)

		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *repository) GetTotalProducts() (int, error) {
	const sql = `SELECT COUNT(*) FROM products`
	var total int

	row := r.db.QueryRow(sql)
	err := row.Scan(&total)

	if err != nil {
		panic(err)
	}

	return total, nil
}
