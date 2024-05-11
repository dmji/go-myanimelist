//go:build debug

package malhttp

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// DumpRequest  functions for debugging
func DumpRequest(req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Printf("request dump failed: %s", err)
		return
	}
	fmt.Println("---------------- Request dump -----------------")
	fmt.Println(string(dump))
	fmt.Println("")
}

// DumpResponse  functions for debugging
func DumpResponse(resp *http.Response) {
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Printf("response dump failed: %s", err)
		return
	}
	fmt.Println("---------------- Response dump ----------------")
	fmt.Println(string(dump))
	fmt.Println("")
	fmt.Println("-----------------------------------------------")
}
