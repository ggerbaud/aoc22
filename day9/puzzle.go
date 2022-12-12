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

const day = "9"

func part1() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	tails := make(map[string]struct{})
	var exists = struct{}{}
	//rope := 1
	h := Pos{0, 0}
	t := Pos{0, 0}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			instr := strings.Split(line, " ")
			steps, err := strconv.Atoi(instr[1])
			if err != nil {
				panic(err)
			}
			for i := 0; i < steps; i++ {
				switch instr[0] {
				case "R":
					h.x++
				case "L":
					h.x--
				case "U":
					h.y++
				case "D":
					h.y--
				}
				// move tail
				if (h.x-t.x)*(h.x-t.x) > 1 || (h.y-t.y)*(h.y-t.y) > 1 {
					if h.x == t.x {
						if h.y > t.y {
							t.y++
						} else if h.y < t.y {
							t.y--
						}
					} else if h.y == t.y {
						if h.x > t.x {
							t.x++
						} else if h.x < t.x {
							t.x--
						}
					} else {
						if h.y > t.y {
							t.y++
						} else if h.y < t.y {
							t.y--
						}
						if h.x > t.x {
							t.x++
						} else if h.x < t.x {
							t.x--
						}
					}
					//tails = handleTail(&h, &t, tails)
				}
				tails[t.name()] = exists
			}
		}
	}
	count := len(tails)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(count))
}
func part2() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	tails := make(map[string]struct{})
	var exists = struct{}{}
	length := 10
	var rope []*Pos
	for i := 0; i < length; i++ {
		rope = append(rope, &Pos{0, 0})
	}
	h := rope[0]
	tails[h.name()] = exists
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			instr := strings.Split(line, " ")
			steps, err := strconv.Atoi(instr[1])
			if err != nil {
				panic(err)
			}
			for i := 0; i < steps; i++ {
				switch instr[0] {
				case "R":
					h.x++
				case "L":
					h.x--
				case "U":
					h.y++
				case "D":
					h.y--
				}
				// move tail
				for i := 1; i < length; i++ {
					h := rope[i-1]
					t := rope[i]
					if (h.x-t.x)*(h.x-t.x) > 1 || (h.y-t.y)*(h.y-t.y) > 1 {
						if h.x == t.x {
							if h.y > t.y {
								t.y++
							} else if h.y < t.y {
								t.y--
							}
						} else if h.y == t.y {
							if h.x > t.x {
								t.x++
							} else if h.x < t.x {
								t.x--
							}
						} else {
							if h.y > t.y {
								t.y++
							} else if h.y < t.y {
								t.y--
							}
							if h.x > t.x {
								t.x++
							} else if h.x < t.x {
								t.x--
							}
						}
					}
				}
				tails[rope[length-1].name()] = exists
			}
		}
	}
	count := len(tails)

	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(count))
}

type Pos struct {
	x, y int
}

func (p *Pos) name() string {
	return fmt.Sprintf("%d-%d", p.x, p.y)
}
