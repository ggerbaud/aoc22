package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//part1()
	part2()
}

func part1() {
	readFile, err := os.Open("./day3/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	score := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			n := len(line) / 2
			h1 := line[0:n]
			h2 := line[n:]
			common := intersection(h1, h2)
			result := 0
			for i := range common {
				result += i
			}
			score += result
		}
	}
	fmt.Println("Day#3.1 : " + strconv.Itoa(score))
}
func part2() {
	readFile, err := os.Open("./day3/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	score := 0
	count := 0
	var s map[int]struct{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			count++
			if count == 1 {
				s = intersection(line, line)
			} else {
				s = intersection2(s, line)
			}
			if count == 3 {
				result := 0
				for i := range s {
					result += i
				}
				score += result
				count = 0
			}
		}
	}
	fmt.Println("Day#3.2 : " + strconv.Itoa(score))
}

func intersection2(m map[int]struct{}, a string) map[int]struct{} {
	s := make(map[int]struct{})
	for _, item := range a {
		n := ctoi(item)
		if _, ok := m[n]; ok {
			s[n] = struct{}{}
		}
	}
	return s
}

func intersection(a, b string) map[int]struct{} {
	s := make(map[int]struct{})
	m := make(map[int]bool)

	for _, item := range a {
		m[ctoi(item)] = true
	}

	for _, item := range b {
		n := ctoi(item)
		if _, ok := m[n]; ok {
			s[n] = struct{}{}
		}
	}
	return s
}

func ctoi(c int32) int {
	if c < 'a' {
		return int(c - 'A' + 27)
	}
	return int(c - 'a' + 1)
}
