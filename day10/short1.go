package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	stop, cycle, score, x := 20, 0, 0, 1

	next := func() {
		if cycle++; cycle == stop {
			score += stop * x
			stop += 40
		}
	}

	for _, line := range lines {
		next()              // addx or noop
		if line[0] == 'a' { // if addx
			next()
			a, _ := strconv.Atoi(line[5:])
			x += a
		}
	}

	println(score)
}
