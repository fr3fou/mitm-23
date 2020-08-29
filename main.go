package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := ""
	flag.StringVar(&dir, "dir", "", "path to the directory of subs")

	flag.Parse()

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// We only need the files
		if info.IsDir() {
			return nil
		}

		fmt.Println(path)
		return nil
	})

	if err != nil {
		panic(err)
	}
}
