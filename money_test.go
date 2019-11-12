package goney

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func TestFromFloat(t *testing.T) {
	type args struct {
		amount   float64
		currency Currency
	}
	tests := []struct {
		name string
		args args
		want Money
	}{
		{"returns a Money instance with provided amount and currency",
			args{1.23, EUR},
			Money{decimal.NewFromFloat(1.23), EUR},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromFloat(tt.args.amount, tt.args.currency); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromInt(t *testing.T) {
	type args struct {
		amount   int64
		currency Currency
	}
	tests := []struct {
		name string
		args args
		want Money
	}{
		{"returns a Money instance with provided amount and currency",
			args{100, EUR},
			Money{decimal.New(100, 10), EUR},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt(tt.args.amount, tt.args.currency); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    Money
		wantErr bool
	}{
		{"returns a Money instance with provided amount and currency",
			args{"EUR 123.34"},
			Money{decimal.NewFromFloat(123.34), EUR},
			false,
		},
		{"returns an error, if currency is unknown",
			args{"OMG 123.34"},
			Money{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromString(tt.args.value)

			if (err != nil) != tt.wantErr {
				t.Errorf("FromString() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Add(t *testing.T) {
	type args struct {
		x Money
	}
	tests := []struct {
		name    string
		input   Money
		args    args
		want    Money
		wantErr bool
	}{
		{"returns a new instance whose amount is the sum of the two amounts, when currencies match",
			Money{
				decimal.NewFromFloat(5.67123),
				EUR,
			},
			args{Money{decimal.NewFromFloat(1.2398), EUR}},
			Money{decimal.NewFromFloat(6.91103), EUR},
			false,
		},
		{"returns an error, if currencies do not match",
			Money{
				decimal.NewFromFloat(1),
				EUR,
			},
			args{Money{decimal.NewFromFloat(2), GBP}},
			Money{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.Add(tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Money.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			eq, err := got.Equal(tt.want)
			if err != nil || !eq {
				t.Errorf("Money.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Subtract(t *testing.T) {
	type args struct {
		x Money
	}
	tests := []struct {
		name    string
		input   Money
		args    args
		want    Money
		wantErr bool
	}{
		{"returns an error, if currencies do not match",
			Money{
				decimal.NewFromFloat(1),
				EUR,
			},
			args{Money{decimal.NewFromFloat(2), GBP}},
			Money{},
			true,
		},
		{"returns a new instance whose amount is the difference between the two amounts, otherwise",
			Money{
				decimal.NewFromFloat(5.67123),
				EUR,
			},
			args{Money{decimal.NewFromFloat(1.2398), EUR}},
			Money{decimal.NewFromFloat(4.43143), EUR},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.Subtract(tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Money.Subtract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			eq, err := got.Equal(tt.want)
			if err != nil || !eq {
				t.Errorf("Money.Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Divide(t *testing.T) {
	tests := []struct {
		name    string
		input   Money
		args    decimal.Decimal
		want    Money
		wantErr bool
	}{
		{"returns a new instance whose amount is the division of the two amounts",
			Money{
				decimal.NewFromFloat(3),
				EUR,
			},
			decimal.NewFromFloat(2),
			Money{decimal.NewFromFloat(1.5), EUR},
			false,
		},
		{"returns an error in case of division by zero",
			Money{
				decimal.NewFromFloat(3),
				EUR,
			},
			decimal.NewFromFloat(0),
			Money{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.Divide(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Money.Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			eq, err := got.Equal(tt.want)
			if err != nil || !eq {
				t.Errorf("Money.Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Multiply(t *testing.T) {
	tests := []struct {
		name  string
		input Money
		args  decimal.Decimal
		want  Money
	}{
		{"returns a new instance whose amount is the multiplication of the two amounts",
			Money{
				decimal.NewFromFloat(5.67123),
				EUR,
			},
			decimal.NewFromFloat(1.2398),
			Money{decimal.NewFromFloat(7.031190954), EUR},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Multiply(tt.args)
			eq, err := got.Equal(tt.want)
			if err != nil || !eq {
				t.Errorf("Money.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Compare(t *testing.T) {
	type args struct {
		x Money
	}
	tests := []struct {
		name    string
		input   Money
		args    args
		want    int
		wantErr bool
	}{
		{"returns an error if currencies don't match",
			Money{
				decimal.NewFromFloat(1.23),
				EUR,
			},
			args{
				Money{
					decimal.NewFromFloat(1.23),
					JPY,
				},
			},
			0,
			true,
		},
		{"returns -1 when currencies match and first amount is lower than the second one",
			Money{
				decimal.NewFromFloat(1.23),
				EUR,
			},
			args{
				Money{
					decimal.NewFromFloat(3.4),
					EUR,
				},
			},
			-1,
			false,
		},
		{"returns 0 when currencies match and amounts are the same",
			Money{
				decimal.NewFromFloat(1.23),
				EUR,
			},
			args{
				Money{
					decimal.NewFromFloat(1.23),
					EUR,
				},
			},
			0,
			false,
		},
		{"returns 1 when currencies match and first amount is greater than the second one",
			Money{
				decimal.NewFromFloat(3.4),
				EUR,
			},
			args{
				Money{
					decimal.NewFromFloat(1.23),
					EUR,
				},
			},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.Compare(tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Money.Compare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Money.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Equal(t *testing.T) {
	type args struct {
		x Money
	}
	tests := []struct {
		name    string
		input   Money
		args    args
		want    bool
		wantErr bool
	}{
		{"returns true when instances have the same currency and amount",
			Money{
				decimal.NewFromFloat(1.23),
				EUR,
			},
			args{
				Money{
					decimal.NewFromFloat(1.23),
					EUR,
				},
			},
			true,
			false,
		},
		{"returns false when instances have the same currency but different amounts",
			Money{
				decimal.NewFromFloat(1.23),
				EUR,
			},
			args{
				Money{
					decimal.NewFromFloat(3.4),
					EUR,
				},
			},
			false,
			false,
		},
		{"returns an error when instances have different currencies",
			Money{
				decimal.NewFromFloat(1.23),
				EUR,
			},
			args{
				Money{
					decimal.NewFromFloat(1.23),
					GBP,
				},
			},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.Equal(tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Money.Equal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Money.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_FormattedString(t *testing.T) {
	tests := []struct {
		name  string
		input Money
		want  string
	}{
		{"formats the amount according to provided precision, rounding it if needed",
			Money{
				decimal.NewFromFloat(1.23456789),
				EUR,
			},
			"EUR 1.23",
		},
		{"formats the amount according to provided precision, rounding it if needed",
			Money{
				decimal.NewFromFloat(1.23456789),
				JPY,
			},
			"JPY 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.FormattedString(); got != tt.want {
				t.Errorf("Money.FormattedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
