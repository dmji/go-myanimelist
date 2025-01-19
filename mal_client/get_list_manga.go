package mal_client

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal_type"
)

// mangaList represents the anime list of a user.
type mangaList listWithPagination[[]mal_type.UserManga]

func (m mangaList) pagination() paging { return m.Paging }

// RequestMangaList sends a GET request to the specified URL.
func (c *Client) RequestMangaList(ctx context.Context, path string, options ...func(v *url.Values)) ([]mal_type.UserManga, *Response, error) {
	p := new(mangaList)
	resp, err := c.requestPagedItem(ctx, path, p, options...)
	if err != nil {
		return nil, resp, err
	}
	return p.Data, resp, nil
}
