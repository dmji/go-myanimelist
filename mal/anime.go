package mal

import (
	"context"
	"fmt"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/api_driver"
	"github.com/dmji/go-myanimelist/mal/containers"
	"github.com/dmji/go-myanimelist/mal/prm"
)

// AnimeService handles communication with the anime related methods of the
// MyAnimeList API:
//
// https://myanimelist.net/apiconfig/references/api/v2#tag/anime
// https://myanimelist.net/apiconfig/references/api/v2#tag/user-animelist
type AnimeService struct {
	client *api_driver.Client

	DetailsOptions            prm.DetailsOptionProvider
	ListOptions               prm.OptionalParamProvider
	SuggestedOptions          prm.OptionalParamProvider
	RankingOptions            prm.OptionalParamProvider
	UpdateMyListStatusOptions prm.UpdateMyAnimeListStatusOptionProvider
	SeasonalOptions           prm.SeasonalAnimeOptionProvider
}

func NewAnimeService(client *api_driver.Client) *AnimeService {
	return &AnimeService{
		client: client,
	}
}

// MARK: List
// List allows an authenticated user to search and list anime data. You may get
// user specific data by using the optional field.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/anime_get
func (s *AnimeService) List(ctx context.Context, search string, options ...prm.OptionalParam) ([]containers.Anime, *api_driver.Response, error) {
	options = append(options, optionFromQuery(search))
	rawOptions := optionsToFuncs(options, func(t prm.OptionalParam) func(*url.Values) { return t.Apply })
	return s.list(ctx, "anime", rawOptions...)
}

// MARK: Details
// Details returns details about an anime. By default, few anime fields are
// populated. Use the Fields option to specify which fields should be included.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/anime_anime_id_get
func (s *AnimeService) Details(ctx context.Context, animeID int, options ...prm.DetailsOption) (*containers.Anime, *api_driver.Response, error) {
	a := new(containers.Anime)
	rawOptions := detailsOptionsToFuncs(options)
	resp, err := s.client.RequestGet(ctx, fmt.Sprintf("anime/%d", animeID), a, rawOptions...)
	if err != nil {
		return nil, resp, err
	}
	return a, resp, nil
}

// MARK: Ranking
// Ranking allows an authenticated user to receive the top anime based on a
// certain ranking.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/anime_ranking_get
func (s *AnimeService) Ranking(ctx context.Context, ranking prm.AnimeRanking, options ...prm.OptionalParam) ([]containers.Anime, *api_driver.Response, error) {
	options = append(
		options,
		prm.OptionFunc(func(v *url.Values) {
			v.Set("ranking_type", string(ranking))
		}),
	)
	rawOptions := optionsToFuncs(options, func(t prm.OptionalParam) func(*url.Values) { return t.Apply })
	return s.list(ctx, "anime/ranking", rawOptions...)
}

// MARK: Seasonal
// Seasonal allows an authenticated user to receive the seasonal anime by
// providing the year and season. The results can be sorted using an option.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/anime_season_year_season_get
func (s *AnimeService) Seasonal(ctx context.Context, year int, season prm.AnimeSeason, options ...prm.SeasonalAnimeOption) ([]containers.Anime, *api_driver.Response, error) {
	rawOptions := optionsToFuncs(options, func(t prm.SeasonalAnimeOption) func(*url.Values) { return t.SeasonalAnimeApply })
	return s.list(ctx, fmt.Sprintf("anime/season/%d/%s", year, season), rawOptions...)
}

// MARK: Suggested
// Suggested returns suggested anime for the authorized user. If the user is new
// comer, this endpoint returns an empty list.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/anime_suggestions_get
func (s *AnimeService) Suggested(ctx context.Context, options ...prm.OptionalParam) ([]containers.Anime, *api_driver.Response, error) {
	rawOptions := optionsToFuncs(options, func(t prm.OptionalParam) func(*url.Values) { return t.Apply })
	return s.list(ctx, "anime/suggestions", rawOptions...)
}

// MARK: UpdateMyListStatus
// UpdateMyListStatus adds the anime specified by animeID to the user's anime
// list with one or more options added to update the status. If the anime
// already exists in the list, only the status is updated.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/anime_anime_id_my_list_status_put
func (s *AnimeService) UpdateMyListStatus(ctx context.Context, animeID int, options ...prm.UpdateMyAnimeListStatusOption) (*containers.AnimeListStatus, *api_driver.Response, error) {
	a := new(containers.AnimeListStatus)
	rawOptions := optionsToFuncs(options, func(t prm.UpdateMyAnimeListStatusOption) func(*url.Values) { return t.UpdateMyAnimeListStatusApply })
	resp, err := s.client.UpdateMyListStatus(ctx, "anime", animeID, a, rawOptions...)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

// MARK: DeleteMyListItem
// DeleteMyListItem deletes an anime from the user's list. If the anime does not
// exist in the user's list, 404 Not Found error is returned.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/anime_anime_id_my_list_status_delete
func (s *AnimeService) DeleteMyListItem(ctx context.Context, animeID int) (*api_driver.Response, error) {
	return s.client.DeleteMyListItem(ctx, "anime", animeID)
}

func (s *AnimeService) list(ctx context.Context, path string, options ...func(v *url.Values)) ([]containers.Anime, *api_driver.Response, error) {
	list, resp, err := s.client.RequestAnimeList(ctx, path, options...)
	if err != nil {
		return nil, resp, err
	}
	anime := make([]containers.Anime, len(list))
	for i := range list {
		anime[i] = list[i].Anime
	}
	return anime, resp, nil
}
