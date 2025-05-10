package mal

import (
	"net/url"

	"github.com/dmji/go-myanimelist/mal_opt"
)

func optionsToFuncs[T any](options []T, fn func(t T) func(*url.Values)) []func(v *url.Values) {
	rawOptions := make([]func(v *url.Values), len(options))
	for i := range options {
		rawOptions[i] = fn(options[i])
	}
	return rawOptions
}

func withOptionFromQuery(query string) mal_opt.OptionFunc {
	return mal_opt.OptionFunc(func(v *url.Values) {
		v.Set("q", query)
	})
}
