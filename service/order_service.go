package service

import (
	"context"
	"time"

	"shop/eventbus"
	"shop/model"

	"github.com/mustafaturan/bus"
	log "github.com/sirupsen/logrus"
)

type (
	orderService struct {
		bus            *bus.Bus
		productService ProductService

		orders []model.Order
	}
)

// NewOrderService :nodoc:
func NewOrderService(ps ProductService, bus *bus.Bus) OrderService {
	return &orderService{
		productService: ps,
		bus:            bus,
	}
}

// CreateOrder :nodoc:
func (o *orderService) CreateOrder(productIDs []int64) *model.Order {
	for _, id := range productIDs {
		if product := o.productService.FindProductByID(id); product == nil {
			return nil
		}
	}

	order := &model.Order{
		ID:         time.Now().UnixNano(),
		ProductIDs: productIDs,
	}

	log.Info("create order, productIDs: ", productIDs)
	go func() {
		_, err := o.bus.Emit(context.Background(), eventbus.OrderCreated, *order)
		if err != nil {
			log.Error(err)
			return
		}

	}()

	return order
}

func (o *orderService) FindOrderByID(id int64) *model.Order {
	for _, order := range o.orders {
		if order.ID == id {
			return &order
		}
	}

	return nil
}

func (o *orderService) add(order *model.Order) {
	if order == nil {
		return
	}

	o.orders = append(o.orders, *order)
}
