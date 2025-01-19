package mal_opt

import (
	"net/url"
	"time"
)

// UpdateMyAnimeListStatusOption are options specific to the
// Service.UpdateMyListStatus method.
type UpdateMyAnimeListStatusOption interface {
	UpdateMyAnimeListStatusApply(fnSet *url.Values)
}

type UpdateMyAnimeListStatusOptionProvider struct {
	AnimeStatus
	RewatchValue
}

func (s UpdateMyAnimeListStatusOptionProvider) Tags(v ...string) Tags {
	return NewTags(v...)
}

func (s UpdateMyAnimeListStatusOptionProvider) Comments(v string) Comments {
	return Comments(v)
}

func (s UpdateMyAnimeListStatusOptionProvider) Priority(v int) Priority {
	return Priority(v)
}

func (s UpdateMyAnimeListStatusOptionProvider) NumTimesRewatched(v int) NumTimesRewatched {
	return NumTimesRewatched(v)
}

func (s UpdateMyAnimeListStatusOptionProvider) NumEpisodesWatched(v int) NumEpisodesWatched {
	return NumEpisodesWatched(v)
}

func (s UpdateMyAnimeListStatusOptionProvider) Score(v int) Score {
	return Score(v)
}

func (d UpdateMyAnimeListStatusOptionProvider) StartDate(v time.Time) StartDate {
	return StartDate(v)
}

func (d UpdateMyAnimeListStatusOptionProvider) FinishDate(v time.Time) FinishDate {
	return FinishDate(v)
}

func (s UpdateMyAnimeListStatusOptionProvider) IsRewatching(v bool) IsRewatching {
	return IsRewatching(v)
}
