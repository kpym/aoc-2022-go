package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

type packet []any

func toPacket(s string) (p packet, r string) {
	s = s[1:] // skip '['
	var new packet
	for len(s) > 0 {
		switch {
		case s[0] == ']':
			return p, s[1:]
		case s[0] == ',':
			s = s[1:]
		case s[0] == '[':
			new, s = toPacket(s)
			p = append(p, new)
		case s[0] >= '0' && s[0] <= '9':
			var i int
			for i = 0; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
			}
			n, _ := strconv.Atoi(s[:i])
			p = append(p, n)
			s = s[i:]
		}
	}
	return // never reached
}

func compare(f, s packet) (r int) {
	for i := 0; i < len(f) && i < len(s) && r == 0; i++ {
		switch f[i].(type) {
		case packet: // f is a packet
			switch s[i].(type) {
			case packet: // both are packets
				r = compare(f[i].(packet), s[i].(packet))
			default: // f is a packet, s is an int
				r = compare(f[i].(packet), packet{s[i].(int)})
			}
		default: // f is an int
			switch s[i].(type) {
			case packet: // f is an int, s is a packet
				r = compare(packet{f[i].(int)}, s[i].(packet))
			default: // both are ints
				r = f[i].(int) - s[i].(int)
			}
		}
	}
	if r != 0 {
		return r
	}
	return len(f) - len(s)
}

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(data), "\n")

	packets := []packet{}
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		p, _ := toPacket(lines[i])
		packets = append(packets, p)
	}
	packets = append(packets, packet{2}, packet{6})

	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) < 0
	})

	prod := 1
	for i, p := range packets {
		if compare(p, packet{2}) == 0 || compare(p, packet{6}) == 0 {
			prod *= i + 1
		}
	}

	println(prod)
}
