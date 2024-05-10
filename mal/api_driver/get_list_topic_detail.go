package api_driver

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/containers"
)

// Topic Details
type topicDetail listWithPagination[containers.TopicDetails]

func (t topicDetail) pagination() paging { return t.Paging }

func (c *Client) RequestTopicDetails(ctx context.Context, path string, options ...func(v *url.Values)) (containers.TopicDetails, *Response, error) {
	p := new(topicDetail)
	resp, err := c.requestPagedItem(ctx, path, p, options...)
	if err != nil {
		return containers.TopicDetails{}, resp, err
	}
	return p.Data, resp, nil
}
