//go:build !debug

package util

import (
	"net/http"
)

func DumpRequest(req *http.Request) {}

func DumpResponse(resp *http.Response) {}
