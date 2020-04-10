package service

import "shop/model"

type (
	// ProductService :nodoc:
	ProductService interface {
		List() []model.Product
		FindProductByID(id int64) *model.Product
	}

	// OrderService :nodoc:
	OrderService interface {
		CreateOrder(productIDs []int64) *model.Order
		FindOrderByID(id int64) *model.Order
	}

	// PaymentService :nodco:
	PaymentService interface {
		// create payment from order
		CreatePayment(orderID int64) *model.Payment
		PayBills(payment *model.Payment) bool
	}
)
