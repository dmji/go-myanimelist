package api_driver

import (
	"context"
	"net/http"
	"net/url"

	"github.com/dmji/go-myanimelist/mal/common"
)

func (c *Client) RequestGet(ctx context.Context, path string, v interface{}, options ...func(v *url.Values)) (*common.Response, error) {
	req, err := c.NewRequest(http.MethodGet, path)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	fillValues(&q, options...)
	req.URL.RawQuery = q.Encode()

	resp, err := c.Do(ctx, req, v)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
