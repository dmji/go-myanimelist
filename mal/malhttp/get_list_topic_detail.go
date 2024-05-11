package malhttp

import (
	"context"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/maltype"
)

// Topic Details
type topicDetail listWithPagination[maltype.TopicDetails]

func (t topicDetail) pagination() paging { return t.Paging }

// RequestTopicDetails sends a GET request to the specified URL.
func (c *Client) RequestTopicDetails(ctx context.Context, path string, options ...func(v *url.Values)) (maltype.TopicDetails, *Response, error) {
	p := new(topicDetail)
	resp, err := c.requestPagedItem(ctx, path, p, options...)
	if err != nil {
		return maltype.TopicDetails{}, resp, err
	}
	return p.Data, resp, nil
}
