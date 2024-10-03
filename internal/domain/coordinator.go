package domain

type Coordinator struct {
	OrderUID string
	Order    Order
	Delivery Delivery
	Payment  Payment
	Items    []Item
}
