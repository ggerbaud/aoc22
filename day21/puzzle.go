package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "21"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	original, monkeys := getData(lines)
	total := part1(original)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(monkeys)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(root monkey) int {
	total := root.yell()
	return total
}
func part2(monkeys map[string]monkey) int {
	root := monkeys["root"].(*opmonkey)
	me := monkeys["humn"].(*proxymonkey)
	myself := &human{me.m.yell()}
	me.m = myself
	vA := root.a.yell()
	vB := root.b.yell()
	root.a.constraint(vB)
	root.b.constraint(vA)
	total := myself.value
	return total
}

type monkey interface {
	yell() int
	constraint(v int)
}

type simplemonkey struct {
	value int
}

func (m *simplemonkey) yell() int {
	return m.value
}

func (m *simplemonkey) constraint(v int) {
}

type human struct {
	value int
}

func (m *human) yell() int {
	return m.value
}

func (m *human) constraint(v int) {
	m.value = v
}

type opmonkey struct {
	a, b        monkey
	f, rfa, rfb func(int, int) int
}

func (m *opmonkey) yell() int {
	return m.f(m.a.yell(), m.b.yell())
}

func (m *opmonkey) constraint(v int) {
	vA := m.a.yell()
	vB := m.b.yell()
	if m.f(vA, vB) != v {
		nvB := m.rfb(v, vA)
		m.b.constraint(nvB)
		nnvb := m.b.yell()
		if nnvb == vB {
			nvA := m.rfa(v, vB)
			m.a.constraint(nvA)
		}
	}
}

type proxymonkey struct {
	m monkey
}

func (m *proxymonkey) yell() int {
	return m.m.yell()
}
func (m *proxymonkey) constraint(v int) {
	m.m.constraint(v)
}

type node struct {
	val, size  int
	prev, next *node
}

func getData(lines []string) (monkey, map[string]monkey) {
	sum := func(a, b int) int { return a + b }
	rsum := func(r, x int) int { return r - x }
	min := func(a, b int) int { return a - b }
	rmina := func(r, b int) int { return r + b }
	rminb := func(r, a int) int { return a - r }
	mul := func(a, b int) int { return a * b }
	rmul := func(r, x int) int { return r / x }
	div := func(a, b int) int { return a / b }
	rdiva := func(r, b int) int { return r * b }
	rdivb := func(r, a int) int {
		if r == 0 {
			return a - 1
		}
		return a / r
	}
	monkeys := make(map[string]monkey)
	for _, line := range lines {
		data := strings.Split(line, ":")
		name := strings.TrimSpace(data[0])
		currM := monkeys[name]
		rest := strings.TrimSpace(data[1])
		var newM monkey
		value, err := strconv.Atoi(rest)
		if err == nil {
			newM = &simplemonkey{value}
		} else {
			var op, ropa, ropb func(int, int) int
			isSum := strings.Split(rest, "+")
			isMin := strings.Split(rest, "-")
			isMul := strings.Split(rest, "*")
			isDiv := strings.Split(rest, "/")
			var data []string
			if len(isSum) == 2 {
				op = sum
				ropa = rsum
				ropb = rsum
				data = isSum
			} else if len(isMin) == 2 {
				op = min
				ropa = rmina
				ropb = rminb
				data = isMin
			} else if len(isMul) == 2 {
				op = mul
				ropa = rmul
				ropb = rmul
				data = isMul
			} else if len(isDiv) == 2 {
				op = div
				ropa = rdiva
				ropb = rdivb
				data = isDiv
			} else {
				panic("wtf")
			}
			newM = &opmonkey{
				f:   op,
				rfa: ropa,
				rfb: ropb,
				a:   getOrCreate(strings.TrimSpace(data[0]), monkeys),
				b:   getOrCreate(strings.TrimSpace(data[1]), monkeys),
			}
		}

		if currM != nil {
			if p, ok := currM.(*proxymonkey); !ok {
				panic("wtf")
			} else {
				p.m = newM
				newM = p
			}
		} else if name == "humn" {
			newM = &proxymonkey{m: newM}
		}
		monkeys[name] = newM
	}
	return monkeys["root"], monkeys
}

func getOrCreate(name string, monkeys map[string]monkey) monkey {
	if _, ok := monkeys[name]; !ok {
		monkeys[name] = &proxymonkey{}
	}
	return monkeys[name]
}
