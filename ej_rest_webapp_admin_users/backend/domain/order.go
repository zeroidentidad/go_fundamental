package domain

import (
	"backend/dto"
	"backend/errs"
)

type Order struct {
	ID         uint        `gorm:"column:id;primaryKey"`
	FirstName  string      `gorm:"column:first_name"`
	LastName   string      `gorm:"column:last_name"`
	Email      string      `gorm:"column:email"`
	UpdatedAt  string      `gorm:"column:updated_at"`
	CreatedAt  string      `gorm:"column:created_at"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID           uint    `gorm:"column:id;primaryKey"`
	OrderID      uint    `gorm:"column:order_id"`
	ProductTitle string  `gorm:"column:product_title"`
	Price        float32 `gorm:"column:price"`
	Quantity     uint    `gorm:"column:quantity"`
}

type OrderStorage interface {
	SelectOrders(int) (*[]Order, int64, *errs.AppError)
}

func (o Order) ToDto() dto.ResponseOrder {
	order_items := make([]dto.ResponseOrderItem, len(o.OrderItems))
	var total float32
	for index, item := range o.OrderItems {
		order_items[index] = dto.ResponseOrderItem{
			ID:           item.ID,
			OrderID:      item.OrderID,
			ProductTitle: item.ProductTitle,
			Price:        item.Price,
			Quantity:     item.Quantity,
		}
		total += item.Price * float32(item.Quantity)
	}
	return dto.ResponseOrder{
		ID:         o.ID,
		Name:       o.FirstName + " " + o.LastName,
		Email:      o.Email,
		Total:      total,
		UpdatedAt:  o.UpdatedAt,
		CreatedAt:  o.CreatedAt,
		OrderItems: order_items,
	}
}
