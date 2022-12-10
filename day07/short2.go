package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	data := bufio.NewScanner(file)

	var size = make(map[string]int)

	var scan func(string)
	scan = func(folder string) {
		for data.Scan() {
			line := data.Text()
			switch {
			case strings.HasPrefix(line, "$ cd .."):
				return
			case strings.HasPrefix(line, "$ cd"):
				subfolder := folder + "/" + line[5:]
				scan(subfolder)
				size[folder] += size[subfolder]
			case line[0] >= '0' && line[0] <= '9':
				s, _ := strconv.Atoi(strings.Fields(line)[0])
				size[folder] += s
			}
		}
	}

	scan("/")

	max := size["/"]
	missing := 30000000 - (70000000 - max)
	for _, s := range size {
		if s > missing && s < max {
			max = s
		}
	}
	println(max)
}
