package main

import (
	"testing"
)

func TestMonkeys(t *testing.T) {
	sm42 := &simplemonkey{value: 42}
	res := sm42.yell()
	if res != 42 {
		t.Fatalf("simple monkey should returns its own value : %d and not %d", 42, res)
	}

	psm42 := &proxymonkey{m: sm42}
	res = psm42.yell()
	if res != 42 {
		t.Fatalf("proxy monkey should returns its inner monkey yelling value : %d and not %d", 42, res)
	}

	opm0 := &opmonkey{a: sm42, b: psm42, f: func(a int, b int) int { return a - b }}
	res = opm0.yell()
	if res != 0 {
		t.Fatalf("op monkey should returns the result of f func with inputs from inner monkeys yelling : %d and not %d", 0, res)
	}
}
func TestPart1(t *testing.T) {
	root, _ := getData(testData())
	result := part1(root)
	expect := 152
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	_, monkeys := getData(testData())
	result := part2(monkeys)
	expect := 301
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"root: pppw + sjmn",
		"dbpl: 5",
		"cczh: sllz + lgvd",
		"zczc: 2",
		"ptdq: humn - dvpt",
		"dvpt: 3",
		"lfqf: 4",
		"humn: 5",
		"ljgn: 2",
		"sjmn: drzm * dbpl",
		"sllz: 4",
		"pppw: cczh / lfqf",
		"lgvd: ljgn * ptdq",
		"drzm: hmdt - zczc",
		"hmdt: 32",
	}
}
