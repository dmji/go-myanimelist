package prm

import (
	"net/url"
	"time"

	"github.com/dmji/go-myanimelist/mal/common"
)

// FinishDate is an option that allows to update the finish date of anime and manga
// in the user's list.
type FinishDate time.Time

func (d FinishDate) UpdateMyAnimeListStatusApply(v *url.Values) {
	v.Set("finish_date", common.FormatMALDate(time.Time(d)))
}
func (d FinishDate) UpdateMyMangaListStatusApply(v *url.Values) {
	v.Set("finish_date", common.FormatMALDate(time.Time(d)))
}

func (d FinishDate) Val(v time.Time) FinishDate { return FinishDate(v) }
