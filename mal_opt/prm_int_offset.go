package mal_opt

import (
	"net/url"
	"strconv"
)

// Offset is an option that sets the offset of the results.
type Offset int

func (o Offset) PagingApply(v *url.Values)        { o.Apply(v) }
func (o Offset) TopicsApply(v *url.Values)        { o.Apply(v) }
func (o Offset) SeasonalAnimeApply(v *url.Values) { o.Apply(v) }
func (o Offset) AnimeListApply(v *url.Values)     { o.Apply(v) }
func (o Offset) MangaListApply(v *url.Values)     { o.Apply(v) }
func (o Offset) Apply(v *url.Values)              { v.Set("offset", strconv.Itoa(int(o))) }
