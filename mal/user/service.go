package user

import (
	"context"
	"fmt"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/api_driver"
	"github.com/dmji/go-myanimelist/mal/common"
	"github.com/dmji/go-myanimelist/mal/containers"
	"github.com/dmji/go-myanimelist/mal/prm"
)

// UserService handles communication with the user related methods of the
// MyAnimeList API:
//
// https://myanimelist.net/apiconfig/references/api/v2#tag/user
// https://myanimelist.net/apiconfig/references/api/v2#operation/users_user_id_animelist_get
// https://myanimelist.net/apiconfig/references/api/v2#operation/users_user_id_mangalist_get
type Service struct {
	client *api_driver.Client
}

func NewService(client *api_driver.Client) *Service {
	return &Service{
		client: client,
	}
}

// MyInfo returns information about the authenticated user.
func (s *Service) MyInfo(ctx context.Context, options ...prm.MyInfoOption) (*containers.User, *common.Response, error) {
	u := new(containers.User)
	rawOptions := common.OptionsToFuncs(options, func(t prm.MyInfoOption) func(*url.Values) { return t.MyInfoApply })
	resp, err := s.client.RequestGet(ctx, "users/@me", u, rawOptions...)
	if err != nil {
		return nil, resp, err
	}
	return u, resp, err
}

// MangaList gets the manga list of the user indicated by username (or use @me).
// The manga can be sorted and filtered using the MangaStatus and SortMangaList
// option functions respectively.
func (s *Service) MangaList(ctx context.Context, username string, options ...prm.MangaListOption) ([]containers.UserManga, *common.Response, error) {
	rawOptions := common.OptionsToFuncs(options, func(t prm.MangaListOption) func(*url.Values) { return t.MangaListApply })
	list, resp, err := s.client.RequestMangaList(ctx, fmt.Sprintf("users/%s/mangalist", username), rawOptions...)
	if err != nil {
		return nil, resp, err
	}
	return list, resp, nil
}

// AnimeList gets the anime list of the user indicated by username (or use @me).
// The anime can be sorted and filtered using the AnimeStatus and SortAnimeList
// option functions respectively.
func (s *Service) AnimeList(ctx context.Context, username string, options ...prm.AnimeListOption) ([]containers.UserAnime, *common.Response, error) {
	rawOptions := common.OptionsToFuncs(options, func(t prm.AnimeListOption) func(*url.Values) { return t.AnimeListApply })
	list, resp, err := s.client.RequestAnimeList(ctx, fmt.Sprintf("users/%s/animelist", username), rawOptions...)
	if err != nil {
		return nil, resp, err
	}
	return list, resp, nil
}
