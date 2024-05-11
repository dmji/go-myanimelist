package malhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v. If v implements the
// io.Writer interface, the raw response body will be written to v, without
// attempting to first decode it.
//
// If the provided ctx is nil then an error will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	if ctx == nil {
		return nil, errors.New("context must not be nil")
	}
	req = req.WithContext(ctx)

	DumpRequest(req)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	DumpResponse(resp)

	response := &Response{Response: resp}
	if err := checkResponse(resp); err != nil {
		return response, err
	}

	if v != nil {
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}

	return response, err
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		// Ignore unmarshal error for undocumented error formats or HTML.
		_ = json.Unmarshal(data, errorResponse)
	}
	// Re-populate error response body in case JSON unmarshal fails.
	r.Body = io.NopCloser(bytes.NewBuffer(data))

	return errorResponse
}
