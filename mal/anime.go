package mal

import (
	"context"
	"fmt"
	"net/url"

	"github.com/dmji/go-myanimelist/mal_client"
	"github.com/dmji/go-myanimelist/mal_opt"
	"github.com/dmji/go-myanimelist/mal_type"
)

type clientAnime interface {
	RequestGet(ctx context.Context, path string, v interface{}, options ...func(v *url.Values)) (*mal_client.Response, error)
	UpdateMyListStatus(ctx context.Context, path string, id int, v interface{}, options ...func(v *url.Values)) (*mal_client.Response, error)
	DeleteMyListItem(ctx context.Context, path string, animeID int) (*mal_client.Response, error)
	RequestAnimeList(ctx context.Context, path string, options ...func(v *url.Values)) ([]mal_type.UserAnime, *mal_client.Response, error)
}

// AnimeService handles communication with the anime related methods of the
// MyAnimeList API:
//
// https://myanimelist.net/apiconfig/references/api/v2#tag/anime
// https://myanimelist.net/apiconfig/references/api/v2#tag/user-animelist
type AnimeService struct {
	client clientAnime

	DetailsOptions            mal_opt.DetailsOptionProvider
	ListOptions               mal_opt.OptionalParamProvider
	SuggestedOptions          mal_opt.OptionalParamProvider
	RankingOptions            mal_opt.OptionalParamProvider
	UpdateMyListStatusOptions mal_opt.UpdateMyAnimeListStatusOptionProvider
	SeasonalOptions           mal_opt.SeasonalAnimeOptionProvider
}

// NewAnimeService returns a new AnimeService.
func NewAnimeService(client clientAnime) *AnimeService {
	return &AnimeService{
		client: client,
	}
}

// MARK: List

// List allows an authenticated user to search and list anime data. You may get
// user specific data by using the optional field.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/anime_get
func (s *AnimeService) List(ctx context.Context, search string, options ...mal_opt.OptionalParam) ([]mal_type.Anime, *mal_client.Response, error) {
	options = append(options, withOptionFromQuery(search))
	rawOptions := optionsToFuncs(options, func(t mal_opt.OptionalParam) func(*url.Values) { return t.Apply })
	return s.list(ctx, "anime", rawOptions...)
}

// MARK: Details

// Details returns details about an anime. By default, few anime fields are
// populated. Use the Fields option to specify which fields should be included.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/anime_anime_id_get
func (s *AnimeService) Details(ctx context.Context, animeID int, options ...mal_opt.DetailsOption) (*mal_type.Anime, *mal_client.Response, error) {
	a := new(mal_type.Anime)
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
func (s *AnimeService) Ranking(ctx context.Context, ranking mal_opt.AnimeRanking, options ...mal_opt.OptionalParam) ([]mal_type.Anime, *mal_client.Response, error) {
	options = append(
		options,
		mal_opt.OptionFunc(func(v *url.Values) {
			v.Set("ranking_type", string(ranking))
		}),
	)
	rawOptions := optionsToFuncs(options, func(t mal_opt.OptionalParam) func(*url.Values) { return t.Apply })
	return s.list(ctx, "anime/ranking", rawOptions...)
}

// MARK: Seasonal

// Seasonal allows an authenticated user to receive the seasonal anime by
// providing the year and season. The results can be sorted using an option.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/anime_season_year_season_get
func (s *AnimeService) Seasonal(ctx context.Context, year int, season mal_opt.AnimeSeason, options ...mal_opt.SeasonalAnimeOption) ([]mal_type.Anime, *mal_client.Response, error) {
	rawOptions := optionsToFuncs(options, func(t mal_opt.SeasonalAnimeOption) func(*url.Values) { return t.SeasonalAnimeApply })
	return s.list(ctx, fmt.Sprintf("anime/season/%d/%s", year, season), rawOptions...)
}

// MARK: Suggested

// Suggested returns suggested anime for the authorized user. If the user is new
// comer, this endpoint returns an empty list.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/anime_suggestions_get
func (s *AnimeService) Suggested(ctx context.Context, options ...mal_opt.OptionalParam) ([]mal_type.Anime, *mal_client.Response, error) {
	rawOptions := optionsToFuncs(options, func(t mal_opt.OptionalParam) func(*url.Values) { return t.Apply })
	return s.list(ctx, "anime/suggestions", rawOptions...)
}

// MARK: UpdateMyListStatus

// UpdateMyListStatus adds the anime specified by animeID to the user's anime
// list with one or more options added to update the status. If the anime
// already exists in the list, only the status is updated.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/anime_anime_id_my_list_status_put
func (s *AnimeService) UpdateMyListStatus(ctx context.Context, animeID int, options ...mal_opt.UpdateMyAnimeListStatusOption) (*mal_type.AnimeListStatus, *mal_client.Response, error) {
	a := new(mal_type.AnimeListStatus)
	rawOptions := optionsToFuncs(options, func(t mal_opt.UpdateMyAnimeListStatusOption) func(*url.Values) { return t.UpdateMyAnimeListStatusApply })
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
func (s *AnimeService) DeleteMyListItem(ctx context.Context, animeID int) (*mal_client.Response, error) {
	return s.client.DeleteMyListItem(ctx, "anime", animeID)
}

func (s *AnimeService) list(ctx context.Context, path string, options ...func(v *url.Values)) ([]mal_type.Anime, *mal_client.Response, error) {
	list, resp, err := s.client.RequestAnimeList(ctx, path, options...)
	if err != nil {
		return nil, resp, err
	}
	anime := make([]mal_type.Anime, len(list))
	for i := range list {
		anime[i] = list[i].Anime
	}
	return anime, resp, nil
}
