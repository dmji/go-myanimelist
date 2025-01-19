package mal_client

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal_type"
)

// Topics
type topics listWithPagination[[]mal_type.Topic]

func (t topics) pagination() paging { return t.Paging }

// RequestTopics sends a GET request to the specified URL.
func (c *Client) RequestTopics(ctx context.Context, path string, options ...func(v *url.Values)) ([]mal_type.Topic, *Response, error) {
	p := new(topics)
	resp, err := c.requestPagedItem(ctx, path, p, options...)
	if err != nil {
		return nil, resp, err
	}
	return p.Data, resp, nil
}
