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

const day = "18"

func part1() {
	//readFile, err := os.Open("./day" + day + "/test.txt")
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	minX, minY, minZ, maxX, maxY, maxZ := 999999, 999999, 999999, 0, 0, 0
	var data []point
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			coords := strings.Split(line, ",")
			xyz := point{utils.ParseInt(coords[0]), utils.ParseInt(coords[1]), utils.ParseInt(coords[2])}
			if xyz[0] < minX {
				minX = xyz[0]
			}
			if xyz[0] > maxX {
				maxX = xyz[0]
			}
			if xyz[1] < minY {
				minY = xyz[1]
			}
			if xyz[1] > maxY {
				maxY = xyz[1]
			}
			if xyz[2] < minZ {
				minZ = xyz[2]
			}
			if xyz[2] > maxZ {
				maxZ = xyz[2]
			}
			// min max
			data = append(data, xyz)
		}
	}
	total := surface(data)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
}
func part2() {
	//readFile, err := os.Open("./day" + day + "/test.txt")
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	minX, minY, minZ, maxX, maxY, maxZ := 999999, 999999, 999999, 0, 0, 0
	var data []point
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			coords := strings.Split(line, ",")
			xyz := point{utils.ParseInt(coords[0]), utils.ParseInt(coords[1]), utils.ParseInt(coords[2])}
			if xyz[0] < minX {
				minX = xyz[0]
			}
			if xyz[0] > maxX {
				maxX = xyz[0]
			}
			if xyz[1] < minY {
				minY = xyz[1]
			}
			if xyz[1] > maxY {
				maxY = xyz[1]
			}
			if xyz[2] < minZ {
				minZ = xyz[2]
			}
			if xyz[2] > maxZ {
				maxZ = xyz[2]
			}
			data = append(data, xyz)
		}
	}

	empties := make([]point, 0)
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			for z := 0; z <= maxZ; z++ {
				empties = append(empties, point{x, y, z})
			}
		}
	}
	for _, p := range data {
		empties, _ = utils.Delete(p, empties)
	}

	pockets := make([][]point, 0)
	for len(empties) > 0 {
		p := empties[0]
		checks := []point{p}
		bubble := make([]point, 0)
		for len(checks) > 0 {
			check := checks[0]
			checks = checks[1:]
			if _, ok := utils.Index(check, empties); ok {
				bubble = append(bubble, check)
				empties, _ = utils.Delete(check, empties)
				for _, mask := range masks {
					checks = append(checks, add(check, mask))
				}
			}
		}
		if _, ok := utils.Index(point{0, 0, 0}, bubble); !ok {
			pockets = append(pockets, bubble)
		}
	}

	total := surface(data)
	for _, pocket := range pockets {
		total -= surface(pocket)
	}
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

var masks = []point{
	{-1, 0, 0},
	{1, 0, 0},
	{0, -1, 0},
	{0, 1, 0},
	{0, 0, -1},
	{0, 0, 1},
}

func surface(data []point) int {
	total := 0
	for _, pt := range data {
		here := 6
		for _, mask := range masks {
			if _, ok := utils.Index(add(pt, mask), data); ok {
				here--
			}
		}
		total += here
	}
	return total
}

func add(p1, p2 point) point {
	return point{p1[0] + p2[0], p1[1] + p2[1], p1[2] + p2[2]}
}

type point [3]int
