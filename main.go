package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func rmHidden(s []fs.DirEntry) []fs.DirEntry {
	src, dst := 0, 0
	for src < len(s) {
		if s[src].Name()[0] != '.' {
			s[dst] = s[src]
			dst++
		}
		src++
	}
	s = s[:dst]
	return s
}

func dir(prev string, f string, n string) {
	files, err := os.ReadDir(f) // os.ReadDir() returns *sorted* entries
	if err != nil {
		log.Fatal(err)
	}
	// Not much diff, esp since os.ReadDir() returns sorted entries
	// files = rmHidden(files)

	if n == "" {
		fmt.Println(f)
	}
	l := prev + "├──"
	for k, file := range files {
		n := file.Name()
		if n[0] == '.' {
			continue
		}
		isLast := false
		if k == len(files)-1 {
			isLast = true
		}

		if !isLast {
			fmt.Println(l, n)
		} else {
			if len(prev) > 0 {
				fmt.Println(prev[:len(prev)-1], "└──", n)
			} else {
				fmt.Println("└──", n)
			}
		}
		if file.IsDir() {
			if !isLast {
				dir(prev+"│   ", f+"/"+n, n)
			} else {
				dir(prev+"    ", f+"/"+n, n)
			}
		}
	}
}

func main() {
	f := "."
	if len(os.Args) > 1 {
		f = strings.TrimSuffix(os.Args[1], string(os.PathSeparator))
	}
	dir("", f, "")
}
