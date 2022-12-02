package main

import (
	"advent/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const day = "1"

func main() {
	part1()
	part2()
}

func part1() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	max := 0
	current := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) == 0 {
			if current > max {
				max = current
			}
			current = 0
		} else {
			current += utils.ParseInt(line)
		}
	}
	if current > max {
		max = current
	}
	fmt.Println("Day#1.1 : " + strconv.Itoa(max))
}

func part2() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	low, middle, max := 0, 0, 0
	current := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) == 0 {
			if current > low {
				low, middle, max = handle3max(low, middle, max, current)
			}
			current = 0
		} else {
			current += utils.ParseInt(line)
		}
	}
	if current > low {
		low, middle, max = handle3max(low, middle, max, current)
	}
	fmt.Println("Day#1.2 : " + strconv.Itoa(low+middle+max))
}

func handle3max(a, b, c, d int) (int, int, int) {
	if d > c {
		return b, c, d
	}
	if d > b {
		return b, d, c
	}
	if d > a {
		return d, b, c
	}
	return a, b, c
}
