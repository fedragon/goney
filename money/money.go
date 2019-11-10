package goney

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// Money represents a monetary amount
type Money struct {
	Amount   decimal.Decimal
	Currency Currency
}

func (m Money) String() string {
	return fmt.Sprintf("%v %v", m.Currency, m.Amount)
}

// RoundedString returns a string representation of this
// instance, rounding its amount according to the currency's own precision
func (m Money) RoundedString() string {
	return fmt.Sprintf("%v %v", m.Currency, m.Amount.RoundBank(m.Currency.Precision))
}

// Equals compares two Money instances for deep equality,
// returning true if their amounts and currencies are themselves equal
func (m Money) Equals(x Money) bool {
	return m.Amount.Equal(x.Amount) && m.Currency == x.Currency
}

// Add returns a Money whose amount is the sum of the amounts in this
// and the other Money instance, if their currencies match.
// It returns an error if their currencies do not match
func (m Money) Add(x Money) (Money, error) {
	if m.Currency != x.Currency {
		return Money{}, fmt.Errorf("different currencies: %v, %v", m.Currency, x.Currency)
	}

	return Money{m.Amount.Add(x.Amount), m.Currency}, nil
}

// Sub returns a Money whose amount is the difference of the amounts in
// this and the other Money instance, if their currencies match.
// It returns an error if their currencies do not match
func (m Money) Sub(x Money) (Money, error) {
	if m.Currency != x.Currency {
		return Money{}, fmt.Errorf("different currencies: %v, %v", m.Currency, x.Currency)
	}

	return Money{m.Amount.Sub(x.Amount), m.Currency}, nil
}

// Mul returns a Money whose amount is the multiplication of the amount
// in this instance by the provided multiplier
func (m Money) Mul(x decimal.Decimal) Money {
	return Money{m.Amount.Mul(x), m.Currency}
}

// Div returns a Money whose amount is the division of the amount
// in this instance by the provided dividend.
// It returns an error if the dividend is zero
func (m Money) Div(x decimal.Decimal) (Money, error) {
	if x.IsZero() {
		return Money{}, fmt.Errorf("cannot divide by zero")
	}

	return Money{m.Amount.Div(x), m.Currency}, nil
}
