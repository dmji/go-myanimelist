package prm

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
