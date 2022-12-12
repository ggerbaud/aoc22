package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

const day = "11"

func part1() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	monkeys := getMonkeys(fileScanner)
	total := makeBusiness(monkeys, 20, 3)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
}
func part2() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	monkeys := getMonkeys(fileScanner)
	total := makeBusiness(monkeys, 10000, 1)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func getMonkeys(fileScanner *bufio.Scanner) []*monkey {
	var monkeys []*monkey
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(line, "Monkey") {
			monkeys = append(monkeys, parseMonkey(fileScanner))
		}
	}
	return monkeys
}

func parseMonkey(fileScanner *bufio.Scanner) *monkey {
	fileScanner.Scan()
	line := fileScanner.Text()
	startItems := strings.Split(line, ": ")
	items := strings.Split(startItems[1], ",")
	fileScanner.Scan()
	ops := strings.Split(fileScanner.Text(), " ")
	fileScanner.Scan()
	divTest := strings.Split(fileScanner.Text(), " ")
	fileScanner.Scan()
	trueTarget := strings.Split(fileScanner.Text(), " ")
	fileScanner.Scan()
	falseTarget := strings.Split(fileScanner.Text(), " ")
	var inspect func(int) int
	k := len(ops)
	if ops[k-2] == "+" {
		inspect = addr(inter(ops[k-1]))
	} else if ops[k-2] == "*" {
		if ops[k-1] == "old" {
			inspect = selfmult()
		} else {
			inspect = multr(inter(ops[k-1]))
		}
	} else {
		panic(fmt.Errorf("wierd ops %v", ops))
	}
	kd, ktt, kft := len(divTest), len(trueTarget), len(falseTarget)
	return &monkey{itemize(items), 0, inspect, inter(divTest[kd-1]), inter(trueTarget[ktt-1]), inter(falseTarget[kft-1])}
}

func makeBusiness(monkeys []*monkey, rounds int, worryDivider int) int {
	bigMod := 1
	for _, m := range monkeys {
		bigMod *= m.divider
	}
	for round := 0; round < rounds; round++ {
		for _, m := range monkeys {
			m.inspections += len(m.items)
			for _, i := range m.items {
				m.items = m.items[1:]
				worry := m.inspect(i, worryDivider, bigMod)
				dest := m.falseTarget
				if m.test(worry) {
					dest = m.trueTarget
				}
				monkeys[dest].items = append(monkeys[dest].items, worry)
			}
		}
	}
	b1, b2 := 0, 0
	for _, m := range monkeys {
		if m.inspections > b1 {
			b2 = b1
			b1 = m.inspections
		} else if m.inspections > b2 {
			b2 = m.inspections
		}
	}
	return b1 * b2
}

func addr(n int) func(int) int {
	return func(i int) int {
		return i + n
	}
}

func multr(n int) func(int) int {
	return func(i int) int {
		return i * n
	}
}

func selfmult() func(int) int {
	return func(i int) int {
		return i * i
	}
}

func itemize(items []string) []int {
	var result []int
	for _, item := range items {
		result = append(result, inter(strings.TrimSpace(item)))
	}
	return result
}

type monkey struct {
	items       []int
	inspections int
	inspecter   func(int) int
	divider     int
	trueTarget  int
	falseTarget int
}

func (m *monkey) test(w int) bool {
	return w%m.divider == 0
}

func (m *monkey) inspect(w, divider, mod int) int {
	return (m.inspecter(w) / divider) % mod
}

func inter(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
