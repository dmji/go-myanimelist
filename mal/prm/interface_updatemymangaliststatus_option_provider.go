package prm

import (
	"net/url"
	"time"
)

type UpdateMyMangaListStatusOption interface {
	UpdateMyMangaListStatusApply(v *url.Values)
}

type UpdateMyMangaListStatusOptionProvider struct {
	MangaStatus
	RereadValue
}

func (s UpdateMyMangaListStatusOptionProvider) Priority(v int) Priority {
	return Priority(v)
}

func (s UpdateMyMangaListStatusOptionProvider) Tags(v ...string) Tags {
	return NewTags(v...)
}

func (s UpdateMyMangaListStatusOptionProvider) Comments(v string) Comments {
	return Comments(v)
}

func (s UpdateMyMangaListStatusOptionProvider) IsRereading(v bool) IsRereading {
	return IsRereading(v)
}

func (s UpdateMyMangaListStatusOptionProvider) NumChaptersRead(v int) NumChaptersRead {
	return NumChaptersRead(v)
}

func (s UpdateMyMangaListStatusOptionProvider) NumTimesReread(v int) NumTimesReread {
	return NumTimesReread(v)
}

func (s UpdateMyMangaListStatusOptionProvider) NumVolumesRead(v int) NumVolumesRead {
	return NumVolumesRead(v)
}

func (s UpdateMyMangaListStatusOptionProvider) Score(v int) Score {
	return Score(v)
}

func (d UpdateMyMangaListStatusOptionProvider) StartDate(v time.Time) StartDate {
	return StartDate(v)
}

func (d UpdateMyMangaListStatusOptionProvider) FinishDate(v time.Time) FinishDate {
	return FinishDate(v)
}
