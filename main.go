package main

import (
	"flag"
)

func main() {
	source := flag.String("file", "", "file to read")
	watch := flag.Bool("watch", false, "watch file for changes")
	flag.Parse()

	if *source == "" {
		flag.PrintDefaults()
		return
	}

	logger := Logger{source: *source}
	if *watch {
		if err := logger.Watch(); err != nil {
			panic(err)
		}
	} else {
		if err := logger.Read(); err != nil {
			panic(err)
		}
	}
}
