package mal_test

import (
	"testing"

	"github.com/dmji/go-myanimelist/mal_opt"
)

func TestOptionRereadValue(t *testing.T) {
	tests := []struct {
		name string
		in   func(mal_opt.RereadValue) mal_opt.RereadValue
		out  mal_opt.RereadValue
	}{
		{
			name: "NoValue",
			in:   mal_opt.RereadValue.NoValue,
			out:  mal_opt.RereadNoValue,
		},
		{
			name: "VeryLow",
			in:   mal_opt.RereadValue.VeryLow,
			out:  mal_opt.RereadVeryLow,
		},
		{
			name: "Low",
			in:   mal_opt.RereadValue.Low,
			out:  mal_opt.RereadLow,
		},
		{
			name: "Medium",
			in:   mal_opt.RereadValue.Medium,
			out:  mal_opt.RereadMedium,
		},
		{
			name: "High",
			in:   mal_opt.RereadValue.High,
			out:  mal_opt.RereadHigh,
		},
		{
			name: "VeryHigh",
			in:   mal_opt.RereadValue.VeryHigh,
			out:  mal_opt.RereadVeryHigh,
		},
	}

	p := new(mal_opt.RereadValue)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in(*p)
			want := tt.out

			if got != want {
				t.Errorf("RereadValue expected '%d', got '%d'", want, got)
			}
		})
	}
}
