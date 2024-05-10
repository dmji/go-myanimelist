package prm

import (
	"strconv"
	"strings"
	"time"
)

var itoa = strconv.Itoa

// MARK: Date / Time format
func formatMALDate(d time.Time) string {
	if d.IsZero() {
		return ""
	}
	return d.Format("2006-01-02")
}

func appendIf(b bool, s []string, v string) {
	if b {
		s = append(s, v)
	}
}

func argJoin(v ...string) string {
	if len(v) == 0 {
		return ""
	}
	return "{" + strings.Join(v, ",") + "}"
}
