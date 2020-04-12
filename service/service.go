package service

import (
	"shop/model"

	"github.com/mustafaturan/bus"
)

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
		Subscriber(event *bus.Event)
		List() []*model.Payment
	}

	// NotificationService :nodoc:
	NotificationService interface {
		List() []string
		CreateNotification(msg string)
		Subscriber(event *bus.Event)
	}
)
