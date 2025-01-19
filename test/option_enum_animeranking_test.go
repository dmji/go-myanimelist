package mal_test

import (
	"testing"

	"github.com/dmji/go-myanimelist/mal_opt"
)

func TestOptionAnimeRanking(t *testing.T) {
	tests := []struct {
		name string
		in   func(mal_opt.AnimeRanking) mal_opt.AnimeRanking
		out  mal_opt.AnimeRanking
	}{
		{
			name: "All",
			in:   mal_opt.AnimeRanking.All,
			out:  mal_opt.AnimeRankingAll,
		},
		{
			name: "Airing",
			in:   mal_opt.AnimeRanking.Airing,
			out:  mal_opt.AnimeRankingAiring,
		},
		{
			name: "Upcoming",
			in:   mal_opt.AnimeRanking.Upcoming,
			out:  mal_opt.AnimeRankingUpcoming,
		},
		{
			name: "TV",
			in:   mal_opt.AnimeRanking.TV,
			out:  mal_opt.AnimeRankingTV,
		},
		{
			name: "OVA",
			in:   mal_opt.AnimeRanking.OVA,
			out:  mal_opt.AnimeRankingOVA,
		},
		{
			name: "Movie",
			in:   mal_opt.AnimeRanking.Movie,
			out:  mal_opt.AnimeRankingMovie,
		},
		{
			name: "Special",
			in:   mal_opt.AnimeRanking.Special,
			out:  mal_opt.AnimeRankingSpecial,
		},
		{
			name: "ByPopularity",
			in:   mal_opt.AnimeRanking.ByPopularity,
			out:  mal_opt.AnimeRankingByPopularity,
		},
		{
			name: "Favorite",
			in:   mal_opt.AnimeRanking.Favorite,
			out:  mal_opt.AnimeRankingFavorite,
		},
	}

	p := new(mal_opt.AnimeRanking)
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
