package main

import (
	"os"
)

type shape [4][4]bool

var rocks [5]shape = [5]shape{
	// ####
	shape{{true, true, true, true}, {false, false, false, false}, {false, false, false, false}, {false, false, false, false}},
	// .#.
	// ###
	// .#.
	shape{{false, true, false, false}, {true, true, true, false}, {false, true, false, false}, {false, false, false, false}},
	// ..#
	// ..#
	// ###
	shape{{true, true, true, false}, {false, false, true, false}, {false, false, true, false}, {false, false, false, false}},
	// #
	// #
	// #
	// #
	shape{{true, false, false, false}, {true, false, false, false}, {true, false, false, false}, {true, false, false, false}},
	// ##
	// ##
	shape{{true, true, false, false}, {true, true, false, false}, {false, false, false, false}, {false, false, false, false}},
}

var width [5]int = [5]int{4, 3, 3, 1, 2}
var height [5]int = [5]int{1, 3, 3, 4, 2}

type line [7]bool // one line of the cave (7 columns)

var empty line = line{}
var full line = line{true, true, true, true, true, true, true}

var cave []line
var top int // index of the top of the cave

func overlap(x, y, rock int) bool {
	for i := 0; y+i <= 0 && i < height[rock]; i++ {
		for j := 0; j < width[rock]; j++ {
			if rocks[rock][i][j] && cave[top+y+i][x+j] {
				return true
			}
		}
	}
	return false
}

func place(x, y, rock int) {
	for i := 0; i < height[rock]; i++ {
		for j := 0; j < width[rock]; j++ {
			if rocks[rock][i][j] {
				cave[top+y+i][x+j] = true
			}
		}
	}
}

func main() {
	data, _ := os.ReadFile(os.Args[1])
	// convert < and > to -1 and 1
	move := make([]int, 0, len(data))
	for _, c := range data {
		if c == '<' {
			move = append(move, -1)
		} else if c == '>' {
			move = append(move, 1)
		}
	}

	cave = make([]line, 0, 4096)
	top = 0 // index of the top of the cave
	cave = append(cave, full, empty, empty, empty, empty, empty, empty, empty)

	count := 0   // number of rocks placed
	rock := 0    // 0..4
	x, y := 2, 4 // y is relative to the top of the cave
	for i := 0; ; i = (i + 1) % len(move) {
		m := move[i]
		if x+m >= 0 && x+width[rock]+m <= 7 && (y > 0 || !overlap(x+m, y, rock)) {
			x += m
		}
		if y > 1 || !overlap(x, y-1, rock) {
			y--
			continue
		}
		// place the rock
		place(x, y, rock)
		if y+height[rock]-1 > 0 {
			top += y + height[rock] - 1
		}
		for len(cave) < top+7 {
			cave = append(cave, empty)
		}
		// next rock
		rock = (rock + 1) % len(rocks)
		x, y = 2, 4
		count++
		if count == 2022 {
			break
		}
	}
	println(top)
}
