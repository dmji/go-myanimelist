package mal_opt

import (
	"net/url"
	"strconv"
)

// MARK: Not Safe For Work

// NSFW is an option which sets the NSFW query option. By default this is set to
// false.
type NSFW bool

func (n NSFW) SeasonalAnimeApply(v *url.Values) { n.Apply(v) }
func (n NSFW) Apply(v *url.Values)              { v.Set("nsfw", strconv.FormatBool(bool(n))) }
