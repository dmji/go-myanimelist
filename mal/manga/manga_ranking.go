package manga

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/common"
	"github.com/dmji/go-myanimelist/mal/containers"
	"github.com/dmji/go-myanimelist/mal/prm"
)

func optionFromMangaRanking(r prm.MangaRanking) common.OptionFunc {
	return common.OptionFunc(func(v *url.Values) {
		v.Set("ranking_type", string(r))
	})
}

// Ranking allows an authenticated user to receive the top manga based on a
// certain ranking.
func (s *Service) Ranking(ctx context.Context, ranking prm.MangaRanking, options ...prm.OptionalParam) ([]containers.Manga, *common.Response, error) {
	options = append(options, optionFromMangaRanking(ranking))
	return s.list(ctx, "manga/ranking", options...)
}
