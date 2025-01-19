package prm

import "net/url"

// Comments is an option that allows to update the comment of anime and manga in
// the user's list.
type Comments string

func (c Comments) UpdateMyAnimeListStatusApply(v *url.Values) {
	v.Set("comments", string(c))
}

func (c Comments) UpdateMyMangaListStatusApply(v *url.Values) {
	v.Set("comments", string(c))
}
