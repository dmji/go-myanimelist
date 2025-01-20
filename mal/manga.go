package mal

import (
	"context"
	"fmt"
	"net/url"

	"github.com/dmji/go-myanimelist/mal_client"
	"github.com/dmji/go-myanimelist/mal_opt"
	"github.com/dmji/go-myanimelist/mal_type"
)

type clientManga interface {
	RequestGet(ctx context.Context, path string, v interface{}, options ...func(v *url.Values)) (*mal_client.Response, error)
	UpdateMyListStatus(ctx context.Context, path string, id int, v interface{}, options ...func(v *url.Values)) (*mal_client.Response, error)
	DeleteMyListItem(ctx context.Context, path string, animeID int) (*mal_client.Response, error)
	RequestMangaList(ctx context.Context, path string, options ...func(v *url.Values)) ([]mal_type.UserManga, *mal_client.Response, error)
}

// MangaService handles communication with the manga related methods of the
// MyAnimeList API:
//
// https://myanimelist.net/apiconfig/references/api/v2#tag/manga
// https://myanimelist.net/apiconfig/references/api/v2#tag/user-mangalist
type MangaService struct {
	client clientManga

	DetailsOptions            mal_opt.DetailsOptionProvider
	ListOptions               mal_opt.OptionalParamProvider
	RankingOptions            mal_opt.OptionalParamProvider
	UpdateMyListStatusOptions mal_opt.UpdateMyMangaListStatusOptionProvider
}

// NewMangaService returns a new MangaService.
func NewMangaService(client clientManga) *MangaService {
	return &MangaService{
		client: client,
	}
}

// List allows an authenticated user to search and list manga data. You may get
// user specific data by using the optional field.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/manga_get
func (s *MangaService) List(ctx context.Context, search string, options ...mal_opt.OptionalParam) ([]mal_type.Manga, *mal_client.Response, error) {
	options = append(options, withOptionFromQuery(search))
	return s.list(ctx, "manga", options...)
}

// Details returns details about a manga. By default, few manga fields are
// populated. Use the Fields option to specify which fields should be included.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/manga_manga_id_get
func (s *MangaService) Details(ctx context.Context, mangaID int, options ...mal_opt.DetailsOption) (*mal_type.Manga, *mal_client.Response, error) {
	m := new(mal_type.Manga)
	rawOptions := detailsOptionsToFuncs(options)
	resp, err := s.client.RequestGet(ctx, fmt.Sprintf("manga/%d", mangaID), m, rawOptions...)
	if err != nil {
		return nil, resp, err
	}
	return m, resp, nil
}

// Ranking allows an authenticated user to receive the top manga based on a
// certain ranking.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/manga_ranking_get
func (s *MangaService) Ranking(ctx context.Context, ranking mal_opt.MangaRanking, options ...mal_opt.OptionalParam) ([]mal_type.Manga, *mal_client.Response, error) {
	options = append(
		options,
		mal_opt.OptionFunc(func(v *url.Values) {
			v.Set("ranking_type", string(ranking))
		}))
	return s.list(ctx, "manga/ranking", options...)
}

// UpdateMyListStatus adds the manga specified by mangaID to the user's manga
// list with one or more options added to update the status. If the manga
// already exists in the list, only the status is updated.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/manga_manga_id_my_list_status_put
func (s *MangaService) UpdateMyListStatus(ctx context.Context, mangaID int, options ...mal_opt.UpdateMyMangaListStatusOption) (*mal_type.MangaListStatus, *mal_client.Response, error) {
	rawOptions := optionsToFuncs(options, func(t mal_opt.UpdateMyMangaListStatusOption) func(*url.Values) { return t.UpdateMyMangaListStatusApply })

	m := new(mal_type.MangaListStatus)
	resp, err := s.client.UpdateMyListStatus(ctx, "manga", mangaID, m, rawOptions...)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, nil
}

// DeleteMyListItem deletes a manga from the user's list. If the manga does not
// exist in the user's list, 404 Not Found error is returned.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/manga_manga_id_my_list_status_delete
func (s *MangaService) DeleteMyListItem(ctx context.Context, mangaID int) (*mal_client.Response, error) {
	return s.client.DeleteMyListItem(ctx, "manga", mangaID)
}

func (s *MangaService) list(ctx context.Context, path string, options ...mal_opt.OptionalParam) ([]mal_type.Manga, *mal_client.Response, error) {
	rawOptions := optionsToFuncs(options, func(t mal_opt.OptionalParam) func(*url.Values) { return t.Apply })
	mangaList, resp, err := s.client.RequestMangaList(ctx, path, rawOptions...)
	if err != nil {
		return nil, resp, err
	}
	manga := make([]mal_type.Manga, len(mangaList))
	for i := range mangaList {
		manga[i] = mangaList[i].Manga
	}
	return manga, resp, nil
}
