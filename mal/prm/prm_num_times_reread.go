package prm

import "net/url"

// NumTimesReread is an option that can update the number of times the user
// has reread a manga in their list.
type NumTimesReread int

func (n NumTimesReread) UpdateMyMangaListStatusApply(v *url.Values) {
	v.Set("num_times_reread", itoa(int(n)))
}
