package mal

import (
	"net/url"

	"github.com/dmji/go-myanimelist/mal/prm"
)

func OptionsToFuncs[T any](options []T, fn func(t T) func(*url.Values)) []func(v *url.Values) {
	rawOptions := make([]func(v *url.Values), len(options))
	for i := range options {
		rawOptions[i] = fn(options[i])
	}
	return rawOptions
}

func DetailsOptionsToFuncs(options []prm.DetailsOption) []func(v *url.Values) {
	fn := func(t prm.DetailsOption) func(*url.Values) { return t.DetailsApply }

	rawOptions := make([]func(v *url.Values), len(options))
	for i := range options {
		rawOptions[i] = fn(options[i])
	}
	return rawOptions
}
