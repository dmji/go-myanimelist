package mal_test

import (
	"testing"

	"github.com/dmji/go-myanimelist/mal_opt"
)

func TestOptionAnimeSeason(t *testing.T) {
	p := new(mal_opt.AnimeSeason)
	tests := []struct {
		name string
		in   func() mal_opt.AnimeSeason
		out  mal_opt.AnimeSeason
	}{
		{
			name: "Winter",
			in:   p.Winter,
			out:  mal_opt.AnimeSeasonWinter,
		},
		{
			name: "Spring",
			in:   p.Spring,
			out:  mal_opt.AnimeSeasonSpring,
		},
		{
			name: "Summer",
			in:   p.Summer,
			out:  mal_opt.AnimeSeasonSummer,
		},
		{
			name: "Fall",
			in:   p.Fall,
			out:  mal_opt.AnimeSeasonFall,
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
