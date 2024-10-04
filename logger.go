package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/fsnotify/fsnotify"
)

type Logger struct {
	source string
}

func (l *Logger) Watch() error {
	file, err := os.OpenFile(l.source, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()
	if err := watcher.Add(l.source); err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	for {
		b, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}

		if err := l.prettyPrint(b); err != nil {
			return err
		}

		if err := l.waitForChange(watcher); err != nil {
			return err
		}
	}
}

func (l *Logger) Read() error {
	file, err := os.OpenFile(l.source, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		b, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}

		if err == io.EOF {
			return nil
		}

		if err := l.prettyPrint(b); err != nil {
			return err
		}
	}
}

func (l *Logger) prettyPrint(source []byte) error {
	var pretty bytes.Buffer
	if err := json.Indent(&pretty, []byte(source), "", "  "); err != nil {
		return err
	}

	os.Stdout.Write(pretty.Bytes())

	return nil
}

func (l *Logger) waitForChange(watcher *fsnotify.Watcher) error {
	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				return nil
			}
		case err := <-watcher.Errors:
			return err
		}
	}
}
