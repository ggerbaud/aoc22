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
	readFile, err := os.Open("./day4/input.txt")
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
			assigns := strings.Split(line, ",")
			if len(assigns) == 2 {
				d1, u1 := bornes(assigns[0])
				d2, u2 := bornes(assigns[1])
				if (d1 <= d2 && u1 >= u2) || (d1 >= d2 && u1 <= u2) {
					score++
				}
			}
		}
	}
	fmt.Println("Day#4.1 : " + strconv.Itoa(score))
}
func part2() {
	readFile, err := os.Open("./day4/input.txt")
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
			assigns := strings.Split(line, ",")
			if len(assigns) == 2 {
				d1, u1 := bornes(assigns[0])
				d2, u2 := bornes(assigns[1])
				if (d1 <= d2 && d2 <= u1) || (d1 >= d2 && d1 <= u2) {
					score++
				}
			}
		}
	}
	fmt.Println("Day#4.2 : " + strconv.Itoa(score))
}

func bornes(s string) (int, int) {
	bs := strings.Split(s, "-")
	if len(bs) == 2 {
		a, err1 := strconv.Atoi(bs[0])
		b, err2 := strconv.Atoi(bs[1])
		if err1 != nil {
			panic(err1)
		}
		if err2 != nil {
			panic(err2)
		}
		return a, b
	}
	panic(fmt.Errorf("wtong input %s", s))
}
