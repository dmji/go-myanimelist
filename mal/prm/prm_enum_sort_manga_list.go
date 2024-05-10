package prm

import "net/url"

// SortMangaList is an option that sorts the results when getting the user's
// manga list.
type SortMangaList string

const (
	// SortMangaListByListScore sorts results by the score of each item in the
	// list in descending order.
	SortMangaListByListScore SortMangaList = "list_score"
	// SortMangaListByListUpdatedAt sorts results by the most updated entries in
	// the list in descending order.
	SortMangaListByListUpdatedAt SortMangaList = "list_updated_at"
	// SortMangaListByMangaTitle sorts results by the manga title in ascending
	// order.
	SortMangaListByMangaTitle SortMangaList = "manga_title"
	// SortMangaListByMangaStartDate sorts results by the manga start date in
	// descending order.
	SortMangaListByMangaStartDate SortMangaList = "manga_start_date"
	// SortMangaListByMangaID sorts results by the manga ID in ascending order.
	// Note: Currently under development.
	SortMangaListByMangaID SortMangaList = "manga_id"
)

func (s SortMangaList) MangaListApply(v *url.Values) { v.Set("sort", string(s)) }

func (s SortMangaList) ByListScore() SortMangaList      { return SortMangaListByListScore }
func (s SortMangaList) ByListUpdatedAt() SortMangaList  { return SortMangaListByListUpdatedAt }
func (s SortMangaList) ByMangaTitle() SortMangaList     { return SortMangaListByMangaTitle }
func (s SortMangaList) ByMangaStartDate() SortMangaList { return SortMangaListByMangaStartDate }
func (s SortMangaList) ByMangaID() SortMangaList        { return SortMangaListByMangaID }
