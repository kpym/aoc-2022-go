package main

import (
	"bytes"
	"os"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type cube [3]int

func nextto(c, d cube) bool {
	count := 0
	for i := 0; i < 3; i++ {
		count += abs(c[i] - d[i])
	}
	return count == 1
}

func neighbours(c cube, cs []cube) int {
	count := 0
	for i := range cs {
		if nextto(c, cs[i]) {
			count++
		}
	}
	return count
}

func main() {
	data, _ := os.ReadFile(os.Args[1])

	var (
		cubes []cube
		c     cube
		i     int
	)
	for _, v := range bytes.FieldsFunc(data, func(r rune) bool { return r < '0' || r > '9' }) {
		c[i], _ = strconv.Atoi(string(v))
		i++
		if i == 3 {
			cubes = append(cubes, c)
			i = 0
		}
	}

	area := 0
	for i := 0; i < len(cubes); i++ {
		area += 6 - 2*neighbours(cubes[i], cubes[:i])
	}

	println(area)
}
