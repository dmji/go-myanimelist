package mal_opt

import "net/url"

// NumChaptersRead is an option that can update the number of chapters read of a
// manga in the user's list.
type NumChaptersRead int

func (n NumChaptersRead) UpdateMyMangaListStatusApply(v *url.Values) {
	v.Set("num_chapters_read", itoa(int(n)))
}
