package goney

import (
	"testing"

	"github.com/shopspring/decimal"
)

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
			if !got.Equals(tt.want) {
				t.Errorf("Money.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Sub(t *testing.T) {
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
		{"returns a new instance whose amount is the difference between the two amounts, when currencies match",
			Money{
				decimal.NewFromFloat(5.67123),
				EUR,
			},
			args{Money{decimal.NewFromFloat(1.2398), EUR}},
			Money{decimal.NewFromFloat(4.43143), EUR},
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
			got, err := tt.input.Sub(tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Money.Sub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equals(tt.want) {
				t.Errorf("Money.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Div(t *testing.T) {
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
			got, err := tt.input.Div(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Money.Div() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equals(tt.want) {
				t.Errorf("Money.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Mul(t *testing.T) {
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
			got := tt.input.Mul(tt.args)
			if !got.Equals(tt.want) {
				t.Errorf("Money.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Equals(t *testing.T) {
	type args struct {
		x Money
	}
	tests := []struct {
		name  string
		input Money
		args  args
		want  bool
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
		},
		{"returns false when instances have the same amount but different currencies",
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.Equals(tt.args.x); got != tt.want {
				t.Errorf("Money.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_RoundedString(t *testing.T) {
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
			if got := tt.input.RoundedString(); got != tt.want {
				t.Errorf("Money.FormattedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
