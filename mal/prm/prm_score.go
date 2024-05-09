package prm

import "net/url"

// Score is an option that can update the anime and manga list scores with
// values 0-10.
type Score int

func (s Score) UpdateMyAnimeListStatusApply(v *url.Values) {
	v.Set("score", itoa(int(s)))
}
func (s Score) UpdateMyMangaListStatusApply(v *url.Values) {
	v.Set("score", itoa(int(s)))
}
