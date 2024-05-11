//go:build !debug

package malhttp

import (
	"net/http"
)

func DumpRequest(req *http.Request) {}

func DumpResponse(resp *http.Response) {}
