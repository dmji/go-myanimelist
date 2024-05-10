package api_driver

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/containers"
)

// mangaList represents the anime list of a user.
type mangaList listWithPagination[[]containers.UserManga]

func (m mangaList) pagination() paging { return m.Paging }

func (c *Client) RequestMangaList(ctx context.Context, path string, options ...func(v *url.Values)) ([]containers.UserManga, *Response, error) {
	p := new(mangaList)
	resp, err := c.requestPagedItem(ctx, path, p, options...)
	if err != nil {
		return nil, resp, err
	}
	return p.Data, resp, nil
}
