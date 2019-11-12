package goney

import (
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want Currency
	}{
		{"returns the currency with provided code, if found",
			args{"JPY"},
			JPY,
		},
		{"returns XXX (no currency), if not found",
			args{"OMG"},
			XXX,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
