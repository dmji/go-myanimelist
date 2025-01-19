package mal_client

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal_type"
)

// Topic Details
type topicDetail listWithPagination[mal_type.TopicDetails]

func (t topicDetail) pagination() paging { return t.Paging }

// RequestTopicDetails sends a GET request to the specified URL.
func (c *Client) RequestTopicDetails(ctx context.Context, path string, options ...func(v *url.Values)) (mal_type.TopicDetails, *Response, error) {
	p := new(topicDetail)
	resp, err := c.requestPagedItem(ctx, path, p, options...)
	if err != nil {
		return mal_type.TopicDetails{}, resp, err
	}
	return p.Data, resp, nil
}
