package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	readFile, err := os.Open("./day5/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	crateMap := make([]string, 3)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			crateMap = append([]string{line}, crateMap...)
		} else {
			break
		}
	}
	// parse crates
	stacks := parseCrateMap(crateMap)
	re := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			groups := re.FindAllStringSubmatch(line, -1)[0]
			n, err := strconv.Atoi(groups[1])
			if err != nil {
				panic(err)
			}
			from, err := strconv.Atoi(groups[2])
			if err != nil {
				panic(err)
			}
			to, err := strconv.Atoi(groups[3])
			if err != nil {
				panic(err)
			}
			moveNFromTo(stacks, n, from-1, to-1)
		}
	}
	fmt.Println("Day#5.1 : " + getTops(stacks))
}
func part2() {
	readFile, err := os.Open("./day5/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	crateMap := make([]string, 3)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			crateMap = append([]string{line}, crateMap...)
		} else {
			break
		}
	}
	// parse crates
	stacks := parseCrateMap(crateMap)
	re := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			groups := re.FindAllStringSubmatch(line, -1)[0]
			n, err := strconv.Atoi(groups[1])
			if err != nil {
				panic(err)
			}
			from, err := strconv.Atoi(groups[2])
			if err != nil {
				panic(err)
			}
			to, err := strconv.Atoi(groups[3])
			if err != nil {
				panic(err)
			}
			moveNFromTo9001(stacks, n, from-1, to-1)
		}
	}
	fmt.Println("Day#5.2 : " + getTops(stacks))
}

func parseCrateMap(crateMap []string) [][]string {
	stacks := strings.Split(crateMap[0], " ")
	n, err := strconv.Atoi(stacks[len(stacks)-1])
	if err != nil {
		panic(err)
	}
	result := make([][]string, n, n)
	for _, e := range crateMap[1:] {
		stack := 0
		for j := 0; j < len(e); j++ {
			if e[j] == '[' {
				result[stack] = append(result[stack], string(e[j+1]))
			}
			j += 3
			stack++
		}
	}
	return result
}

func moveNFromTo9001(stacks [][]string, n, from, to int) {
	stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-n:]...)
	stacks[from] = stacks[from][:len(stacks[from])-n]
}

func moveNFromTo(stacks [][]string, n, from, to int) {
	for i := 0; i < n; i++ {
		moveHeadFromTo(stacks, from, to)
	}
}

func moveHeadFromTo(stacks [][]string, from, to int) {
	stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
	stacks[from] = stacks[from][:len(stacks[from])-1]
}
func getTops(stacks [][]string) string {
	result := ""
	for _, stack := range stacks {
		result += stack[len(stack)-1]
	}
	return result
}
