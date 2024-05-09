package prm

import "net/url"

// RereadValue is an option that can update the reread value of a manga in the
type RereadValue int

const (
	RereadNoValue  RereadValue = 0
	RereadVeryLow  RereadValue = 1
	RereadLow      RereadValue = 2
	RereadMedium   RereadValue = 3
	RereadHigh     RereadValue = 4
	RereadVeryHigh RereadValue = 5
)

func (r RereadValue) UpdateMyMangaListStatusApply(v *url.Values) {
	v.Set("reread_value", itoa(int(r)))
}
