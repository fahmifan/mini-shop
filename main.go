package main

import (
	"os"
	"os/signal"
	"syscall"

	"shop/eventbus"
	"shop/eventhandler"
	"shop/service"

	"github.com/mustafaturan/bus"
	log "github.com/sirupsen/logrus"
)

func main() {
	handler := &eventhandler.EventHandler{}

	bbus := eventbus.NewBus()
	bbus.RegisterTopics(eventbus.BusTopics()...)
	bbus.RegisterHandler("order-channel", &bus.Handler{
		Matcher: eventbus.OrderAll,
		Handle:  handler.HandleOrder,
	})

	productService := service.NewProductService()
	orderService := service.NewOrderService(productService, bbus)
	paymentSerivce := service.NewPaymentService(orderService)

	handler.PaymentService = paymentSerivce

	products := productService.List()
	orderService.CreateOrder([]int64{products[0].ID})

	sigCh := make(chan os.Signal)
	done := make(chan bool)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Info("exiting...")
		done <- true
	}()
	<-done
	close(done)
	close(sigCh)
}
