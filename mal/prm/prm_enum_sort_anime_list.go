package prm

import "net/url"

// SortAnimeList is an option that sorts the results when getting the user's
// anime list.
type SortAnimeList string

const (
	// SortAnimeListByListScore sorts results by the score of each item in the
	// list in descending order.
	SortAnimeListByListScore SortAnimeList = "list_score"
	// SortAnimeListByListUpdatedAt sorts results by the most updated entries in
	// the list in descending order.
	SortAnimeListByListUpdatedAt SortAnimeList = "list_updated_at"
	// SortAnimeListByAnimeTitle sorts results by the anime title in ascending
	// order.
	SortAnimeListByAnimeTitle SortAnimeList = "anime_title"
	// SortAnimeListByAnimeStartDate sorts results by the anime start date in
	// descending order.
	SortAnimeListByAnimeStartDate SortAnimeList = "anime_start_date"
	// SortAnimeListByAnimeID sorts results by the anime ID in ascending order.
	// Note: Currently under development.
	SortAnimeListByAnimeID SortAnimeList = "anime_id"
)

func (s SortAnimeList) AnimeListApply(v *url.Values) { v.Set("sort", string(s)) }

func (n SortAnimeList) ByListScore() SortAnimeList      { return SortAnimeListByListScore }
func (n SortAnimeList) ByListUpdatedAt() SortAnimeList  { return SortAnimeListByListUpdatedAt }
func (n SortAnimeList) ByAnimeTitle() SortAnimeList     { return SortAnimeListByAnimeTitle }
func (n SortAnimeList) ByAnimeStartDate() SortAnimeList { return SortAnimeListByAnimeStartDate }
func (n SortAnimeList) ByAnimeID() SortAnimeList        { return SortAnimeListByAnimeID }
