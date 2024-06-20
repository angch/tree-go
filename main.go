package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func dir(prev string, f string, n string) {
	files, err := os.ReadDir(f)
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	if n == "" {
		fmt.Println(f)
	} else {
		fmt.Println(prev[:len(prev)-1], "├──", n)
	}
	l := prev + "├──"
	for _, file := range files {
		n := file.Name()
		if n[0] == '.' {
			continue
		}
		if file.IsDir() {
			dir(prev+"│  ", f+"/"+n, n)
		} else {
			fmt.Println(l, n)
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
