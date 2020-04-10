package eventbus

// bus topics ...
const (
	OrderAll   = "order.*"
	PaymentAll = "payment.*"

	OrderCreated   = "order.created"
	OrderDeleted   = "order.deleted"
	PaymentCreated = "payment.created"
	PaymentPaid    = "payment.paid"
)

// BusTopics return all bus topics
func BusTopics() []string {
	return []string{
		OrderCreated,
		OrderDeleted,
		PaymentCreated,
		PaymentPaid,
	}
}
