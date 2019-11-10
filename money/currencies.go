package goney

// Currency represents an ISO4217 currency
type Currency struct {
	Code      string
	Precision int32
}

func (c Currency) String() string {
	return c.Code
}

// EUR Euro
var EUR = Currency{"EUR", 2}

// GBP British Pound
var GBP = Currency{"GBP", 2}

// JPY Japanese Yen
var JPY = Currency{"JPY", 0}
