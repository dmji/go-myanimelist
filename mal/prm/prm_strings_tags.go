package prm

import (
	"net/url"
	"strings"
)

// Tags is an option that allows to update the comma-separated tags of anime and
// manga in the user's list.
type Tags []string

func (t Tags) UpdateMyAnimeListStatusApply(v *url.Values) {
	v.Set("tags", strings.Join(t, ","))
}
func (t Tags) UpdateMyMangaListStatusApply(v *url.Values) {
	v.Set("tags", strings.Join(t, ","))
}

func (f Tags) f(v ...string) Tags {
	res := make([]string, 0, len(v))
	res = append(res, v...)
	return Tags(res)
}
