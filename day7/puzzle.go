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

const day = "7"

func part1() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	root := walk(fileScanner)
	total := addSize(root)
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

	root := walk(fileScanner)
	calcSize(root)
	free := 70000000 - root.size
	need := 30000000 - free
	if need <= 0 {
		panic(fmt.Errorf("something wrond, need %d", need))
	}
	smallest := smallestToDelete(root, need, root.size)

	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(smallest))
}

func walk(fileScanner *bufio.Scanner) *dir {
	root := &dir{"/", false, 0, 0, nil, nil}
	cwd := root
	listing := false
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			if line == "$ cd /" {
				listing = false
				cwd = root
			} else if line == "$ cd .." {
				listing = false
				if cwd.parent == nil {
					panic(fmt.Errorf("no parent directory for %s", cwd.name))
				}
				cwd = cwd.parent
			} else if strings.HasPrefix(line, "$ cd ") {
				listing = false
				name := line[5:]
				if i, ok := utils.IndexKey(name, cwd.deps, func(t *dir) string { return t.name }); ok {
					cwd = cwd.deps[i]
				} else {
					nd := &dir{name, false, 0, 0, nil, cwd}
					cwd.deps = append(cwd.deps, nd)
					cwd = nd
				}
			} else if line == "$ ls" {
				listing = true
			} else if listing {
				if strings.HasPrefix(line, "dir ") {
					name := line[4:]
					if _, ok := utils.IndexKey(name, cwd.deps, func(t *dir) string { return t.name }); !ok {
						nd := &dir{name, false, 0, 0, nil, cwd}
						cwd.deps = append(cwd.deps, nd)
					}
				} else {
					file := strings.Split(line, " ")
					size, err := strconv.Atoi(file[0])
					if err != nil {
						panic(err)
					}
					cwd.directSize += size
				}
			} else {
				panic(fmt.Errorf("wtf %s", line))
			}
		}
	}
	return root
}

func addSize(d *dir) int {
	total := 0
	calcSize(d)
	if d.size <= 100000 {
		total += d.size
	}
	for _, sub := range d.deps {
		total += addSize(sub)
	}
	return total
}
func calcSize(d *dir) {
	if !d.calc {
		if len(d.deps) > 0 {
			for _, sub := range d.deps {
				calcSize(sub)
				d.size += sub.size
			}
		}
		d.size += d.directSize
		d.calc = true
	}
}

func smallestToDelete(d *dir, need, target int) int {
	result := target
	if need == target {
		return target
	}
	if d.size < need {
		return target
	} else if d.size == need {
		return d.size
	} else {
		result = d.size
		for _, sub := range d.deps {
			nt := smallestToDelete(sub, need, result)
			if nt < result {
				result = nt
			}
		}
	}
	return result
}

type dir struct {
	name       string
	calc       bool
	size       int
	directSize int
	deps       []*dir
	parent     *dir
}
