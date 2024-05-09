package anime

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/common"
	"github.com/dmji/go-myanimelist/mal/containers"
	"github.com/dmji/go-myanimelist/mal/prm"
)

func optionFromAnimeRanking(r prm.AnimeRanking) common.OptionFunc {
	return common.OptionFunc(func(v *url.Values) {
		v.Set("ranking_type", string(r))
	})
}

// Ranking allows an authenticated user to receive the top anime based on a
// certain ranking.
func (s *Service) Ranking(ctx context.Context, ranking prm.AnimeRanking, options ...prm.OptionalParam) ([]containers.Anime, *common.Response, error) {
	options = append(options, optionFromAnimeRanking(ranking))
	return s.list(ctx, "anime/ranking", options...)
}
