package prm

import "net/url"

// AnimeListOption are options specific to the UserService.AnimeList method.
type AnimeListOption interface {
	AnimeListApply(v *url.Values)
}

type AnimeListOptionProvider struct {
	SortAnimeList
	AnimeStatus
	UserListFields
	//AnimeFields
}

func (s AnimeListOptionProvider) NSFW(v bool) NSFW {
	return NSFW(v)
}

func (s AnimeListOptionProvider) Limit(v int) Limit {
	return NewLimit(v)
}

func (s AnimeListOptionProvider) Offset(v int) Offset {
	return Offset(v)
}

func (s AnimeListOptionProvider) Fields(v ...string) Fields {
	return NewFields(v...)
}
