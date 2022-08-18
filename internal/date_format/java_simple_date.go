package date_format

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	singleQuotedEscape = "${SQ}"
	tokenTmpl          = "${T%d}"
)

type javaSimpleDateConvertor struct {
	format string
}

func NewJavaSimpleDateConvertor(format string) javaSimpleDateConvertor {
	return javaSimpleDateConvertor{
		format: format,
	}
}

func (c javaSimpleDateConvertor) ToLayout() string {
	return c.parseFmt(c.format)
}

func (c javaSimpleDateConvertor) parseFmt(fmt string) string {
	// create tokens
	l, tokens := c.generateTokens(fmt)

	// hour
	l = strings.Replace(l, "HH", "15", -1)
	l = strings.Replace(l, "H", "15", -1)
	l = strings.Replace(l, "hh", "03", -1)
	l = strings.Replace(l, "h", "3", -1)
	l = strings.Replace(l, "a", "PM", -1)
	// minute
	l = strings.Replace(l, "mm", "04", -1)
	// second
	l = strings.Replace(l, "ss", "05", -1)
	// millisecond
	l = strings.Replace(l, "SSS", "000", -1)

	// year
	l = strings.Replace(l, "yyyy", "2006", -1)
	l = strings.Replace(l, "yy", "06", -1)
	// month
	l = strings.Replace(l, "MMM", "Jan", -1)
	l = strings.Replace(l, "MM", "01", -1)
	// day
	l = strings.Replace(l, "dd", "02", -1)
	l = strings.Replace(l, "d", "_2", -1)

	// weekday
	l = strings.Replace(l, "EEEE", "Monday", -1)
	l = strings.Replace(l, "EEE", "Mon", -1)

	// timezone
	l = strings.Replace(l, "XXX", "-07:00", -1)
	l = strings.Replace(l, "Z", "-0700", -1)
	l = strings.Replace(l, "z", "MST", -1)

	// recover tokens
	l = c.recoverTokens(l, tokens)

	return l
}

func (c javaSimpleDateConvertor) generateTokens(f string) (string, []string) {
	// create temporary tokens for single-quoted escape char
	f = strings.ReplaceAll(f, "''", singleQuotedEscape)

	// create temporary tokens for single-quoted phrases
	r, _ := regexp.Compile("('.*?')")
	ts := r.FindAllString(f, -1)
	tokens := make([]string, len(ts))
	for i, t := range ts {
		f = strings.Replace(f, t, fmt.Sprintf(tokenTmpl, i), -1)
		// remove leading & trailing single quote from each token
		tokens[i] = t[1 : len(t)-1]
	}

	return f, tokens
}

func (c javaSimpleDateConvertor) recoverTokens(f string, tokens []string) string {
	if len(tokens) == 0 {
		return f
	}
	for i, t := range tokens {
		f = strings.Replace(f, fmt.Sprintf(tokenTmpl, i), t, -1)
	}

	// recover escaped chars
	f = strings.ReplaceAll(f, singleQuotedEscape, "'")
	return f
}
