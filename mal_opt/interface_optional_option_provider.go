package mal_opt

import "net/url"

// common.OptionalParam is implemented by types that can be used as options in most methods
// such as Limit, Offset and Fields.
type OptionalParam interface {
	Apply(v *url.Values)
}

type OptionalParamProvider struct {
	AnimeFields // moved
	MangaFields // moved
}

func (s OptionalParamProvider) Limit(v int) Limit {
	return NewLimit(v)
}

func (s OptionalParamProvider) Offset(v int) Offset {
	return Offset(v)
}

func (s OptionalParamProvider) Fields(v ...string) Fields {
	return NewFields(v...)
}

func (s OptionalParamProvider) NSFW(v bool) NSFW {
	return NSFW(v)
}
