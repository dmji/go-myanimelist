package anime

import (
	"context"
	"fmt"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/api_driver"
	"github.com/dmji/go-myanimelist/mal/common"
	"github.com/dmji/go-myanimelist/mal/containers"
	"github.com/dmji/go-myanimelist/mal/prm"
)

// Service handles communication with the anime related methods of the
// MyAnimeList API:
//
// https://myanimelist.net/apiconfig/references/api/v2#tag/anime
// https://myanimelist.net/apiconfig/references/api/v2#tag/user-animelist
type Service struct {
	client *api_driver.Client
}

func NewService(client *api_driver.Client) *Service {
	return &Service{
		client: client,
	}
}

// Details returns details about an anime. By default, few anime fields are
// populated. Use the Fields option to specify which fields should be included.
func (s *Service) Details(ctx context.Context, animeID int, options ...prm.DetailsOption) (*containers.Anime, *common.Response, error) {
	a := new(containers.Anime)
	rawOptions := common.OptionsToFuncs(options, func(t prm.DetailsOption) func(*url.Values) { return t.DetailsApply })
	resp, err := s.client.RequestGet(ctx, fmt.Sprintf("anime/%d", animeID), a, rawOptions...)
	if err != nil {
		return nil, resp, err
	}
	return a, resp, nil
}

// List allows an authenticated user to search and list anime data. You may get
// user specific data by using the optional field "my_list_status".
func (s *Service) List(ctx context.Context, search string, options ...prm.OptionalParam) ([]containers.Anime, *common.Response, error) {
	options = append(options, common.OptionFromQuery(search))
	return s.list(ctx, "anime", options...)
}

func (s *Service) list(ctx context.Context, path string, options ...prm.OptionalParam) ([]containers.Anime, *common.Response, error) {
	rawOptions := common.OptionsToFuncs(options, func(t prm.OptionalParam) func(*url.Values) { return t.Apply })
	list, resp, err := s.client.RequestAnimeList(ctx, path, rawOptions...)
	if err != nil {
		return nil, resp, err
	}
	anime := make([]containers.Anime, len(list))
	for i := range list {
		anime[i] = list[i].Anime
	}
	return anime, resp, nil
}

// Suggested returns suggested anime for the authorized user. If the user is new
// comer, this endpoint returns an empty list.
func (s *Service) Suggested(ctx context.Context, options ...prm.OptionalParam) ([]containers.Anime, *common.Response, error) {
	return s.list(ctx, "anime/suggestions", options...)
}

// MARK: DeleteMyListItem
// DeleteMyListItem deletes an anime from the user's list. If the anime does not
// exist in the user's list, 404 Not Found error is returned.
func (s *Service) DeleteMyListItem(ctx context.Context, animeID int) (*common.Response, error) {
	return s.client.DeleteMyListItem(ctx, "anime", animeID)
}
