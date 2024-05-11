package mal

import (
	"context"
	"fmt"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/malhttp"
	"github.com/dmji/go-myanimelist/mal/maltype"
	"github.com/dmji/go-myanimelist/mal/prm"
)

// UserService handles communication with the user related methods of the
// MyAnimeList API:
//
// https://myanimelist.net/apiconfig/references/api/v2#tag/user
// https://myanimelist.net/apiconfig/references/api/v2#operation/users_user_id_animelist_get
// https://myanimelist.net/apiconfig/references/api/v2#operation/users_user_id_mangalist_get
type UserService struct {
	client *malhttp.Client

	AnimeListOptions prm.AnimeListOptionProvider
	MangaListOptions prm.MangaListOptionProvider
	MyInfoOptions    prm.MyInfoOptionProvider
}

func NewUserService(client *malhttp.Client) *UserService {
	return &UserService{
		client: client,
	}
}

// MyInfo returns information about the authenticated user.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/users_user_id_get
func (s *UserService) MyInfo(ctx context.Context, options ...prm.MyInfoOption) (*maltype.User, *malhttp.Response, error) {
	u := new(maltype.User)
	rawOptions := optionsToFuncs(options, func(t prm.MyInfoOption) func(*url.Values) { return t.MyInfoApply })
	resp, err := s.client.RequestGet(ctx, "users/@me", u, rawOptions...)
	if err != nil {
		return nil, resp, err
	}
	return u, resp, err
}

// MangaList gets the manga list of the user indicated by username (or use @me).
// The manga can be sorted and filtered using the MangaStatus and SortMangaList
// option functions respectively.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/users_user_id_mangalist_get
func (s *UserService) MangaList(ctx context.Context, username string, options ...prm.MangaListOption) ([]maltype.UserManga, *malhttp.Response, error) {
	rawOptions := optionsToFuncs(options, func(t prm.MangaListOption) func(*url.Values) { return t.MangaListApply })
	list, resp, err := s.client.RequestMangaList(ctx, fmt.Sprintf("users/%s/mangalist", username), rawOptions...)
	if err != nil {
		return nil, resp, err
	}
	return list, resp, nil
}

// AnimeList gets the anime list of the user indicated by username (or use @me).
// The anime can be sorted and filtered using the AnimeStatus and SortAnimeList
// option functions respectively.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/users_user_id_animelist_get
func (s *UserService) AnimeList(ctx context.Context, username string, options ...prm.AnimeListOption) ([]maltype.UserAnime, *malhttp.Response, error) {
	rawOptions := optionsToFuncs(options, func(t prm.AnimeListOption) func(*url.Values) { return t.AnimeListApply })
	list, resp, err := s.client.RequestAnimeList(ctx, fmt.Sprintf("users/%s/animelist", username), rawOptions...)
	if err != nil {
		return nil, resp, err
	}
	return list, resp, nil
}
