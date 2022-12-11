package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items           []int
	op              func(int) int
	div             int
	iftrue, iffalse int
	inspected       int
}

func atoi(s string) int {
	n, _ := strconv.Atoi(strings.TrimSpace(s))
	return n
}

func newop(o, s string) func(int) int {
	if s != "old" {
		if o == "*" {
			return func(i int) int { return i * atoi(s) }
		} else {
			return func(i int) int { return i + atoi(s) }
		}
	}
	return func(i int) int { return i * i }
}

func newMonkey(lines []string) *Monkey {
	m := &Monkey{}
	for _, s := range strings.Split(lines[0][18:], ", ") {
		m.items = append(m.items, atoi(s))
	}
	m.op = newop(lines[1][23:24], lines[1][25:])
	m.div = atoi(lines[2][21:])
	m.iftrue = atoi(lines[3][29:])
	m.iffalse = atoi(lines[4][30:])
	return m
}

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(data), "\n")

	var monkeys []*Monkey
	div := 1 // we will work in modulo div
	for from := 1; from < len(lines); from += 7 {
		new := newMonkey(lines[from:])
		monkeys = append(monkeys, new)
		div *= new.div
	}

	n := 0 // to monkey
	for round := 0; round < 10000; round++ {
		for _, monkey := range monkeys {
			for _, i := range monkey.items {
				item := monkey.op(i) % div
				if item%monkey.div == 0 {
					n = monkey.iftrue
				} else {
					n = monkey.iffalse
				}
				monkeys[n].items = append(monkeys[n].items, item)
			}
			monkey.inspected += len(monkey.items)
			monkey.items = nil
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
	println(monkeys[0].inspected * monkeys[1].inspected)
}
