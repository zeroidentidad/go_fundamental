package product

import "database/sql"

type Repository interface {
	GetProductById(productId int) (*Product, error)
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *getUpdateProductRequest) (int64, error)
	DeleteProduct(params *getDeleteProductRequest) (int64, error)
	GetBestSellers() ([]*ProductTop, error)
	GetTotalVentas() (float64, error)
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

func (r *repository) InsertProduct(params *getAddProductRequest) (int64, error) {
	const sql = `INSERT INTO products (product_code, product_name, category, description, list_price, standard_cost) VALUES(?,?,?,?,?,?)`

	insert, err := r.db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	result, _err := insert.Exec(params.ProductCode, params.ProductName, params.Category, params.Description, params.ListPrice, params.StandarCost)

	if _err != nil {
		panic(_err)
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func (r *repository) UpdateProduct(params *getUpdateProductRequest) (int64, error) {
	const sql = `UPDATE products SET product_code=?, product_name=?, category=?, 
				description=?, list_price=?, standard_cost=? WHERE id=?`

	update, err := r.db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	_, _err := update.Exec(params.ProductCode, params.ProductName, params.Category, params.Description, params.ListPrice, params.StandarCost, params.ID)

	if _err != nil {
		panic(_err)
	}

	return params.ID, nil
}

func (r *repository) DeleteProduct(params *getDeleteProductRequest) (int64, error) {
	const sql = `DELETE FROM products WHERE id=?`

	delete, err := r.db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	result, _err := delete.Exec(params.ProductID)

	if _err != nil {
		panic(_err)
	}

	rows, _ := result.RowsAffected()

	return rows, nil
}

func (r *repository) GetBestSellers() ([]*ProductTop, error) {
	const sql = `SELECT  od.product_id, p.product_name, SUM(od.quantity * od.unit_price) vendido
				FROM  order_details od INNER JOIN products p
				ON od.product_id=p.id
				GROUP BY od.product_id ORDER BY vendido DESC LIMIT 10`

	results, err := r.db.Query(sql)

	if err != nil {
		panic(err)
	}

	var productsTop []*ProductTop
	for results.Next() {
		productTop := &ProductTop{}
		err = results.Scan(&productTop.ID, &productTop.ProductName, &productTop.Vendidos)

		if err != nil {
			panic(err)
		}

		productsTop = append(productsTop, productTop)
	}

	return productsTop, nil
}

func (r *repository) GetTotalVentas() (float64, error) {
	const sql = `SELECT SUM(od.quantity * od.unit_price) vendido FROM order_details od`

	var total float64

	row := r.db.QueryRow(sql)
	err := row.Scan(&total)

	if err != nil {
		panic(err)
	}

	return total, nil
}
