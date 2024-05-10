//go:build !debug

package api_driver

import (
	"net/http"
)

func DumpRequest(req *http.Request) {}

func DumpResponse(resp *http.Response) {}
