package mal_client

import (
	"context"

	"github.com/dmji/go-myanimelist/mal_prm"
	"github.com/dmji/go-myanimelist/mal_type"
)

// RequestAnimeList sends a GET request to the specified URL.
func (c *Client) RequestAnimeList(ctx context.Context, path string, opts *mal_prm.UserAnimeListRequestParameters) ([]mal_type.UserAnime, *Response, error) {
	p := new(listWithPagination[[]mal_type.UserAnime])
	resp, err := c.requestPagedItem(ctx, path, p, opts)
	if err != nil {
		return nil, resp, err
	}
	return p.Data, resp, nil
}

// RequestMangaList sends a GET request to the specified URL.
func (c *Client) RequestMangaList(ctx context.Context, path string, opts *mal_prm.UserMangaListRequestParameters) ([]mal_type.UserManga, *Response, error) {
	p := new(listWithPagination[[]mal_type.UserManga])
	resp, err := c.requestPagedItem(ctx, path, p, opts)
	if err != nil {
		return nil, resp, err
	}
	return p.Data, resp, nil
}

// RequestTopicDetails sends a GET request to the specified URL.
func (c *Client) RequestTopicDetails(ctx context.Context, path string, opts *mal_prm.ForumTopicDetailsRequestParameters) (mal_type.TopicDetails, *Response, error) {
	p := new(listWithPagination[mal_type.TopicDetails])

	resp, err := c.requestPagedItem(ctx, path, p, opts)
	if err != nil {
		return mal_type.TopicDetails{}, resp, err
	}
	return p.Data, resp, nil
}

// RequestTopics sends a GET request to the specified URL.
func (c *Client) RequestTopics(ctx context.Context, path string, opts *mal_prm.ForumTopicsRequestParameters) ([]mal_type.Topic, *Response, error) {
	p := new(listWithPagination[[]mal_type.Topic])
	resp, err := c.requestPagedItem(ctx, path, p, opts)
	if err != nil {
		return nil, resp, err
	}
	return p.Data, resp, nil
}
