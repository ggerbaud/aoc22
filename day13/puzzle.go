package main

import (
	"advent/utils"
	"fmt"
	"golang.org/x/exp/slices"
	"sort"
	"strconv"
)

func main() {
	part1()
	part2()
}

const day = "13"
const test = false

func part1() {
	lines := utils.ReadFileLinesForDay(day, test)

	pair := make([][]interface{}, 0)
	idx := 1
	total := 0
	for _, line := range lines {
		if len(line) > 0 {
			pair = append(pair, parseLine(line))
		} else {
			if len(pair) != 2 {
				panic("no pair")
			}
			if compare(pair[0], pair[1]) {
				total += idx
			}
			pair = make([][]interface{}, 0)
			idx++
		}
	}
	if len(pair) == 2 {
		if compare(pair[0], pair[1]) {
			total += idx
		}
		pair = make([][]interface{}, 0)
	} else if len(pair) != 0 {
		panic("incomplete  pair")
	}
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
}
func part2() {
	lines := utils.ReadFileLinesForDay(day, false)

	packets := make(packets, 0)
	divider1 := parseLine("[[2]]")
	divider2 := parseLine("[[6]]")
	packets = append(packets, divider1)
	packets = append(packets, divider2)
	for _, line := range lines {
		if len(line) > 0 {
			packets = append(packets, parseLine(line))
		}
	}
	sort.Sort(packets)
	idx1 := slices.IndexFunc(packets, func(i interface{}) bool { return compareInt(i, divider1) == 0 }) + 1
	idx2 := slices.IndexFunc(packets, func(i interface{}) bool { return compareInt(i, divider2) == 0 }) + 1
	total := idx1 * idx2
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

type packets []interface{}

func (p packets) Len() int {
	return len(p)
}

func (p packets) Less(i, j int) bool {
	return compareInt(p[i], p[j]) < 0
}
func (p packets) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func compare(a, b []interface{}) bool {
	return compareList(a, b) <= 0
}

func compareInt(a, b interface{}) int {
	if n1, ok := a.(int); ok {
		if n2, ok := b.(int); ok {
			return n1 - n2
		} else if l, ok := b.([]interface{}); ok {
			return compareList([]interface{}{a}, l)
		} else {
			panic(fmt.Errorf("unknown type for %v : %t", b, b))
		}
	} else if l1, ok := a.([]interface{}); ok {
		if n2, ok := b.(int); ok {
			return compareList(l1, []interface{}{n2})
		} else if l2, ok := b.([]interface{}); ok {
			return compareList(l1, l2)
		} else {
			panic(fmt.Errorf("unknown type for %v : %t", b, b))
		}
	} else {
		panic(fmt.Errorf("unknown type for %v : %t", a, a))
	}
}

func compareList(a, b []interface{}) int {
	res := -1
	n := len(a)
	if len(b) == len(a) {
		res = 0
	} else if len(b) < len(a) {
		res = 1
		n = len(b)
	}
	for i := 0; i < n; i++ {
		cmp := compareInt(a[i], b[i])
		if cmp == 0 {
			continue
		}
		return cmp
	}
	return res
}

func parseLine(line string) []interface{} {
	l, _ := newList(line[1:])
	return l
}

func newList(data string) ([]interface{}, int) {
	out := make([]interface{}, 0)
	buf := ""
	for i := 0; i < len(data); i++ {
		if data[i] == ']' {
			if len(buf) > 0 {
				out = append(out, utils.ParseInt(buf))
				buf = ""
			}
			return out, i + 1
		} else if data[i] == '[' {
			l, n := newList(data[i+1:])
			out = append(out, l)
			i += n
		} else if data[i] == ',' {
			if len(buf) > 0 {
				out = append(out, utils.ParseInt(buf))
				buf = ""
			}
		} else {
			buf += data[i : i+1]
		}
	}
	panic("no end of list")
}
