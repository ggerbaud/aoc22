package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

const day = "12"

func part1() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	first, target := parseMap(fileScanner)
	ok, total := solve(first, target)
	if !ok {
		panic("wtf")
	}
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

	lowest, target := parseMap2(fileScanner)
	ok, total := solve2(lowest, target)
	if !ok {
		panic("wtf")
	}
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func parseMap(fileScanner *bufio.Scanner) (first, target *square) {
	var out [][]*square
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var ol []*square
		for _, c := range line {
			sqr := &square{height: c}
			if c == 'S' {
				first = sqr
				sqr.height = 'a'
			} else if c == 'E' {
				target = sqr
				sqr.height = 'z'
			}
			ol = append(ol, sqr)
		}
		out = append(out, ol)
	}
	calcVisitables(out)
	return
}

func parseMap2(fileScanner *bufio.Scanner) ([]*square, *square) {
	var out [][]*square
	var target *square
	var lowest []*square
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var ol []*square
		for _, c := range line {
			sqr := &square{height: c}
			if c == 'S' || c == 'a' {
				lowest = append(lowest, sqr)
				sqr.height = 'a'
			} else if c == 'E' {
				target = sqr
				sqr.height = 'z'
			}
			ol = append(ol, sqr)
		}
		out = append(out, ol)
	}
	calcVisitables(out)
	return lowest, target
}

func calcVisitables(dataMap [][]*square) {
	for l := 0; l < len(dataMap); l++ {
		for c := 0; c < len(dataMap[l]); c++ {
			addVisitable(dataMap, dataMap[l][c], l-1, c)
			addVisitable(dataMap, dataMap[l][c], l+1, c)
			addVisitable(dataMap, dataMap[l][c], l, c-1)
			addVisitable(dataMap, dataMap[l][c], l, c+1)
		}
	}
}

func addVisitable(dataMap [][]*square, sqr *square, i, j int) {
	if i >= 0 && i < len(dataMap) && j >= 0 && j < len(dataMap[i]) && dataMap[i][j].height-sqr.height <= 1 {
		sqr.visitables = append(sqr.visitables, dataMap[i][j])
	}
}

type square struct {
	height     rune
	visitables []*square
}

type step struct {
	square *square
	level  int
}

func solve(first, target *square) (bool, int) {
	solver := make(chan step, 200)
	defer close(solver)

	visited := make(map[*square]bool)
	visited[first] = true

	solver <- step{square: first, level: 0}

	for sqr := range solver {
		if sqr.square == target {
			return true, sqr.level
		}

		for _, visitable := range sqr.square.visitables {
			if _, ok := visited[visitable]; !ok {
				visited[visitable] = true
				solver <- step{square: visitable, level: sqr.level + 1}
			}
		}
	}

	return false, -1
}

func solve2(lowest []*square, target *square) (bool, int) {
	solver := make(chan step, 3000)
	defer close(solver)

	visited := make(map[*square]bool)

	for _, sqr := range lowest {
		visited[sqr] = true
		solver <- step{square: sqr, level: 0}
	}

	for sqr := range solver {
		if sqr.square == target {
			return true, sqr.level
		}

		for _, visitable := range sqr.square.visitables {
			if _, ok := visited[visitable]; !ok {
				visited[visitable] = true
				solver <- step{square: visitable, level: sqr.level + 1}
			}
		}
	}

	return false, -1
}
