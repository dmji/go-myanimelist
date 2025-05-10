package mal_opt

import (
	"net/url"
	"strings"
)

// Fields is an option that allows to choose the fields that should be returned
// as by default, the API doesn't return all fields.
//
// Example:
//
//	Fields{"synopsis", "my_list_status{priority,comments}"}
type Fields []string

func (f Fields) SeasonalAnimeApply(v *url.Values) { f.Apply(v) }
func (f Fields) Apply(v *url.Values) {
	if len(f) != 0 {
		v.Set("fields", strings.Join(f, ","))
	}
}

func NewFields(v ...string) Fields {
	res := make([]string, 0, len(v))
	res = append(res, v...)
	return Fields(res)
}
