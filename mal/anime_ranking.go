package mal

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/common"
)

// AnimeRanking allows to choose how the anime will be ranked.
type AnimeRanking string

const (
	// AnimeRankingAll returns the top anime series.
	AnimeRankingAll AnimeRanking = "all"
	// AnimeRankingAiring returns the top airing anime.
	AnimeRankingAiring AnimeRanking = "airing"
	// AnimeRankingUpcoming returns the top upcoming anime.
	AnimeRankingUpcoming AnimeRanking = "upcoming"
	// AnimeRankingTV returns the top Anime TV series.
	AnimeRankingTV AnimeRanking = "tv"
	// AnimeRankingOVA returns the top anime OVA series.
	AnimeRankingOVA AnimeRanking = "ova"
	// AnimeRankingMovie returns the top anime movies.
	AnimeRankingMovie AnimeRanking = "movie"
	// AnimeRankingSpecial returns the top anime specials.
	AnimeRankingSpecial AnimeRanking = "special"
	// AnimeRankingByPopularity returns the top anime by popularity.
	AnimeRankingByPopularity AnimeRanking = "bypopularity"
	// AnimeRankingFavorite returns the top favorite Anime.
	AnimeRankingFavorite AnimeRanking = "favorite"
)

func optionFromAnimeRanking(r AnimeRanking) common.OptionFunc {
	return common.OptionFunc(func(v *url.Values) {
		v.Set("ranking_type", string(r))
	})
}

// Ranking allows an authenticated user to receive the top anime based on a
// certain ranking.
func (s *AnimeService) Ranking(ctx context.Context, ranking AnimeRanking, options ...common.OptionalParam) ([]Anime, *Response, error) {
	options = append(options, optionFromAnimeRanking(ranking))
	return s.list(ctx, "anime/ranking", options...)
}
