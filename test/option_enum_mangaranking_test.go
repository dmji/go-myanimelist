package mal_test

import (
	"testing"

	"github.com/dmji/go-myanimelist/mal/prm"
)

func TestOptionMangaRanking(t *testing.T) {
	tests := []struct {
		name string
		in   func(prm.MangaRanking) prm.MangaRanking
		out  prm.MangaRanking
	}{
		{
			name: "All",
			in:   prm.MangaRanking.All,
			out:  prm.MangaRankingAll,
		},
		{
			name: "Manga",
			in:   prm.MangaRanking.Manga,
			out:  prm.MangaRankingManga,
		},
		{
			name: "Oneshots",
			in:   prm.MangaRanking.Oneshots,
			out:  prm.MangaRankingOneshots,
		},
		{
			name: "Doujinshi",
			in:   prm.MangaRanking.Doujinshi,
			out:  prm.MangaRankingDoujinshi,
		},
		{
			name: "LightNovels",
			in:   prm.MangaRanking.LightNovels,
			out:  prm.MangaRankingLightNovels,
		},
		{
			name: "Novels",
			in:   prm.MangaRanking.Novels,
			out:  prm.MangaRankingNovels,
		},
		{
			name: "Manhwa",
			in:   prm.MangaRanking.Manhwa,
			out:  prm.MangaRankingManhwa,
		},
		{
			name: "Manhua",
			in:   prm.MangaRanking.Manhua,
			out:  prm.MangaRankingManhua,
		},
		{
			name: "ByPopularity",
			in:   prm.MangaRanking.ByPopularity,
			out:  prm.MangaRankingByPopularity,
		},
		{
			name: "Favorite",
			in:   prm.MangaRanking.Favorite,
			out:  prm.MangaRankingFavorite,
		},
	}

	p := new(prm.MangaRanking)
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
