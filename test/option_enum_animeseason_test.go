package mal_test

import (
	"testing"

	"github.com/dmji/go-myanimelist/mal/prm"
)

func TestOptionAnimeSeason(t *testing.T) {
	p := new(prm.AnimeSeason)
	tests := []struct {
		name string
		in   func() prm.AnimeSeason
		out  prm.AnimeSeason
	}{
		{
			name: "Winter",
			in:   p.Winter,
			out:  prm.AnimeSeasonWinter,
		},
		{
			name: "Spring",
			in:   p.Spring,
			out:  prm.AnimeSeasonSpring,
		},
		{
			name: "Summer",
			in:   p.Summer,
			out:  prm.AnimeSeasonSummer,
		},
		{
			name: "Fall",
			in:   p.Fall,
			out:  prm.AnimeSeasonFall,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in()
			want := tt.out

			if got != want {
				t.Errorf("AnimeSeason expected '%s', got '%s'", want, got)
			}
		})
	}
}
