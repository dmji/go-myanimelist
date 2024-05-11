package mal_test

import (
	"testing"

	"github.com/dmji/go-myanimelist/mal/prm"
)

func TestOptionRereadValue(t *testing.T) {
	tests := []struct {
		name string
		in   func(prm.RereadValue) prm.RereadValue
		out  prm.RereadValue
	}{
		{
			name: "NoValue",
			in:   prm.RereadValue.NoValue,
			out:  prm.RereadNoValue,
		},
		{
			name: "VeryLow",
			in:   prm.RereadValue.VeryLow,
			out:  prm.RereadVeryLow,
		},
		{
			name: "Low",
			in:   prm.RereadValue.Low,
			out:  prm.RereadLow,
		},
		{
			name: "Medium",
			in:   prm.RereadValue.Medium,
			out:  prm.RereadMedium,
		},
		{
			name: "High",
			in:   prm.RereadValue.High,
			out:  prm.RereadHigh,
		},
		{
			name: "VeryHigh",
			in:   prm.RereadValue.VeryHigh,
			out:  prm.RereadVeryHigh,
		},
	}

	p := new(prm.RereadValue)
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
