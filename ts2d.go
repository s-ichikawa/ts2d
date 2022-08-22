package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/s-ichikawa/ts2d/internal"
)

var customFormat, customGoDateFormat string

func main() {
	flag.StringVar(&customFormat, "format", "", "")
	flag.StringVar(&customFormat, "f", "", "")
	flag.StringVar(&customGoDateFormat, "go-format", "", "")
	flag.StringVar(&customGoDateFormat, "gf", "", "")

	flag.Parse()

	// `customFormat`, i.e. Java date template is the top-class citizen.
	if customFormat != "" {
		internal.SetCustomFormatInJavaDataPattern(customFormat)
	} else if customGoDateFormat != "" {
		internal.SetCustomFormatInGoLayout(customGoDateFormat)
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
