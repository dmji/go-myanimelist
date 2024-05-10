package api_driver

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/containers"
)

// Topics
type topics listWithPagination[[]containers.Topic]

func (t topics) pagination() paging { return t.Paging }

func (c *Client) RequestTopics(ctx context.Context, path string, options ...func(v *url.Values)) ([]containers.Topic, *Response, error) {
	p := new(topics)
	resp, err := c.requestPagedItem(ctx, path, p, options...)
	if err != nil {
		return nil, resp, err
	}
	return p.Data, resp, nil
}
