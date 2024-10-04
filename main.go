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

	scanner := Scanner{
		source: *source,
		from:   *from,
	}
	if err := scanner.Scan(); err != nil {
		panic(err)
	}
}

type Scanner struct {
	source string
	from   int
}

func (s *Scanner) Scan() error {
	file, err := os.OpenFile(s.source, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		if line <= s.from {
			line++
			continue
		}

		var pretty bytes.Buffer
		if err := json.Indent(&pretty, []byte(scanner.Text()), "", "  "); err != nil {
			return err
		}

		fmt.Println(pretty.String())
		line++
	}

	return nil
}
