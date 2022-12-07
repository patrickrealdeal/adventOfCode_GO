package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	root = iota
	prev
	to
	fileInfo
)

type filesystem struct {
	dirs  []string
	sizes map[string]int
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := string(data)

	fs := filesystem{sizes: make(map[string]int)}

	for _, line := range strings.Split(lines, "\n") {
		switch fs.parse(line) {
		case root:
			fs.dirs = []string{"/"}
		case prev:
			fs.dirs = fs.dirs[:len(fs.dirs)-1]
		case to:
			fs.dirs = append(fs.dirs, (strings.Split(line, " ")[2]))
		case fileInfo:
			for i := len(fs.dirs); i != 0; i-- {
				path := strings.Join(fs.dirs[:i], "/")
				if strings.HasPrefix(path, "//") {
					path = strings.TrimPrefix(path, "/")
				}
				fs.sizes[path] += computeSize(line)
			}
		default:
			continue
		}
	}

	spaceUsed := 70000000 - fs.sizes["/"]
	needed := 30000000
	dirUsages := []int{}

	for k := range fs.sizes {
		if fs.sizes[k] >= needed-spaceUsed {
			dirUsages = append(dirUsages, fs.sizes[k])
		}
	}

	sort.Ints(dirUsages)

	res := dirUsages[0]
	fmt.Println(res)
}

func (fs *filesystem) parse(s string) int {
	if ok, _ := regexp.Match("\\$ cd /$", []byte(s)); ok {
		return root
	}
	if ok, _ := regexp.Match("\\$ cd \\.\\.", []byte(s)); ok {
		return prev
	}
	if ok, _ := regexp.Match("\\$ cd .+", []byte(s)); ok {
		return to
	}
	if ok, _ := regexp.Match("[0-9]+ .+", []byte(s)); ok {
		return fileInfo
	}
	return -1
}

func computeSize(s string) int {
	ss := strings.Split(s, " ")
	v, _ := strconv.Atoi(ss[0])
	return v
}

// 272298
