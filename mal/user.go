package mal

import (
	"context"
	"fmt"
	"net/url"

	"github.com/dmji/go-myanimelist/mal_client"
	"github.com/dmji/go-myanimelist/mal_opt"
	"github.com/dmji/go-myanimelist/mal_type"
)

type clientUser interface {
	RequestGet(ctx context.Context, path string, v interface{}, options ...func(v *url.Values)) (*mal_client.Response, error)
	RequestMangaList(ctx context.Context, path string, options ...func(v *url.Values)) ([]mal_type.UserManga, *mal_client.Response, error)
	RequestAnimeList(ctx context.Context, path string, options ...func(v *url.Values)) ([]mal_type.UserAnime, *mal_client.Response, error)
}

// UserService handles communication with the user related methods of the
// MyAnimeList API:
//
// https://myanimelist.net/apiconfig/references/api/v2#tag/user
// https://myanimelist.net/apiconfig/references/api/v2#operation/users_user_id_animelist_get
// https://myanimelist.net/apiconfig/references/api/v2#operation/users_user_id_mangalist_get
type UserService struct {
	client clientUser

	AnimeListOptions mal_opt.AnimeListOptionProvider
	MangaListOptions mal_opt.MangaListOptionProvider
	MyInfoOptions    mal_opt.MyInfoOptionProvider
}

// NewUserService returns a new UserService.
func NewUserService(client clientUser) *UserService {
	return &UserService{
		client: client,
	}
}

// MyInfo returns information about the authenticated user.
// Reference API docs: https://myanimelist.net/apiconfig/references/api/v2#operation/users_user_id_get
func (s *UserService) MyInfo(ctx context.Context, options ...mal_opt.MyInfoOption) (*mal_type.User, *mal_client.Response, error) {
	u := new(mal_type.User)
	rawOptions := optionsToFuncs(options, func(t mal_opt.MyInfoOption) func(*url.Values) { return t.MyInfoApply })
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
func (s *UserService) MangaList(ctx context.Context, username string, options ...mal_opt.MangaListOption) ([]mal_type.UserManga, *mal_client.Response, error) {
	rawOptions := optionsToFuncs(options, func(t mal_opt.MangaListOption) func(*url.Values) { return t.MangaListApply })
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
func (s *UserService) AnimeList(ctx context.Context, username string, options ...mal_opt.AnimeListOption) ([]mal_type.UserAnime, *mal_client.Response, error) {
	rawOptions := optionsToFuncs(options, func(t mal_opt.AnimeListOption) func(*url.Values) { return t.AnimeListApply })
	list, resp, err := s.client.RequestAnimeList(ctx, fmt.Sprintf("users/%s/animelist", username), rawOptions...)
	if err != nil {
		return nil, resp, err
	}
	return list, resp, nil
}
