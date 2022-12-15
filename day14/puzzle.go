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
	fmt.Println()
}

const day = "14"

func part1() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	dm := parseMap(fileScanner)
	dm.locked = true
	total := 0
	for fall(dm.source) {
		total++
	}
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	fmt.Println(dm.print())
}
func part2() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	dm := parseMap(fileScanner)
	dm.v2 = true
	dm.locked = true
	total := 0
	for fall(dm.source) {
		total++
	}
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total+1))
	fmt.Println(dm.print())
}

func fall(source *square) bool {
	//fmt.Println(source.dm.print())
	var ref *square
	down := source.down()
	downleft := source.downleft()
	downright := source.downright()
	if down == nil {
		return false
	}
	if down.content == AIR {
		ref = down
	} else if downleft == nil {
		return false
	} else if downleft.content == AIR {
		ref = downleft
	} else if downright == nil {
		return false
	} else if downright.content == AIR {
		ref = downright
	} else {
		if source.content == SOURCE && source.dm.v2 {
			return false
		}
		return true
	}
	if source.content != SOURCE {
		source.content = AIR
	}
	ref.content = SAND
	return fall(ref)
}

func parseMap(fileScanner *bufio.Scanner) *dmap {
	dm := newDMap()
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			pairs := strings.Split(line, " -> ")
			for i := 0; i < len(pairs)-1; i++ {
				p := pairs[i]
				pn := pairs[i+1]
				xy1 := strings.Split(p, ",")
				x1 := parseInt(xy1[0])
				y1 := parseInt(xy1[1])
				xy2 := strings.Split(pn, ",")
				x2 := parseInt(xy2[0])
				y2 := parseInt(xy2[1])
				if x1 == x2 {
					minY, maxY := y1, y2
					if y1 > y2 {
						minY, maxY = y2, y1
					}
					fillDown(dm, x1, minY, maxY)
				} else if y1 == y2 {
					minX, maxX := x1, x2
					if x1 > x2 {
						minX, maxX = x2, x1
					}
					fillRight(dm, y1, minX, maxX)
				} else {
					panic(fmt.Errorf("not straight : %d,%d,%d,%d", x1, y1, x2, y2))
				}
			}
		}
	}
	return dm
}

func fillDown(dm *dmap, x, start, end int) {
	current := dm.get(x, start)
	for i := start; i < end; i++ {
		current.content = ROCK
		current = current.down()
	}
	current.content = ROCK
}

func fillRight(dm *dmap, y, start, end int) {
	current := dm.get(start, y)
	for i := start; i < end; i++ {
		current.content = ROCK
		current = current.right()
	}
	current.content = ROCK
}

type kind rune

const (
	SOURCE kind = '+'
	AIR    kind = '.'
	ROCK   kind = '#'
	SAND   kind = 'o'
)

type dmap struct {
	v2               bool
	locked           bool
	minX, maxX, maxY int
	dict             *map[string]*square
	source           *square
}

func newDMap() *dmap {
	dm := new(dmap)
	dict := make(map[string]*square)
	dm.locked = false
	dm.maxX = 500
	dm.maxY = 0
	dm.minX = dm.maxX
	dm.dict = &dict
	dm.source = &square{content: SOURCE, x: dm.minX, y: dm.maxY, dm: dm}
	dict[key(500, 0)] = dm.source
	return dm
}

func (dm *dmap) get(x, y int) *square {
	key := key(x, y)
	local := *dm.dict
	if ref, ok := local[key]; ok {
		return ref
	}
	qr := &square{content: AIR, x: x, y: y}
	if !dm.locked || (x >= dm.minX && x <= dm.maxX && y <= dm.maxY) {
		return dm.add(qr)
	}
	if dm.v2 {
		memoMaxY := dm.maxY
		if y <= dm.maxY+1 {
			dm.add(qr)
		} else if y == dm.maxY+2 {
			qr.content = ROCK
			dm.add(qr)
		} else {
			panic(fmt.Errorf("wtf y=%d, maxY=%d, memoMaxY=%d", y, dm.maxY, memoMaxY))
		}
		dm.maxY = memoMaxY
		return qr
	}
	return nil
}

func (dm *dmap) add(sqr *square) *square {
	key := key(sqr.x, sqr.y)
	local := *dm.dict
	if sqr.x < dm.minX {
		dm.minX = sqr.x
	}
	if sqr.x > dm.maxX {
		dm.maxX = sqr.x
	}
	if sqr.y > dm.maxY {
		dm.maxY = sqr.y
	}
	sqr.dm = dm
	local[key] = sqr
	return sqr
}

func (dm *dmap) print() string {
	momoLocked := dm.locked
	dm.locked = false
	y := dm.maxY
	if dm.v2 {
		y += 2
	}
	out := ""
	for i := 0; i <= y; i++ {
		for j := dm.minX; j <= dm.maxX; j++ {
			out += string(dm.get(j, i).content)
		}
		out += "\n"
	}
	dm.locked = momoLocked
	return out
}

type square struct {
	content kind
	x       int
	y       int
	dm      *dmap
}

func (sqr *square) down() *square {
	return sqr.dm.get(sqr.x, sqr.y+1)
}

func (sqr *square) left() *square {
	return sqr.dm.get(sqr.x-1, sqr.y)
}

func (sqr *square) right() *square {
	return sqr.dm.get(sqr.x+1, sqr.y)
}

func (sqr *square) downleft() *square {
	return sqr.dm.get(sqr.x-1, sqr.y+1)
}

func (sqr *square) downright() *square {
	return sqr.dm.get(sqr.x+1, sqr.y+1)
}

func key(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func parseInt(s string) int {
	k, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return k
}
