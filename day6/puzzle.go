package main

import (
	"advent/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	readFile, err := os.Open("./day6/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	cpt := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			cpt = getStartOfPacket(line)
			fmt.Println("Day#6.1 : " + strconv.Itoa(cpt))
		}
	}
}

func part2() {
	readFile, err := os.Open("./day6/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	cpt := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			cpt = getStartOfMessage(line)
			fmt.Println("Day#6.2 : " + strconv.Itoa(cpt))
		}
	}
}

func getStartOfPacket(input string) int {
	chars := make([]int32, 0)
	for i, c := range input {
		if len(chars) < 4 {
			chars = append(chars, c)
			continue
		} else {
			if allDifferent(chars) {
				return i
			}
			chars = append(chars[1:], c)
		}
	}
	panic("nope")
}

func getStartOfMessage(input string) int {
	chars := make([]int32, 0)
	for i, c := range input {
		if len(chars) < 14 {
			chars = append(chars, c)
			continue
		} else {
			if allDifferent(chars) {
				return i
			}
			chars = append(chars[1:], c)
		}
	}
	panic("nope")
}

func allDifferent(chars []int32) bool {
	tmp := make([]int32, len(chars))
	for _, c := range chars {
		if _, ok := utils.Index(c, tmp); ok {
			return false
		}
		tmp = append(tmp, c)
	}
	return true
}
