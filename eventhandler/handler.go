package eventhandler

import (
	"shop/eventbus"
	"shop/model"
	"shop/service"

	"github.com/mustafaturan/bus"
	log "github.com/sirupsen/logrus"
)

// EventHandler :nodoc:
type EventHandler struct {
	PaymentService service.PaymentService
}

// HandleOrder :nodoc:
func (e *EventHandler) HandleOrder(event *bus.Event) {
	switch event.Topic {
	case eventbus.OrderCreated:
		log.Infof("recieved event %v", event.ID)
		order, ok := event.Data.(model.Order)
		if !ok {
			return
		}

		payment := e.PaymentService.CreatePayment(order.ID)
		// TODO: sent payment bill notification
		log.Info("create payment", payment)
	}
}
