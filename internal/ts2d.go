package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	df "github.com/s-ichikawa/ts2d/internal/date_format"
)

const (
	defaultQuotation      = `"`
	defaultBufferDuration = 1 * time.Second
	defaultTerm           = 30 * 24 * time.Hour
)

var (
	tsnreg = regexp.MustCompile(`\d+\.\d+`)
	tsreg  = regexp.MustCompile(`\d+`)

	customFormat = ""
)

func SetCustomFormatInGoLayout(f string) {
	customFormat = f
}

func SetCustomFormatInJavaDataPattern(f string) {
	r := df.NewJavaSimpleDateConvertor(f)
	customFormat = r.ToLayout()
}

func TimestampToDate(in string) string {
	out := tsnreg.ReplaceAllStringFunc(in, func(s string) string {
		ts := strings.Split(s, ".")
		sec, _ := strconv.ParseInt(ts[0], 10, 64)
		nsec, _ := strconv.ParseInt(ts[1], 10, 64)

		if !isTargetTerm(sec) {
			return s
		}

		return format(time.Unix(sec, nsec))
	})

	out = tsreg.ReplaceAllStringFunc(out, func(s string) string {
		sec, _ := strconv.ParseInt(s, 10, 64)
		if !isTargetTerm(sec) {
			return s
		}
		return format(time.Unix(sec, 0))
	})

	return out
}

func isTargetTerm(sec int64) bool {
	now := time.Now()
	end := now.Add(defaultBufferDuration)
	start := now.Add(-defaultTerm)
	return start.Unix() < sec && sec < end.Unix()
}

func format(t time.Time) string {
	var d string
	if customFormat != "" {
		d = t.Format(customFormat)
	} else {
		d = t.String()
	}
	return fmt.Sprintf("%s%s%s", defaultQuotation, d, defaultQuotation)
}
