package common

import (
	"net/url"
	"strconv"
	"strings"
)

// common.OptionalParam is implemented by types that can be used as options in most methods
// such as Limit, Offset and Fields.
type OptionalParam interface {
	Apply(v *url.Values)
}

// MARK: Parameters for endpoints that return a list
// Limit is an option that limits the results returned by a request.
type Limit int

func (l Limit) PagingApply(v *url.Values)        { l.Apply(v) }
func (l Limit) TopicsApply(v *url.Values)        { l.Apply(v) }
func (l Limit) SeasonalAnimeApply(v *url.Values) { l.Apply(v) }
func (l Limit) AnimeListApply(v *url.Values)     { l.Apply(v) }
func (l Limit) MangaListApply(v *url.Values)     { l.Apply(v) }
func (l Limit) Apply(v *url.Values)              { v.Set("limit", strconv.Itoa(int(l))) }

// Offset is an option that sets the offset of the results.
type Offset int

func (o Offset) PagingApply(v *url.Values)        { o.Apply(v) }
func (o Offset) TopicsApply(v *url.Values)        { o.Apply(v) }
func (o Offset) SeasonalAnimeApply(v *url.Values) { o.Apply(v) }
func (o Offset) AnimeListApply(v *url.Values)     { o.Apply(v) }
func (o Offset) MangaListApply(v *url.Values)     { o.Apply(v) }
func (o Offset) Apply(v *url.Values)              { v.Set("offset", strconv.Itoa(int(o))) }

// MARK: Choosing fields
// Fields is an option that allows to choose the fields that should be returned
// as by default, the API doesn't return all fields.
//
// Example:
//
//	Fields{"synopsis", "my_list_status{priority,comments}"}
type Fields []string

func (f Fields) SeasonalAnimeApply(v *url.Values) { f.Apply(v) }
func (f Fields) AnimeListApply(v *url.Values)     { f.Apply(v) }
func (f Fields) MangaListApply(v *url.Values)     { f.Apply(v) }
func (f Fields) DetailsApply(v *url.Values)       { f.Apply(v) }
func (f Fields) MyInfoApply(v *url.Values)        { f.Apply(v) }
func (f Fields) Apply(v *url.Values) {
	if len(f) != 0 {
		v.Set("fields", strings.Join(f, ","))
	}
}

// MARK: Not Safe For Work
// NSFW is an option which sets the NSFW query option. By default this is set to
// false.
type NSFW bool

func (n NSFW) SeasonalAnimeApply(v *url.Values) { n.Apply(v) }
func (n NSFW) AnimeListApply(v *url.Values)     { n.Apply(v) }
func (n NSFW) MangaListApply(v *url.Values)     { n.Apply(v) }
func (n NSFW) Apply(v *url.Values)              { v.Set("nsfw", strconv.FormatBool(bool(n))) }

type OptionFunc func(v *url.Values)

func (f OptionFunc) Apply(v *url.Values) {
	f(v)
}

func OptionFromQuery(query string) OptionFunc {
	return OptionFunc(func(v *url.Values) {
		v.Set("q", query)
	})
}
