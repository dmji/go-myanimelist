package mal_opt

import "net/url"

type OptionFunc func(v *url.Values)

func (f OptionFunc) Apply(v *url.Values) { f(v) }
