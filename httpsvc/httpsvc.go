package httpsvc

import (
	"encoding/json"
	"net/http"
	"os"
	"os/signal"

	"shop/eventbus"
	"shop/model"
	"shop/service"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/kumparan/go-lib/utils"
	log "github.com/sirupsen/logrus"
)

// Run run http service
func Run() {
	bbus := eventbus.NewBus()
	bbus.RegisterTopics(eventbus.BusTopics()...)

	productService := service.NewProductService()
	orderService := service.NewOrderService(productService, bbus)
	paymentSerivce := service.NewPaymentService(orderService, bbus)
	notificationService := service.NewNotificationService(bbus)

	r := chi.NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/api/products", func(w http.ResponseWriter, r *http.Request) {
		products := productService.List()
		w.Write(dumpJSON(products))
	})

	r.Get("/api/notifications", func(w http.ResponseWriter, r *http.Request) {
		d := dumpJSON(notificationService.List())
		w.Write(d)
	})

	// order?productID=xxx-xxxx
	r.Post("/api/order", func(w http.ResponseWriter, r *http.Request) {
		productID := utils.String2Int64(r.URL.Query().Get("productID"))
		if productID <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			dumpJSON(map[string]string{
				"message": "invalid productID",
			})
			return
		}

		order := orderService.CreateOrder([]int64{productID})
		w.Write(dumpJSON(order))
	})

	r.Get("/api/payments", func(w http.ResponseWriter, r *http.Request) {
		d := dumpJSON(paymentSerivce.List())
		w.Write(d)
	})

	// ?paymentID=xxx-xxx
	r.Post("/api/payments/paid", func(w http.ResponseWriter, r *http.Request) {
		paymentID := utils.String2Int64(r.URL.Query().Get("paymentID"))
		if paymentID <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			dumpJSON(map[string]string{
				"message": "invalid paymentID",
			})
			return
		}

		pm := &model.Payment{
			ID: paymentID,
		}
		if ok := paymentSerivce.PayBills(pm); !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(dumpJSON(map[string]string{
				"message": "failed paid bills",
			}))
			return
		}

		w.Write(dumpJSON(pm))
	})

	port := ":3000"

	go func() {
		log.Info("service running at ", port)
		http.ListenAndServe(port, r)
	}()

	// listen interrupt
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

func dumpJSON(v interface{}) []byte {
	if v == nil {
		return nil
	}

	b, err := json.Marshal(v)
	if err != nil {
		log.Error(err)
	}

	return b
}
