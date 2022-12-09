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

const day = "8"

func part1() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var forest [][]int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			var fl []int
			for _, d := range line {
				fl = append(fl, int(d))
			}
			forest = append(forest, fl)
		}
	}
	visible := 0
	for i, l := range forest {
		if i == 0 || i == len(forest)-1 {
			visible += len(l)
		} else {
			for j, t := range l {
				if j == 0 || j == len(l)-1 {
					visible++
				} else {
					// before sur ligne
					vb := true
					for k := j - 1; k >= 0; k-- {
						if l[k] >= t {
							vb = false
							break
						}
					}
					if vb {
						visible++
						continue
					}
					// after sur ligne
					vb = true
					for k := j + 1; k < len(l); k++ {
						if l[k] >= t {
							vb = false
							break
						}
					}
					if vb {
						visible++
						continue
					}
					// before sur colonne
					vb = true
					for k := i - 1; k >= 0; k-- {
						if forest[k][j] >= t {
							vb = false
							break
						}
					}
					if vb {
						visible++
						continue
					}
					// after sur colonne
					vb = true
					for k := i + 1; k < len(forest); k++ {
						if forest[k][j] >= t {
							vb = false
							break
						}
					}
					if vb {
						visible++
						continue
					}
				}
			}
		}
	}
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(visible))
}
func part2() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var forest [][]int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			var fl []int
			for _, d := range line {
				fl = append(fl, int(d))
			}
			forest = append(forest, fl)
		}
	}
	score := 0
	for i, l := range forest {
		if i == 0 || i == len(forest)-1 {
			continue
		} else {
			for j, t := range l {
				if j == 0 || j == len(l)-1 {
					continue
				} else {
					localSc := 1
					cpt := 0
					// before sur ligne
					for k := j - 1; k >= 0; k-- {
						cpt++
						if l[k] >= t {
							break
						}
					}
					localSc = localSc * cpt
					cpt = 0
					// after sur ligne
					for k := j + 1; k < len(l); k++ {
						cpt++
						if l[k] >= t {
							break
						}
					}
					localSc = localSc * cpt
					cpt = 0
					// before sur colonne
					for k := i - 1; k >= 0; k-- {
						cpt++
						if forest[k][j] >= t {
							break
						}
					}
					localSc = localSc * cpt
					cpt = 0
					// after sur colonne
					for k := i + 1; k < len(forest); k++ {
						cpt++
						if forest[k][j] >= t {
							break
						}
					}
					localSc = localSc * cpt
					if localSc > score {
						score = localSc
					}
				}
			}
		}
	}
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(score))
}
