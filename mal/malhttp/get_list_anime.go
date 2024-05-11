package malhttp

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/maltype"
)

// animeList represents the anime list of a user.
type animeList listWithPagination[[]maltype.UserAnime]

func (a animeList) pagination() paging { return a.Paging }

func (c *Client) RequestAnimeList(ctx context.Context, path string, options ...func(v *url.Values)) ([]maltype.UserAnime, *Response, error) {
	p := new(animeList)
	resp, err := c.requestPagedItem(ctx, path, p, options...)
	if err != nil {
		return nil, resp, err
	}
	return p.Data, resp, nil
}
