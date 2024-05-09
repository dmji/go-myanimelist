package anime

import (
	"context"
	"fmt"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/common"
	"github.com/dmji/go-myanimelist/mal/containers"
	"github.com/dmji/go-myanimelist/mal/prm"
)

func optionFromSeasonalAnimeOption(o prm.SeasonalAnimeOption) common.OptionFunc {
	return common.OptionFunc(func(v *url.Values) {
		o.SeasonalAnimeApply(v)
	})
}

// Seasonal allows an authenticated user to receive the seasonal anime by
// providing the year and season. The results can be sorted using an option.
func (s *Service) Seasonal(ctx context.Context, year int, season prm.AnimeSeason, options ...prm.SeasonalAnimeOption) ([]containers.Anime, *common.Response, error) {
	oo := make([]prm.OptionalParam, len(options))
	for i := range options {
		oo[i] = optionFromSeasonalAnimeOption(options[i])
	}
	return s.list(ctx, fmt.Sprintf("anime/season/%d/%s", year, season), oo...)
}
