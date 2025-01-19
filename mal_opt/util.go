package mal_opt

import (
	"strconv"
	"strings"
	"time"
)

var itoa = strconv.Itoa

func formatMALDate(d time.Time) string {
	if d.IsZero() {
		return ""
	}
	return d.Format("2006-01-02")
}

func argJoin(v ...string) string {
	if len(v) == 0 {
		return ""
	}
	return "{" + strings.Join(v, ",") + "}"
}
