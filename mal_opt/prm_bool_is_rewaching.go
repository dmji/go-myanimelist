package mal_opt

import (
	"net/url"
	"strconv"
)

// IsRewatching is an option that can update if a user is rewatching an anime in
// their list.
type IsRewatching bool

func (r IsRewatching) UpdateMyAnimeListStatusApply(v *url.Values) {
	v.Set("is_rewatching", strconv.FormatBool(bool(r)))
}
