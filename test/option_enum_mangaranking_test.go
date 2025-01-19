package mal_test

import (
	"testing"

	"github.com/dmji/go-myanimelist/mal_opt"
)

func TestOptionMangaRanking(t *testing.T) {
	tests := []struct {
		name string
		in   func(mal_opt.MangaRanking) mal_opt.MangaRanking
		out  mal_opt.MangaRanking
	}{
		{
			name: "All",
			in:   mal_opt.MangaRanking.All,
			out:  mal_opt.MangaRankingAll,
		},
		{
			name: "Manga",
			in:   mal_opt.MangaRanking.Manga,
			out:  mal_opt.MangaRankingManga,
		},
		{
			name: "Oneshots",
			in:   mal_opt.MangaRanking.Oneshots,
			out:  mal_opt.MangaRankingOneshots,
		},
		{
			name: "Doujinshi",
			in:   mal_opt.MangaRanking.Doujinshi,
			out:  mal_opt.MangaRankingDoujinshi,
		},
		{
			name: "LightNovels",
			in:   mal_opt.MangaRanking.LightNovels,
			out:  mal_opt.MangaRankingLightNovels,
		},
		{
			name: "Novels",
			in:   mal_opt.MangaRanking.Novels,
			out:  mal_opt.MangaRankingNovels,
		},
		{
			name: "Manhwa",
			in:   mal_opt.MangaRanking.Manhwa,
			out:  mal_opt.MangaRankingManhwa,
		},
		{
			name: "Manhua",
			in:   mal_opt.MangaRanking.Manhua,
			out:  mal_opt.MangaRankingManhua,
		},
		{
			name: "ByPopularity",
			in:   mal_opt.MangaRanking.ByPopularity,
			out:  mal_opt.MangaRankingByPopularity,
		},
		{
			name: "Favorite",
			in:   mal_opt.MangaRanking.Favorite,
			out:  mal_opt.MangaRankingFavorite,
		},
	}

	p := new(mal_opt.MangaRanking)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in(*p)
			want := tt.out

			if got != want {
				t.Errorf("MangaRanking expected '%s', got '%s'", want, got)
			}
		})
	}
}
