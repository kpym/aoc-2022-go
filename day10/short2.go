package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var crt [6][40]byte
	pixel, row, x := 0, 0, 1

	next := func() {
		if pixel >= x-1 && pixel <= x+1 {
			crt[row][pixel] = '#'
		} else {
			crt[row][pixel] = '.'
		}
		pixel++
		if pixel == 40 {
			pixel = 0
			row++
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

	for _, row := range crt {
		for i := 0; i < len(row); i += 5 {
			print(string(row[i : i+5]))
			print("|")
		}
		println()
	}
}
