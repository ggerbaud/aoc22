package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "20"

func main() {
	testData := getData(true)
	part1Test := part1(testData)
	part1Expect := 3
	if part1Test != part1Expect {
		panic(fmt.Errorf("expected %d but was %d for part1 test", part1Expect, part1Test))
	}
	original := getData(false)
	total := part1(original)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	part2Test := part2(testData)
	part2Expect := 1623178306
	if part2Test != part2Expect {
		panic(fmt.Errorf("expected %d but was %d for part2 test", part2Expect, part2Test))
	}
	total = part2(original)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(original []int) int {

	nodes, zero := buildNodes(original, 1)

	for _, n := range nodes {
		n.move()
	}

	total := zero.get(1000) + zero.get(2000) + zero.get(3000)
	return total
}
func part2(original []int) int {

	nodes, zero := buildNodes(original, 811589153)

	for i := 0; i < 10; i++ {
		for _, n := range nodes {
			n.move()
		}
	}

	total := zero.get(1000) + zero.get(2000) + zero.get(3000)
	return total
}

func buildNodes(data []int, key int) ([]*node, *node) {
	n := len(data)
	first := &node{val: data[0] * key, size: n}
	zero := first
	nodes := []*node{first}
	for _, v := range data[1:] {
		prev := nodes[len(nodes)-1]
		nn := &node{val: v * key, size: n, prev: prev}
		prev.next = nn
		nodes = append(nodes, nn)
		if v == 0 {
			zero = nn
		}
	}
	last := nodes[len(nodes)-1]
	last.next = first
	first.prev = last
	return nodes, zero
}

type node struct {
	val, size  int
	prev, next *node
}

func (n *node) move() {
	// detach
	n.prev.next = n.next
	n.next.prev = n.prev

	if n.val > 0 {
		for i := 0; i < (n.val % (n.size - 1)); i++ {
			n.next = n.next.next
			n.prev = n.prev.next
		}
	} else if n.val < 0 {
		for i := 0; i < (-n.val % (n.size - 1)); i++ {
			n.next = n.prev
			n.prev = n.prev.prev
		}
	}

	// attach
	n.prev.next = n
	n.next.prev = n
}

func (n *node) get(k int) int {
	nd := n
	for i := 0; i < (k % n.size); i++ {
		nd = nd.next
	}
	return nd.val
}

func getData(test bool) []int {
	lines := utils.ReadFileLinesForDay(day, test)
	var out []int
	for _, line := range lines {
		out = append(out, utils.ParseInt(strings.TrimSpace(line)))
	}
	return out
}
