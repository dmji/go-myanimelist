package mal_opt

import (
	"net/url"
	"time"
)

// StartDate is an option that allows to update the start date of anime and manga
// in the user's list.
type StartDate time.Time

func (d StartDate) UpdateMyAnimeListStatusApply(v *url.Values) {
	v.Set("start_date", formatMALDate(time.Time(d)))
}

func (d StartDate) UpdateMyMangaListStatusApply(v *url.Values) {
	v.Set("start_date", formatMALDate(time.Time(d)))
}

// FinishDate is an option that allows to update the finish date of anime and manga
// in the user's list.
type FinishDate time.Time

func (d FinishDate) UpdateMyAnimeListStatusApply(v *url.Values) {
	v.Set("finish_date", formatMALDate(time.Time(d)))
}

func (d FinishDate) UpdateMyMangaListStatusApply(v *url.Values) {
	v.Set("finish_date", formatMALDate(time.Time(d)))
}
