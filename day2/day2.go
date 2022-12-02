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

func part1() {
	readFile, err := os.Open("./day2/input2.txt")
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
			round := strings.Split(line, " ")
			score += strats1[round[0]][round[1]]
		}
	}
	fmt.Println("Day#2.1 : " + strconv.Itoa(score))
}

func part2() {
	readFile, err := os.Open("./day2/input2.txt")
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
			round := strings.Split(line, " ")
			score += strats2[round[0]][round[1]]
		}
	}
	fmt.Println("Day#2.2 : " + strconv.Itoa(score))
}

var strats1 = map[string]map[string]int{
	"A": {
		"X": 1 + 3,
		"Y": 2 + 6,
		"Z": 3 + 0,
	},
	"B": {
		"X": 1 + 0,
		"Y": 2 + 3,
		"Z": 3 + 6,
	},
	"C": {
		"X": 1 + 6,
		"Y": 2 + 0,
		"Z": 3 + 3,
	},
}

var strats2 = map[string]map[string]int{
	"A": {
		"X": 3 + 0,
		"Y": 1 + 3,
		"Z": 2 + 6,
	},
	"B": {
		"X": 1 + 0,
		"Y": 2 + 3,
		"Z": 3 + 6,
	},
	"C": {
		"X": 2 + 0,
		"Y": 3 + 3,
		"Z": 1 + 6,
	},
}
