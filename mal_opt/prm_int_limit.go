package mal_opt

import (
	"net/url"
	"strconv"
)

// Limit is an option that limits the results returned by a request. (<=100)
type Limit int

func (l Limit) PagingApply(v *url.Values)        { l.Apply(v) }
func (l Limit) TopicsApply(v *url.Values)        { l.Apply(v) }
func (l Limit) SeasonalAnimeApply(v *url.Values) { l.Apply(v) }
func (l Limit) AnimeListApply(v *url.Values)     { l.Apply(v) }
func (l Limit) MangaListApply(v *url.Values)     { l.Apply(v) }
func (l Limit) Apply(v *url.Values)              { v.Set("limit", strconv.Itoa(int(l))) }

func NewLimit(v int) Limit {
	if v > 100 {
		panic("limit must be <= 100")
	}
	return Limit(v)
}
