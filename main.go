package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/asticode/go-astisub"
)

var twentyThree = []string{
	"23",
	"Twenty Three",
	"Twenty-Three",
}

func main() {
	dir := ""
	flag.StringVar(&dir, "dir", "", "path to the directory of subs")

	flag.Parse()

	lines := []string{}
	start := time.Now()
	log.Printf("Beginning parsing SRT files in %s", dir)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// We only need the files
		if info.IsDir() {
			return nil
		}

		// Parse the file
		srt, err := astisub.OpenFile(path)
		if err != nil {
			return err
		}

		for _, item := range srt.Items {
			lines = append(lines, item.String())
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
	log.Printf("Parsed SRT files. Time elapsed: %s.", time.Since(start))
}
