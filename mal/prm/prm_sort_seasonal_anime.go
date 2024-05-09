package prm

import "net/url"

// SortSeasonalAnime is an option that allows to sort the anime results.
type SortSeasonalAnime string

const (
	// SortSeasonalByAnimeScore sorts seasonal results by anime score in
	// descending order.
	SortSeasonalByAnimeScore SortSeasonalAnime = "anime_score"
	// SortSeasonalByAnimeNumListUsers sorts seasonal results by anime num list
	// users in descending order.
	SortSeasonalByAnimeNumListUsers SortSeasonalAnime = "anime_num_list_users"
)

func (s SortSeasonalAnime) SeasonalAnimeApply(v *url.Values) { v.Set("sort", string(s)) }
