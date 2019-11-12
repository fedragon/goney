package goney

import (
	"errors"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

// Money represents a monetary amount
type Money struct {
	Amount   decimal.Decimal
	Currency Currency
}

// FromFloat returns a Money instance representing the provided
// amount and currency
func FromFloat(amount float64, currency Currency) Money {
	return Money{decimal.NewFromFloat(amount), currency}
}

// FromInt returns a Money instance representing the provided
// amount and currency
func FromInt(amount int64, currency Currency) Money {
	return Money{decimal.New(amount, 10), currency}
}

// FromString returns a Money instance representing the provided
// amount and currency. It assumes that the input string will
// have the form: <currency code> <amount>
// e.g. EUR 1.234
func FromString(value string) (Money, error) {
	if parts := strings.Split(value, " "); len(parts) == 2 {
		amount, err := decimal.NewFromString(parts[1])

		if err != nil {
			return Money{}, fmt.Errorf("invalid amount %v", amount)
		}
		currency := Find(parts[0])

		if currency == XXX {
			return Money{}, fmt.Errorf("unknown currency %v", parts[0])
		}

		return Money{amount, currency}, nil
	}

	return Money{}, errors.New("unable to parse string into money. expected format: currency-code amount")
}

func (m Money) String() string {
	return fmt.Sprintf("%v %v", m.Currency, m.Amount)
}

// FormattedString returns a string representation of this
// instance, rounding its amount according to the currency's own precision
func (m Money) FormattedString() string {
	return fmt.Sprintf("%v %v", m.Currency, m.Amount.RoundBank(m.Currency.Precision))
}

// Equal compares two Money instances for deep equality,
// returning true if their amounts and currencies are themselves equal
func (m Money) Equal(x Money) (bool, error) {
	cmp, err := m.Compare(x)

	if err != nil {
		return false, err
	}

	return cmp == 0, err
}

func (m Money) sameCurrency(x Money) error {
	if m.Currency != x.Currency {
		return errors.New("currencies do not match")
	}

	return nil
}

// Compare compares two Money instances and returns -1, 0, 1
// if the first amount is respectively lower than, equal to,
// or greater than the second one.
// It returns an error if their currencies don't match.
func (m Money) Compare(x Money) (int, error) {
	if err := m.sameCurrency(x); err != nil {
		return 0, err
	}

	return m.Amount.Cmp(x.Amount), nil
}

// Add returns a Money whose amount is the sum of the amounts in this
// and the other Money instance, if their currencies match.
// It returns an error if their currencies do not match
func (m Money) Add(x Money) (Money, error) {
	if err := m.sameCurrency(x); err != nil {
		return Money{}, err
	}

	return Money{m.Amount.Add(x.Amount), m.Currency}, nil
}

// Subtract returns a Money whose amount is the difference of the amounts in
// this and the other Money instance, if their currencies match.
// It returns an error if their currencies do not match
func (m Money) Subtract(x Money) (Money, error) {
	if err := m.sameCurrency(x); err != nil {
		return Money{}, err
	}

	return Money{m.Amount.Sub(x.Amount), m.Currency}, nil
}

// Multiply returns a Money whose amount is the multiplication of
// the amount in this instance by the provided multiplier
func (m Money) Multiply(x decimal.Decimal) Money {
	return Money{m.Amount.Mul(x), m.Currency}
}

// Divide returns a Money whose amount is the division of the amount
// in this instance by the provided dividend.
// It returns an error if the dividend is zero.
func (m Money) Divide(x decimal.Decimal) (Money, error) {
	if x.IsZero() {
		return Money{}, fmt.Errorf("cannot divide by zero")
	}

	return Money{m.Amount.Div(x), m.Currency}, nil
}
