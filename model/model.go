package model

import "fmt"

// Product :nodoc:
type Product struct {
	ID    int64
	Price float64
}

func (p Product) String() string {
	return fmt.Sprintf(`{"id": %d, "price":	%f}`, p.ID, p.Price)
}

// Order :nodoc:
type Order struct {
	ID         int64
	ProductIDs []int64
}

func (o Order) String() string {
	return fmt.Sprintf(`{"id": %d, "productIDs": %v}`, o.ID, o.ProductIDs)
}

// PaymentStatus :nodoc:
type PaymentStatus int

func (p PaymentStatus) String() string {
	switch p {
	case PaymentStatusPending:
		return "pending"
	case PaymentStatusPaid:
		return "paid"
	case PaymentStatusCanceled:
		return "canceled"
	default:
		return ""
	}
}

// PaymentStatus enum
const (
	PaymentStatusPending  = PaymentStatus(1)
	PaymentStatusPaid     = PaymentStatus(2)
	PaymentStatusCanceled = PaymentStatus(3)
)

// Payment :nodoc:
type Payment struct {
	ID      int64
	OrderID int64
	Status  PaymentStatus
}

func (p Payment) String() string {
	return fmt.Sprintf("{id: %d, order_id: %d, status: %v}", p.ID, p.OrderID, p.Status)
}
