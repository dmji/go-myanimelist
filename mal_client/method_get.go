package mal_client

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal_type"
)

// RequestAnimeList sends a GET request to the specified URL.
func (c *Client) RequestAnimeList(ctx context.Context, path string, options ...func(v *url.Values)) ([]mal_type.UserAnime, *Response, error) {
	p := new(listWithPagination[[]mal_type.UserAnime])
	resp, err := c.requestPagedItem(ctx, path, p, options...)
	if err != nil {
		return nil, resp, err
	}
	return p.Data, resp, nil
}

// RequestMangaList sends a GET request to the specified URL.
func (c *Client) RequestMangaList(ctx context.Context, path string, options ...func(v *url.Values)) ([]mal_type.UserManga, *Response, error) {
	p := new(listWithPagination[[]mal_type.UserManga])
	resp, err := c.requestPagedItem(ctx, path, p, options...)
	if err != nil {
		return nil, resp, err
	}
	return p.Data, resp, nil
}

// RequestTopicDetails sends a GET request to the specified URL.
func (c *Client) RequestTopicDetails(ctx context.Context, path string, options ...func(v *url.Values)) (mal_type.TopicDetails, *Response, error) {
	p := new(listWithPagination[mal_type.TopicDetails])

	resp, err := c.requestPagedItem(ctx, path, p, options...)
	if err != nil {
		return mal_type.TopicDetails{}, resp, err
	}
	return p.Data, resp, nil
}

// RequestTopics sends a GET request to the specified URL.
func (c *Client) RequestTopics(ctx context.Context, path string, options ...func(v *url.Values)) ([]mal_type.Topic, *Response, error) {
	p := new(listWithPagination[[]mal_type.Topic])
	resp, err := c.requestPagedItem(ctx, path, p, options...)
	if err != nil {
		return nil, resp, err
	}
	return p.Data, resp, nil
}
