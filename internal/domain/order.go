package domain

import (
	"time"
)

type Order struct {
	OrderUID          string
	TrackNumber       string
	Entry             string
	Locale            string
	InternalSignature string
	CustomerID        string
	DeliveryService   string
	ShardKey          string
	SmID              string
	CreatedAt         time.Time
	OofShard          string
}
