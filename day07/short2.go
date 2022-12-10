package main

import (
	"os"
	"strconv"
	"strings"
)

var (
	lines   []string
	folders = make(map[string]int)
)

func scan(folder string, from int) (next int) {
	for i := from; i < len(lines); i++ {
		line := lines[i]
		switch {
		case line[:7] == "$ cd ..": //return at line i
			return i
		case line[:4] == "$ cd": //read subfolder
			subfolder := folder + "/" + line[5:]
			i = scan(subfolder, i+2)
			folders[folder] += folders[subfolder]
		case line[0] >= '0' && line[0] <= '9': // add file size
			s, _ := strconv.Atoi(line[:strings.IndexByte(line, ' ')])
			folders[folder] += s
		}
	}
	return len(lines)
}

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines = strings.Split(strings.TrimSpace(string(data)), "\n")

	scan("/", 2)

	size := folders["/"]
	missing := 30000000 - (70000000 - size)
	for _, s := range folders {
		if s > missing && s < size {
			size = s
		}
	}
	println(size)
}
