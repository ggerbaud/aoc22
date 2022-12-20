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

const day = "16"

func part1() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	valves := make(map[string]*valve)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			label := line[6:8]
			rate := utils.ParseInt(strings.Split(strings.Split(line, ";")[0], "=")[1])
			v := getOrCreate(valves, label)
			v.rate = rate
			leads := strings.Split(line, "to valve")[1]
			if leads[:1] == "s" {
				leads = leads[1:]
			}
			leads = strings.TrimSpace(leads)
			for _, l := range strings.Split(leads, ",") {
				lv := getOrCreate(valves, strings.TrimSpace(l))
				v.connect = append(v.connect, lv)
			}
		}
	}

	rates := make(map[string]int)
	dists := make(map[string]map[string]int)
	for _, v := range valves {
		if v.rate > 0 || v.label == "AA" {
			rates[v.label] = v.rate
			for _, v2 := range valves {
				if v2.rate > 0 {
					_, p := shortestPath(v, v2, make(map[string]bool))
					if dists[v.label] == nil {
						dists[v.label] = make(map[string]int)
					}
					dists[v.label][v2.label] = len(p[0])
				}
			}
		}
	}
	best := scoreFor1(rates, dists)
	fmt.Println("Day#" + day + ".1.2 : " + strconv.Itoa(best))
}
func part2() {
	readFile, err := os.Open("./day" + day + "/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	valves := make(map[string]*valve)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			label := line[6:8]
			rate := utils.ParseInt(strings.Split(strings.Split(line, ";")[0], "=")[1])
			v := getOrCreate(valves, label)
			v.rate = rate
			leads := strings.Split(line, "to valve")[1]
			if leads[:1] == "s" {
				leads = leads[1:]
			}
			leads = strings.TrimSpace(leads)
			for _, l := range strings.Split(leads, ",") {
				lv := getOrCreate(valves, strings.TrimSpace(l))
				v.connect = append(v.connect, lv)
			}
		}
	}

	rates := make(map[string]int)
	dists := make(map[string]map[string]int)
	for _, v := range valves {
		if v.rate > 0 || v.label == "AA" {
			rates[v.label] = v.rate
			for _, v2 := range valves {
				if v2.rate > 0 {
					_, p := shortestPath(v, v2, make(map[string]bool))
					if dists[v.label] == nil {
						dists[v.label] = make(map[string]int)
					}
					dists[v.label][v2.label] = len(p[0])
				}
			}
		}
	}
	best := scoreFor2(rates, dists)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(best))
}
func scoreFor1(rates map[string]int, dists map[string]map[string]int) int {
	type key struct {
		remain int
		pos    string
	}
	type state struct {
		k      key
		total  int
		opened []string
	}
	best := 0
	states := []state{{key{31, "AA"}, 0, make([]string, 0)}}
	//loop:
	for len(states) > 0 {
		s := states[0]
		states = states[1:]

		if s.k.remain <= 0 {
			best = utils.Max(best, s.total)
			continue
		}

		remain := s.k.remain - 1
		opened := make([]string, len(s.opened))
		copy(opened, s.opened)
		ns := s.total
		opened = append(opened, s.k.pos)
		ns += remain * rates[s.k.pos]
		newstates := false
		for conn, d := range dists[s.k.pos] {
			if _, ok := utils.Index(conn, opened); !ok {
				newstates = true
				remain := remain
				if d > remain {
					remain = 0
				} else {
					remain -= d
				}
				states = append(states, state{key{remain, conn}, ns, opened})
			}
		}
		if !newstates {
			best = utils.Max(best, ns)
		}
	}
	return best
}
func scoreFor2(rates map[string]int, dists map[string]map[string]int) int {
	best := 0
	states := []state{{27, []pos{{dst: "AA"}, {dst: "AA"}}, 0, make([]string, 0)}}

	vIdx := make(map[string]int)
	idx := 0
	for k, _ := range rates {
		vIdx[k] = idx
		idx++
	}
	seen := make(map[int]map[int]int)

	for len(states) > 0 {
		s := states[0]
		states = states[1:]

		if s.remain <= 0 || (s.pos[0].stop && s.pos[1].stop) {
			best = utils.Max(best, s.total)
			continue
		}

		opIdx := computeOpened(s.opened, vIdx)
		if saw, ok := seen[opIdx]; !ok {
			seen[opIdx] = make(map[int]int)
		} else {
			if sc, ok := saw[s.remain]; ok && sc > s.total {
				continue
			}
		}
		seen[opIdx][s.remain] = s.total

		remain := s.remain - 1
		opened := make([]string, len(s.opened))
		copy(opened, s.opened)
		ns := s.total
		for i := 0; i < len(s.pos); i++ {
			if !s.pos[i].stop {
				if !s.pos[i].move {
					opened = append(opened, s.pos[i].dst)
					ns += remain * rates[s.pos[i].dst]
				} else {
					s.pos[i].d--
				}
			}
		}
		mvs0 := moves(s.pos[0], opened, dists)
		mvs1 := moves(s.pos[1], opened, dists)
		newstates := false
		for _, m0 := range mvs0 {
			for _, m1 := range mvs1 {
				if m0.dst != m1.dst {
					newstates = true
					remain := remain
					minD := utils.Min(m0.d, m1.d)
					if minD > remain {
						remain = 0
					} else {
						remain -= minD
					}
					p0, p1 := pos{dst: m0.dst, d: m0.d - minD, move: m0.d > minD}, pos{dst: m1.dst, d: m1.d - minD, move: m1.d > minD}
					news := state{remain, []pos{p0, p1}, ns, opened}
					states = append(states, news)
				}
			}
		}
		if !newstates {
			if s.pos[0].move {
				remain := remain - s.pos[0].d
				news := state{remain, []pos{{dst: s.pos[0].dst}, {stop: true}}, ns, opened}
				states = append(states, news)
			} else if s.pos[1].move {
				remain := remain - s.pos[1].d
				news := state{remain, []pos{{dst: s.pos[1].dst}, {stop: true}}, ns, opened}
				states = append(states, news)
			} else {
				best = utils.Max(best, ns)
			}
		}
	}
	return best
}

