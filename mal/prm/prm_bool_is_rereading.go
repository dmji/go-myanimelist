package prm

import (
	"net/url"
	"strconv"
)

// IsRereading is an option that can update if a user is rereading a manga in
// their list.
type IsRereading bool

func (r IsRereading) UpdateMyMangaListStatusApply(v *url.Values) {
	v.Set("is_rereading", strconv.FormatBool(bool(r)))
}

func (n IsRereading) True() IsRereading  { return IsRereading(true) }
func (n IsRereading) False() IsRereading { return IsRereading(false) }
