package prm

import "net/url"

// SeasonalAnimeOption are options specific to the Service.Seasonal method.
// +Fields
type SeasonalAnimeOption interface {
	SeasonalAnimeApply(v *url.Values)
}

type SeasonalAnimeOptionProvider struct {
	AnimeSeason

	AnimeFields
	MangaFields
	SortSeasonalAnime
}

func (s SeasonalAnimeOptionProvider) Limit(v int) Limit {
	return NewLimit(v)
}

func (s SeasonalAnimeOptionProvider) Offset(v int) Offset {
	return Offset(v)
}

func (s SeasonalAnimeOptionProvider) Fields(v ...string) Fields {
	return NewFields(v...)
}

func (s SeasonalAnimeOptionProvider) NSFW(v bool) NSFW {
	return NSFW(v)
}
