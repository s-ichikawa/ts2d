package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	tsnreg,
	tsreg *regexp.Regexp

	start,
	end time.Time
)

func main() {
	tsnreg = regexp.MustCompile(`\d+\.\d+`)
	tsreg = regexp.MustCompile(`\d+`)

	r := bufio.NewReader(os.Stdin)
	defer func() {
		bf, _ := r.ReadByte()
		fmt.Println(string(bf))
	}()

	for true {
		by, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			break
		}
		fmt.Print(c(string(by)))
		if err == io.EOF {
			break
		}
	}
}

func c(in string) string {
	end = time.Now().Add(1 * time.Second)
	start = time.Unix(end.Unix()-(24*60*60*30), int64(end.Nanosecond()))

	out := tsnreg.ReplaceAllStringFunc(in, func(s string) string {
		ts := strings.Split(s, ".")
		sec, _ := strconv.ParseInt(ts[0], 10, 64)
		nsec, _ := strconv.ParseInt(ts[1], 10, 64)

		if !isTargetTerm(sec) {
			return s
		}

		return fmt.Sprintf(`"%s"`, time.Unix(sec, nsec).String())
	})

	out = tsreg.ReplaceAllStringFunc(out, func(s string) string {
		sec, _ := strconv.ParseInt(s, 10, 64)
		if !isTargetTerm(sec) {
			return s
		}
		return fmt.Sprintf(`"%s"`, time.Unix(sec, 0).String())
	})

	return out
}

func isTargetTerm(sec int64) bool {
	return start.Unix() < sec && sec < end.Unix()
}
