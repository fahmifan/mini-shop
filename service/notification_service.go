package service

import (
	"fmt"
	"shop/eventbus"
	"shop/model"
	"sync"

	"github.com/mustafaturan/bus"
	"github.com/sirupsen/logrus"
)

type notificationService struct {
	notifs []string
	bbus   *bus.Bus
}

// NewNotificationService ...
func NewNotificationService(bbus *bus.Bus) NotificationService {
	ns := &notificationService{
		bbus: bbus,
	}

	ns.RegisterSubscriber()

	return ns
}

func (n *notificationService) RegisterSubscriber() {
	n.bbus.RegisterHandler("notification-service-order.*", &bus.Handler{
		Matcher: eventbus.OrderAll,
		Handle:  n.Subscriber,
	})

	n.bbus.RegisterHandler("notification-service-payment.*", &bus.Handler{
		Matcher: eventbus.PaymentAll,
		Handle:  n.Subscriber,
	})

}

func (n *notificationService) CreateNotification(msg string) {
	if msg == "" {
		return
	}

	mu := sync.Mutex{}
	mu.Lock()
	n.notifs = append(n.notifs, msg)
	mu.Unlock()
	return
}

func (n *notificationService) List() []string {
	return n.notifs
}

func (n *notificationService) Subscriber(event *bus.Event) {
	go func() {
		switch event.Topic {
		case eventbus.OrderCreated:
			order, ok := event.Data.(model.Order)
			if !ok {
				return
			}

			msg := fmt.Sprintf("order %d created", order.ID)
			logrus.Info(msg)
			n.CreateNotification(msg)

		case eventbus.PaymentPaid:
			payment, ok := event.Data.(model.Payment)
			if !ok {
				return
			}

			msg := fmt.Sprintf("payment %d paid", payment.ID)
			logrus.Info(msg)
			n.CreateNotification(msg)
		}
	}()
}
