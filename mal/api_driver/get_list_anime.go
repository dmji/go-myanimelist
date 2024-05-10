package api_driver

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/containers"
)

// animeList represents the anime list of a user.
type animeList listWithPagination[[]containers.UserAnime]

func (a animeList) pagination() paging { return a.Paging }

func (c *Client) RequestAnimeList(ctx context.Context, path string, options ...func(v *url.Values)) ([]containers.UserAnime, *Response, error) {
	p := new(animeList)
	resp, err := c.requestPagedItem(ctx, path, p, options...)
	if err != nil {
		return nil, resp, err
	}
	return p.Data, resp, nil
}
