package order

import (
	"database/sql"

	"github.com/zeroidentidad/backend/helper"
)

type Repository interface {
	GetOrderById(param *getOrderByIdRequest) (*OrderItem, error)
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
