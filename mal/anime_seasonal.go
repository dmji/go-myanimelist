package mal

import (
	"context"
	"fmt"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/common"
)

// AnimeSeason is the airing season of the anime.
type AnimeSeason string

const (
	// AnimeSeasonWinter is the winter season of January, February and March.
	AnimeSeasonWinter AnimeSeason = "winter"
	// AnimeSeasonSpring is the spring season of April, May and June.
	AnimeSeasonSpring AnimeSeason = "spring"
	// AnimeSeasonSummer is the summer season of July, August and September.
	AnimeSeasonSummer AnimeSeason = "summer"
	// AnimeSeasonFall is the fall season of October, November and December.
	AnimeSeasonFall AnimeSeason = "fall"
)

// SortSeasonalAnime is an option that allows to sort the anime results.
type SortSeasonalAnime string

const (
	// SortSeasonalByAnimeScore sorts seasonal results by anime score in
	// descending order.
	SortSeasonalByAnimeScore SortSeasonalAnime = "anime_score"
	// SortSeasonalByAnimeNumListUsers sorts seasonal results by anime num list
	// users in descending order.
	SortSeasonalByAnimeNumListUsers SortSeasonalAnime = "anime_num_list_users"
)

func (s SortSeasonalAnime) SeasonalAnimeApply(v *url.Values) { v.Set("sort", string(s)) }

// SeasonalAnimeOption are options specific to the AnimeService.Seasonal method.
type SeasonalAnimeOption interface {
	SeasonalAnimeApply(v *url.Values)
}

func optionFromSeasonalAnimeOption(o SeasonalAnimeOption) common.OptionFunc {
	return common.OptionFunc(func(v *url.Values) {
		o.SeasonalAnimeApply(v)
	})
}

// Seasonal allows an authenticated user to receive the seasonal anime by
// providing the year and season. The results can be sorted using an option.
func (s *AnimeService) Seasonal(ctx context.Context, year int, season AnimeSeason, options ...SeasonalAnimeOption) ([]Anime, *Response, error) {
	oo := make([]common.OptionalParam, len(options))
	for i := range options {
		oo[i] = optionFromSeasonalAnimeOption(options[i])
	}
	return s.list(ctx, fmt.Sprintf("anime/season/%d/%s", year, season), oo...)
}
