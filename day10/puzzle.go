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

const day = "10"

func part1() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	total := 0
	cycle := 1
	X := 1
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			total += cycle * X
		}
		if line == "noop" {
			cycle++
			continue
		}
		addr := strings.Split(line, " ")
		v, err := strconv.Atoi(addr[1])
		if err != nil {
			panic(err)
		}
		if cycle == 19 || cycle == 59 || cycle == 99 || cycle == 139 || cycle == 179 || cycle == 219 {
			total += (cycle + 1) * X
		}
		cycle += 2
		X += v
	}
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total) + " (X=" + strconv.Itoa(X) + ")")
}
func part2() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	cycle := 1
	X := 1
	crt := ""
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "noop" {
			crt = draw(crt, &cycle, X)
			cycle++
		} else {
			addr := strings.Split(line, " ")
			v, err := strconv.Atoi(addr[1])
			if err != nil {
				panic(err)
			}
			crt = draw(crt, &cycle, X)
			cycle++
			crt = draw(crt, &cycle, X)
			cycle++
			X += v
		}
	}
	fmt.Println("Day#" + day + ".2 : \n" + crt)
}

func draw(crt string, cycle *int, X int) string {
	if *cycle >= X && *cycle <= X+2 {
		crt += "#"
	} else {
		crt += "."
	}
	if *cycle == 40 {
		crt += "\n"
		*cycle = 0
	}
	return crt
}
