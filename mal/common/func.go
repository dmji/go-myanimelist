package common

import (
	"net/http"
	"net/url"
)

type OptionFunc func(v *url.Values)

func (f OptionFunc) Apply(v *url.Values) {
	f(v)
}

func OptionFromQuery(query string) OptionFunc {
	return OptionFunc(func(v *url.Values) {
		v.Set("q", query)
	})
}

// Response wraps http.Response and is returned in all the library functions
// that communicate with the MyAnimeList API. Even if an error occurs the
// response will always be returned along with the actual error so that the
// caller can further inspect it if needed. For the same reason it also keeps
// a copy of the http.Response.Body that was read when the response was first
// received.
type Response struct {
	*http.Response
	Body []byte

	NextOffset int
	PrevOffset int
}

func OptionsToFuncs[T any](options []T, fn func(t T) func(*url.Values)) []func(v *url.Values) {
	rawOptions := make([]func(v *url.Values), len(options))
	for i := range options {
		rawOptions[i] = fn(options[i])
	}
	return rawOptions
}
