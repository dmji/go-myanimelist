package mal_client

import (
	"context"
	"net/http"
)

// RequestGet sends a GET request to the specified URL.
func (c *Client) RequestGet(ctx context.Context, path string, v interface{}, qdata interface{}) (*Response, error) {
	req, err := c.NewRequest(http.MethodGet, path)
	if err != nil {
		return nil, err
	}

	if qdata != nil {
		req.URL.RawQuery, err = c.urlMarshaler.Marshal(qdata)
		if err != nil {
			return nil, err
		}
	}

	resp, err := c.Do(ctx, req, v)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
