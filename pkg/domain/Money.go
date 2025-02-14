package domain

import "math"

type Quantity int64 // @name Quantity

type Currency string // @name Currency

const (
	CurrencyBRL Currency = "BRL"
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
)

func (a Quantity) Int64() int64 {
	return int64(a)
}

type Money struct {
	Quantity Quantity `json:"quantity,omitempty" bson:"quantity,omitempty" example:"1000"`
	Currency Currency `json:"currency,omitempty" bson:"currency,omitempty" example:"BRL"` // default code BRL
} // @name Money

func (m Money) Float() float64 {
	if m.IsZero() {
		return 0
	}
	d := m.DigitsAsCents()
	return float64(m.Quantity) / float64(d)
}

func (m Money) IsZero() bool {
	return m.Quantity == 0
}

func (m Money) DigitsAsCents() int {
	// minor unit = 2 BRL
	ce := math.Pow(10, float64(2))
	return int(ce)
}
