package mal_client

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal_type"
)

// animeList represents the anime list of a user.
type animeList listWithPagination[[]mal_type.UserAnime]

func (a animeList) pagination() paging { return a.Paging }

// RequestAnimeList sends a GET request to the specified URL.
func (c *Client) RequestAnimeList(ctx context.Context, path string, options ...func(v *url.Values)) ([]mal_type.UserAnime, *Response, error) {
	p := new(animeList)
	resp, err := c.requestPagedItem(ctx, path, p, options...)
	if err != nil {
		return nil, resp, err
	}
	return p.Data, resp, nil
}
