package dto

import (
	"time"
	"github.com/mysterium/node/money"
)

const PAYMENT_METHOD_PER_TIME = "PER_TIME"

type PaymentMethodPerTime struct {
	Price money.Money `json:"price"`

	// Service duration provided for paid price
	Duration time.Duration `json:"duration"`
}

func (method PaymentMethodPerTime) GetPrice() money.Money {
	return method.Price
}
