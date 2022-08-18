package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/s-ichikawa/ts2d/internal"
)

var customFormat, customJavaDateFormat string

func main() {
	flag.StringVar(&customFormat, "format", "", "")
	flag.StringVar(&customFormat, "f", "", "")
	flag.StringVar(&customJavaDateFormat, "java-format", "", "")
	flag.StringVar(&customJavaDateFormat, "jf", "", "")

	flag.Parse()

	// `customFormat`, i.e. golang layout is top-class citizen.
	if customFormat != "" {
		internal.SetCustomFormat(customFormat)
	} else if customJavaDateFormat != "" {
		internal.SetCustomFormatInJavaDataPattern(customJavaDateFormat)
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
