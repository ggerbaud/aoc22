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

const day = "17"

func part1() {
	jets := getData()
	//jets := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	total := simulate(2022, jets)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
}

func part2() {
	jets := getData()
	//jets := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

	total := simulate(1000000000000, jets)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func simulate(target int, jets string) int {
	emptyLine := [7]int{0, 0, 0, 0, 0, 0, 0}
	ground := [7]int{1, 1, 1, 1, 1, 1, 1}
	land := [][7]int{ground}
	steamIdx := 0
	rockIdx := 0
	ok := true
	skip := 0
	patterns := make(map[[7]int]map[int]map[int][2]int)
	patterns[ground] = make(map[int]map[int][2]int)
	patterns[ground][0] = make(map[int][2]int)
	patterns[ground][rockIdx][steamIdx] = [2]int{0, 0}
	for i := 0; i < target; i++ {
		land = append(land, emptyLine, emptyLine, emptyLine)
		rp := [2]int{2, len(land)}
		r := rocks[rockIdx]
		for _, sh := range r.shape {
			var line [7]int
			copy(line[2:], sh)
			land = append(land, line)
		}
		for {
			steam := jets[steamIdx]
			steamIdx = (steamIdx + 1) % len(jets)
			land, rp = steamJet(land, r, rp, steam)
			ok, land, rp = moveDown(land, r, rp)
			if !ok {
				break
			}
		}
		last := land[len(land)-1]
		if _, ok := patterns[last]; !ok {
			patterns[last] = make(map[int]map[int][2]int)
		}
		if s2, ok := patterns[last][rockIdx]; !ok {
			patterns[last][rockIdx] = make(map[int][2]int)
		} else if pattern, ok := s2[steamIdx]; ok {
			if skip == 0 {
				size := len(land) - 1
				cycle := i - pattern[0]
				jump := (target - i) / cycle
				dsize := size - pattern[1]
				skip = jump * dsize
				i += jump * cycle
			}
		}
		patterns[last][rockIdx][steamIdx] = [2]int{i, len(land) - 1}
		rockIdx = (rockIdx + 1) % len(rocks)
	}
	return len(land) - 1 + skip
}

func moveDown(land [][7]int, r rock, rp [2]int) (bool, [][7]int, [2]int) {
	if rp[1] <= 1 {
		return false, land, rp
	}
	out := make([][7]int, 0)
	for i, l := range land {
		out = append(out, [7]int{})
		for j := 0; j < 7; j++ {
			out[i][j] = l[j]
		}
	}
	for i, sh := range r.shape {
		prev := out[rp[1]+i-1]
		line := out[rp[1]+i]
		for j, pt := range sh {
			prev[rp[0]+j] += pt
			line[rp[0]+j] -= pt
		}
		for _, v := range prev {
			if v > 1 {
				return false, land, rp
			}
		}
		out[rp[1]+i-1] = prev
		out[rp[1]+i] = line
	}
	rp[1]--
	for {
		n := len(out) - 1
		last := out[n]
		empty := true
		for _, v := range last {
			if v > 0 {
				empty = false
				break
			}
		}
		if !empty {
			break
		}
		out = out[:n]
	}
	return true, out, rp
}

func steamJet(land [][7]int, r rock, rp [2]int, steam uint8) ([][7]int, [2]int) {
	if steam == '<' {
		return steamJetLeft(land, r, rp)
	}
	if steam == '>' {
		return steamJetRight(land, r, rp)
	}
	return land, rp
}

func steamJetLeft(land [][7]int, r rock, rp [2]int) ([][7]int, [2]int) {
	if rp[0] == 0 {
		return land, rp
	}
	out := make([][7]int, 0)
	for i, l := range land {
		out = append(out, [7]int{})
		for j := 0; j < 7; j++ {
			out[i][j] = l[j]
		}
	}
	for i, sh := range r.shape {
		line := out[rp[1]+i]
		for j, pt := range sh {
			line[rp[0]+j-1] += pt
			line[rp[0]+j] -= pt
		}
		for _, v := range line {
			if v > 1 {
				return land, rp
			}
		}
		out[rp[1]+i] = line
	}
	rp[0]--
	return out, rp
}

func steamJetRight(land [][7]int, r rock, rp [2]int) ([][7]int, [2]int) {
	if rp[0]+len(r.shape[0]) == 7 {
		return land, rp
	}
	out := make([][7]int, 0)
	for i, l := range land {
		out = append(out, [7]int{})
		for j := 0; j < 7; j++ {
			out[i][j] = l[j]
		}
	}
	for i, sh := range r.shape {
		line := out[rp[1]+i]
		for j := len(sh) - 1; j >= 0; j-- {
			line[rp[0]+j+1] += sh[j]
			line[rp[0]+j] -= sh[j]
		}
		for _, v := range line {
			if v > 1 {
				return land, rp
			}
		}
		out[rp[1]+i] = line
	}
	rp[0]++
	return out, rp

}

type rock struct {
	shape [][]int
}

var minusR = rock{
	shape: [][]int{{1, 1, 1, 1}},
}

var plusR = rock{
	shape: [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}},
}

var noLR = rock{
	shape: [][]int{{1, 1, 1}, {0, 0, 1}, {0, 0, 1}},
}

var iR = rock{
	shape: [][]int{{1}, {1}, {1}, {1}},
}

var cubeR = rock{
	shape: [][]int{{1, 1}, {1, 1}},
}

var rocks = [5]rock{minusR, plusR, noLR, iR, cubeR}

func getData() string {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	return fileScanner.Text()
}
