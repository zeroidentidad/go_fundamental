package order

import (
	"database/sql"
	"fmt"

	"github.com/zeroidentidad/backend/helper"
)

type Repository interface {
	GetOrderById(param *getOrderByIdRequest) (*OrderItem, error)
	GetOrders(params *getOrdersRequest) ([]*OrderItem, error)
	GetTotalOrders(params *getOrdersRequest) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetOrderById(param *getOrderByIdRequest) (*OrderItem, error) {
	const sql = `SELECT o.id, o.customer_id, o.order_date, o.status_id, os.status_name, 
				CONCAT(c.first_name,' ',c.last_name) as costumer_name,
				c.company, c.address, c.business_phone, c.city
				FROM orders o
				INNER JOIN orders_status os ON o.status_id=os.id 
				INNER JOIN customers c ON o.customer_id=c.id 
				WHERE o.id=?`

	order := &OrderItem{}

	row := r.db.QueryRow(sql, param.orderId)
	err := row.Scan(&order.ID, &order.CustomerID, &order.OrderDate, &order.StatusId, &order.StatusName, &order.Customer, &order.Company, &order.Address, &order.Phone, &order.City)
	helper.Catch(err)

	orderDetail, err := GetOrderDetail(r, &param.orderId)
	helper.Catch(err)

	order.Data = orderDetail

	return order, nil
}

func GetOrderDetail(r *repository, orderId *int64) ([]*OrderDetailItem, error) {
	const sql = `SELECT od.order_id, od.id, od.quantity, od.unit_price, 
				p.product_name, od.product_id FROM order_details od
				INNER JOIN products p ON od.product_id=p.id
				WHERE od.order_id=?`

	results, err := r.db.Query(sql, orderId)
	helper.Catch(err)

	var orders []*OrderDetailItem
	for results.Next() {
		orderDetailItem := &OrderDetailItem{}
		err := results.Scan(&orderDetailItem.OrderId, &orderDetailItem.ID, &orderDetailItem.Quantity, &orderDetailItem.UnitPrice, &orderDetailItem.ProductName, &orderDetailItem.ProductId)
		helper.Catch(err)

		orders = append(orders, orderDetailItem)
	}

	return orders, nil
}

func (r *repository) GetOrders(params *getOrdersRequest) ([]*OrderItem, error) {
	filter := Filter(params)

	var sql = `SELECT o.id, o.customer_id, o.order_date, o.status_id,
			os.status_name, CONCAT(c.first_name,' ',c.last_name) AS customer_name 
			FROM orders o
			INNER JOIN orders_status os ON o.status_id=os.id 
			INNER JOIN customers c ON o.customer_id=c.id
			WHERE 1=1 `

	sql = sql + *filter + "LIMIT ? OFFSET ?"

	results, err := r.db.Query(sql, params.Limit, params.Offset)
	helper.Catch(err)

	var orders []*OrderItem

	for results.Next() {
		order := &OrderItem{}

		err := results.Scan(&order.ID, &order.CustomerID, &order.OrderDate, &order.StatusId, &order.StatusName, &order.Customer)
		helper.Catch(err)

		orderDetail, err := GetOrderDetail(r, &order.ID)
		helper.Catch(err)

		order.Data = orderDetail

		orders = append(orders, order)
	}

	return orders, err
}

func (r *repository) GetTotalOrders(params *getOrdersRequest) (int64, error) {
	filter := Filter(params)

	var sql = "SELECT COUNT(*) FROM orders o WHERE 1=1 " + *filter
	row := r.db.QueryRow(sql)

	var total int64

	err := row.Scan(&total)
	helper.Catch(err)

	return total, nil
}

func Filter(params *getOrdersRequest) *string {
	var filter string

	if params.Status != nil {
		filter += fmt.Sprintf(" AND o.status_id = %v ", params.Status.(float64))
	}

	if params.DateFrom != nil && params.DateTo == nil {
		filter += fmt.Sprintf(" AND o.order_date >= '%v'", params.DateFrom.(string))
	}

	if params.DateFrom == nil && params.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date <= '%v'", params.DateTo.(string))
	}

	if params.DateFrom != nil && params.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date BETWEEN '%v' AND '%v'", params.DateFrom.(string), params.DateTo.(string))
	}

	return &filter
}
