package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/asticode/go-astisub"
)

var whitelist = []string{
	"23",
	"Twenty Three",
	"Twenty-Three",
	"twenty three",
	"twenty-three",
	"two three",
	"two-three",
}

var blacklist = []string{
	"Synchro to version",
}

func main() {
	dir := ""
	flag.StringVar(&dir, "dir", "", "path to the directory of subs")

	flag.Parse()

	subs := []astisub.Subtitles{}
	start := time.Now()
	log.Printf("Beginning parsing SRT files in %s...", dir)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// We only need the files
		if info.IsDir() {
			return nil
		}

		// Parse the file
		srt, err := astisub.OpenFile(path)
		if err != nil {
			return err
		}

		subs = append(subs, *srt)

		return nil
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Parsed SRT files. Time elapsed: %s.", time.Since(start))

	start = time.Now()
	log.Printf("Beginning to count occurrences of 23...")
	counter := 0
	for i, sub := range subs {
		for _, item := range sub.Items {
			str := item.String()
			if containsArr(str, whitelist) && !containsArr(str, blacklist) {
				fmt.Println(printLine(*item))
				fmt.Printf("found in Episode %d\n", i+1)
				fmt.Println()
				counter++
			}
		}
	}

	log.Printf("Finished counting. Time elapsed: %s", time.Since(start))
	fmt.Printf("Occurrences found: %d\n", counter)
}

func containsArr(str string, arr []string) bool {
	for _, s := range arr {
		if strings.Contains(str, s) {
			return true
		}
	}

	return false
}

func printLine(item astisub.Item) string {
	var os []string

	for _, l := range item.Lines {
		os = append(os, l.String())
	}

	return strings.Join(os, "\n")
}
