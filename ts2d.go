package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/s-ichikawa/ts2d/internal"
)

var customFormat string

func main() {
	flag.StringVar(&customFormat, "format", "", "")
	flag.StringVar(&customFormat, "f", "", "")

	flag.Parse()

	if customFormat != "" {
		internal.SetCustomFormat(customFormat)
	}

	r := bufio.NewReader(os.Stdin)
	for true {
		by, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			break
		}
		fmt.Print(internal.TimestampToDate(string(by)))
		if err == io.EOF {
			break
		}
	}
}
