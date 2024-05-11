//go:build !debug

package malhttp

import (
	"net/http"
)

// DumpRequest Stub functions for debugging
func DumpRequest(req *http.Request) {}

// DumpResponse Stub functions for debugging
func DumpResponse(resp *http.Response) {}