func computeOpened(opened []string, vIdx map[string]int) int {
	x := 0
	for _, v := range opened {
		x |= 1 << vIdx[v]
	}
	return x
}

func moves(x pos, opened []string, dists map[string]map[string]int) []pos {
	if x.stop || x.move {
		return []pos{x}
	}
	var mvs []pos
	for conn, d := range dists[x.dst] {
		if _, ok := utils.Index(conn, opened); !ok {
			mvs = append(mvs, pos{dst: conn, d: d, move: false})
		}
	}
	return mvs
}

func shortestPath(from, to *valve, visited map[string]bool) (int, [][]string) {
	if _, ok := utils.Index(to, from.connect); ok {
		return 1, [][]string{{to.label}}
	}
	best := 999999999
	var bestPaths [][]string
	for _, c := range from.connect {
		if visit, ok := visited[c.label]; !ok || !visit {
			visited[from.label] = true
			l, p := shortestPath(c, to, visited)
			if l == best {
				for _, sp := range p {
					bestPaths = append(bestPaths, append([]string{c.label}, sp...))
				}
			} else if l < best {
				best = l
				bestPaths = make([][]string, 0)
				for _, sp := range p {
					bestPaths = append(bestPaths, append([]string{c.label}, sp...))
				}
			}
			visited[from.label] = false
		}
	}
	return best + 1, bestPaths
}

type valve struct {
	label   string
	rate    int
	connect []*valve
}

type result struct {
	label string
	best  int
	next  *result
}
type pos struct {
	dst  string
	d    int
	move bool
	stop bool
}
type state struct {
	remain int
	pos    []pos
	total  int
	opened []string
}

func getOrCreate(valves map[string]*valve, label string) *valve {
	if v, ok := valves[label]; ok {
		return v
	}
	v := &valve{label: label}
	valves[label] = v
	return v
}
