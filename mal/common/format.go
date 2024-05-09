package common

import (
	"fmt"
	"net/http"
	"time"
)

// Paging provides access to the next and previous page URLs when there are
// pages of results.
type Paging struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

// MARK: Date / Time format
func FormatMALDate(d time.Time) string {
	if d.IsZero() {
		return ""
	}
	return d.Format("2006-01-02")
}

// MARK: Error format
// An ErrorResponse reports an error caused by an API request.
//
// https://myanimelist.net/apiconfig/references/api/v2#section/Common-formats - Error format
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"`
	Err      string         `json:"error"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message, r.Err)
}
