package prm

import "net/url"

// Priority is an option that allows to update the priority of an anime or manga
// in the user's list with values 0=Low, 1=Medium, 2=High.
type Priority int

func (p Priority) UpdateMyAnimeListStatusApply(v *url.Values) {
	v.Set("priority", itoa(int(p)))
}
func (p Priority) UpdateMyMangaListStatusApply(v *url.Values) {
	v.Set("priority", itoa(int(p)))
}

func (p Priority) Val(v int) Priority { return Priority(v) }
