package prm

import "net/url"

// NumVolumesRead is an option that can update the number of volumes read of a
// manga in the user's list.
type NumVolumesRead int

func (n NumVolumesRead) UpdateMyMangaListStatusApply(v *url.Values) {
	v.Set("num_volumes_read", itoa(int(n)))
}
