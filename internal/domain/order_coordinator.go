package domain

type OrderCoordinator struct {
	OrderUID string
	Order    Order
	Delivery Delivery
	Payment  Payment
	Items    []Item
}
