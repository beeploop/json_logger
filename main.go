package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	source := flag.String("file", "", "file to read")
	from := flag.Int("from", 0, "line to start reading")
	flag.Parse()

	if *source == "" {
		flag.PrintDefaults()
		return
	}

	file, err := os.OpenFile(*source, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		if line <= *from {
			line++
			continue
		}

        var pretty bytes.Buffer
        if err := json.Indent(&pretty, []byte(scanner.Text()), "", "  "); err != nil {
            panic(err)
        }

        fmt.Println(pretty.String())
		line++
	}
}
