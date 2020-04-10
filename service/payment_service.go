package service

import (
	"shop/model"
	"time"
)

type (
	paymentService struct {
		payments     []*model.Payment
		orderService OrderService
	}
)

// NewPaymentService :nodoc:
func NewPaymentService(os OrderService) PaymentService {
	return &paymentService{
		orderService: os,
	}
}

func (ps *paymentService) CreatePayment(orderID int64) *model.Payment {
	return &model.Payment{
		ID:      time.Now().UnixNano(),
		OrderID: orderID,
		Status:  model.PaymentStatusPending,
	}
}

func (ps *paymentService) PayBills(payment *model.Payment) (ok bool) {
	if payment == nil {
		return
	}

	pm := ps.findPaymentByID(payment.ID)
	if pm == nil {
		return
	}

	if order := ps.orderService.FindOrderByID(payment.OrderID); order == nil {
		return
	}

	pm.Status = model.PaymentStatusPaid

	return
}

func (ps *paymentService) findPaymentByID(id int64) *model.Payment {
	for _, pm := range ps.payments {
		if pm.ID == id {
			return pm
		}
	}

	return nil
}
