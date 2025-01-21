package mal_prm

import "strings"

func argJoin(v ...string) string {
	if len(v) == 0 {
		return ""
	}
	return "{" + strings.Join(v, ",") + "}"
}
