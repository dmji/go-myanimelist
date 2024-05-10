package prm

import (
	"net/url"
)

// NumTimesRewatched is an option that can update the number of times the user
// has rewatched an anime in their list.
type NumTimesRewatched int

func (n NumTimesRewatched) UpdateMyAnimeListStatusApply(v *url.Values) {
	v.Set("num_times_rewatched", itoa(int(n)))
}

func (n NumTimesRewatched) Val(v int) NumTimesRewatched { return NumTimesRewatched(v) }
