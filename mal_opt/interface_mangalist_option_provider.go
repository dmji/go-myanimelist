package mal_opt

import "net/url"

// MangaListOption are options specific to the UserService.MangaList method.
type MangaListOption interface {
	MangaListApply(v *url.Values)
}

type MangaListOptionProvider struct {
	UserListFields
	SortMangaList
	MangaStatus
}

func (s MangaListOptionProvider) NSFW(v bool) NSFW {
	return NSFW(v)
}

func (s MangaListOptionProvider) Limit(v int) Limit {
	return NewLimit(v)
}

func (s MangaListOptionProvider) Offset(v int) Offset {
	return Offset(v)
}

func (s MangaListOptionProvider) Fields(v ...string) Fields {
	return NewFields(v...)
}
