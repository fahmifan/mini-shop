package service

import (
	"context"
	"shop/eventbus"
	"shop/model"
	"sync"
	"time"

	"github.com/mustafaturan/bus"
	log "github.com/sirupsen/logrus"
)

type (
	paymentService struct {
		payments     []*model.Payment
		orderService OrderService
		bbus         *bus.Bus
	}
)

// NewPaymentService :nodoc:
func NewPaymentService(os OrderService, bbus *bus.Bus) PaymentService {

	ps := &paymentService{
		orderService: os,
		bbus:         bbus,
	}

	ps.bbus.RegisterHandler("payment-service", &bus.Handler{
		Matcher: eventbus.OrderAll,
		Handle:  ps.Subscriber,
	})

	return ps
}

func (ps *paymentService) CreatePayment(orderID int64) *model.Payment {
	pm := &model.Payment{
		ID:      time.Now().UnixNano(),
		OrderID: orderID,
		Status:  model.PaymentStatusPending,
	}

	mu := sync.Mutex{}
	mu.Lock()
	ps.payments = append(ps.payments, pm)
	mu.Unlock()
	return pm
}

func (ps *paymentService) PayBills(payment *model.Payment) (ok bool) {
	if payment == nil {
		return
	}

	pm := ps.findPaymentByID(payment.ID)
	if pm == nil {
		log.Error("payment not found not found")
		return
	}

	// if order := ps.orderService.FindOrderByID(pm.OrderID); order == nil {
	// 	log.Error("order not found")
	// 	return
	// }

	pm.Status = model.PaymentStatusPaid
	ok = true

	payment = pm

	go func() {
		_, err := ps.bbus.Emit(context.Background(), eventbus.PaymentPaid, *pm)
		if err != nil {
			log.Error(err)
		}
	}()

	return
}

func (ps *paymentService) Subscriber(event *bus.Event) {
	go func() {
		switch event.Topic {
		case eventbus.OrderCreated:
			order, ok := event.Data.(model.Order)
			if !ok {
				return
			}

			payment := ps.CreatePayment(order.ID)
			log.Info("create payment", payment)
		}
	}()
}

func (ps *paymentService) List() []*model.Payment {
	return ps.payments
}

func (ps *paymentService) findPaymentByID(id int64) *model.Payment {
	for _, pm := range ps.payments {
		if pm.ID == id {
			return pm
		}
	}

	return nil
}
