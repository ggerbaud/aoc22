package main

import (
	"advent/utils"
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

const day = "15"

func part1() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	sensors := make([]*sensor, 0)
	beacons := make(map[int]map[int]bool)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			values := strings.Split(line, "=")
			sx, sy, bx, by := utils.ParseInt(strings.Split(values[1], ",")[0]), utils.ParseInt(strings.Split(values[2], ":")[0]), utils.ParseInt(strings.Split(values[3], ",")[0]), utils.ParseInt(values[4])
			setBeacon(beacons, bx, by)
			sensor := &sensor{x: sx, y: sy, dist: dist(sx, bx, sy, by)}
			sensors = append(sensors, sensor)
		}
	}
	ty := 10
	bl := len(beacons[ty])
	minX := int(^uint(0) >> 1)
	maxX := -minX - 1
	for _, s := range sensors {
		if ok, x1, x2 := s.seeLine(ty); ok {
			if x1 < minX {
				minX = x1
			}
			if x2 > maxX {
				maxX = x2
			}
		}
	}
	total := maxX - minX + 1 - bl
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
}
func part2() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	sensors := make([]*sensor, 0)
	//beacons := make(map[int]map[int]bool)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			values := strings.Split(line, "=")
			sx, sy, bx, by := utils.ParseInt(strings.Split(values[1], ",")[0]), utils.ParseInt(strings.Split(values[2], ":")[0]), utils.ParseInt(strings.Split(values[3], ",")[0]), utils.ParseInt(values[4])
			//setBeacon(beacons, bx, by)
			sensor := &sensor{x: sx, y: sy, dist: dist(sx, bx, sy, by)}
			sensors = append(sensors, sensor)
		}
	}
	max := 4000000
	x, y := 0, 0
	for i := 0; i <= max; i++ {
		free := append([][]int{}, []int{0, max})
		for _, s := range sensors {
			if ok, x1, x2 := s.seeLine(i); ok {
				free = split(free, x1, x2)
				if len(free) == 0 {
					break
				}
			}
		}
		if len(free) > 0 {
			y = i
			x = free[0][0]
			break
		}
	}
	total := x*4000000 + y
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func dist(x1, x2, y1, y2 int) int {
	dist := 0
	if x1 < x2 {
		dist += x2 - x1
	} else {
		dist += x1 - x2
	}
	if y1 < y2 {
		dist += y2 - y1
	} else {
		dist += y1 - y2
	}
	return dist
}

func split(data [][]int, a, b int) [][]int {
	var out [][]int
	for k, i := range data {
		if a < i[0] {
			if b < i[0] {
				out = append(out, data[k:]...)
				break
			} else if b >= i[1] {
				continue
			} else {
				if i[1] >= b+1 {
					out = append(out, []int{b + 1, i[1]})
				}
				out = append(out, data[k+1:]...)
				break
			}
		} else if a > i[1] {
			out = append(out, i)
			continue
		} else {
			if a-1 >= i[0] {
				out = append(out, []int{i[0], a - 1})
			}
			if b < i[1] {
				if i[1] >= b+1 {
					out = append(out, []int{b + 1, i[1]})
				}
				out = append(out, data[k+1:]...)
				break
			} else {
				continue
			}
		}
	}
	return out
}

type sensor struct {
	x, y, dist int
}

func (s *sensor) seeLine(l int) (bool, int, int) {
	if l >= s.y-s.dist && l <= s.y+s.dist {
		cons := s.y - l
		if l > s.y {
			cons = l - s.y
		}
		delta := s.dist - cons
		return true, s.x - delta, s.x + delta
	}
	return false, 0, 0
}

func setBeacon(beacons map[int]map[int]bool, x, y int) {
	if _, ok := beacons[y]; !ok {
		beacons[y] = make(map[int]bool)
	}
	if _, ok := beacons[y][x]; !ok {
		beacons[y][x] = true
	}
}
