package domain

import "time"
type Payment struct {
	Transaction   string
	Currency      string
	Provider      string
	Amount        int64
	PaymentDt     time.Time
	Bank          string
	DeliveryCost  int64
	GoodsTotal    int64
	CustomFee     int64
}