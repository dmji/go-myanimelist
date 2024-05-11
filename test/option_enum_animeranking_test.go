package mal_test

import (
	"testing"

	"github.com/dmji/go-myanimelist/mal/prm"
)

func TestOptionAnimeRanking(t *testing.T) {
	tests := []struct {
		name string
		in   func(prm.AnimeRanking) prm.AnimeRanking
		out  prm.AnimeRanking
	}{
		{
			name: "All",
			in:   prm.AnimeRanking.All,
			out:  prm.AnimeRankingAll,
		},
		{
			name: "Airing",
			in:   prm.AnimeRanking.Airing,
			out:  prm.AnimeRankingAiring,
		},
		{
			name: "Upcoming",
			in:   prm.AnimeRanking.Upcoming,
			out:  prm.AnimeRankingUpcoming,
		},
		{
			name: "TV",
			in:   prm.AnimeRanking.TV,
			out:  prm.AnimeRankingTV,
		},
		{
			name: "OVA",
			in:   prm.AnimeRanking.OVA,
			out:  prm.AnimeRankingOVA,
		},
		{
			name: "Movie",
			in:   prm.AnimeRanking.Movie,
			out:  prm.AnimeRankingMovie,
		},
		{
			name: "Special",
			in:   prm.AnimeRanking.Special,
			out:  prm.AnimeRankingSpecial,
		},
		{
			name: "ByPopularity",
			in:   prm.AnimeRanking.ByPopularity,
			out:  prm.AnimeRankingByPopularity,
		},
		{
			name: "Favorite",
			in:   prm.AnimeRanking.Favorite,
			out:  prm.AnimeRankingFavorite,
		},
	}

	p := new(prm.AnimeRanking)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in(*p)
			want := tt.out

			if got != want {
				t.Errorf("AnimeRanking expected '%s', got '%s'", want, got)
			}
		})
	}
}
