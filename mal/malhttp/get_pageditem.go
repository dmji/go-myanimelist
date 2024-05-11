package malhttp

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// MARK: List / pagination template
type listWithPagination[T any] struct {
	Data   T      `json:"data"`
	Paging paging `json:"paging"`
}

type pagination interface {
	pagination() paging
}

func (c *Client) requestPagedItem(ctx context.Context, path string, p pagination, options ...func(v *url.Values)) (*Response, error) {
	req, err := c.NewRequest(http.MethodGet, path)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	fillValues(&q, options...)
	req.URL.RawQuery = q.Encode()

	resp, err := c.Do(ctx, req, p)
	if err != nil {
		return resp, err
	}

	prev, next, err := parsePaging(p.pagination())
	if err != nil {
		return resp, err
	}
	resp.PrevOffset = prev
	resp.NextOffset = next

	return resp, nil
}

func parsePaging(p paging) (prev, next int, err error) {
	if p.Previous != "" {
		offset, err := parseOffset(p.Previous)
		if err != nil {
			return 0, 0, fmt.Errorf("paging: previous: %s", err)
		}
		prev = offset
	}
	if p.Next != "" {
		offset, err := parseOffset(p.Next)
		if err != nil {
			return 0, 0, fmt.Errorf("paging: next: %s", err)
		}
		next = offset
	}
	return prev, next, nil
}

func parseOffset(urlStr string) (int, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return 0, fmt.Errorf("parsing URL %q: %s", urlStr, err)
	}
	offset, err := strconv.Atoi(u.Query().Get("offset"))
	if err != nil {
		return 0, fmt.Errorf("parsing offset: %s", err)
	}
	return offset, nil
}
