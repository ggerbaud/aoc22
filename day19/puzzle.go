package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	_blueprints := buildBluePrints()
	part1(_blueprints)
	part2(_blueprints)
}

const day = "19"
const test = false

func part1(_blueprints map[int]*blueprint) {
	total := 0
	for i, b := range _blueprints {
		costs := [6]int{b.robots[ore].oreQ, b.robots[clay].oreQ, b.robots[obsidian].oreQ, b.robots[obsidian].clayQ, b.robots[geode].oreQ, b.robots[geode].obsidianQ}
		total += i * recone(costs, 0, 0, 0, 0, 1, 0, 0, 0, 24)
	}
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
}
func part2(_blueprints map[int]*blueprint) {
	total := 1
	for i, b := range _blueprints {
		if i <= 3 {
			costs := [6]int{b.robots[ore].oreQ, b.robots[clay].oreQ, b.robots[obsidian].oreQ, b.robots[obsidian].clayQ, b.robots[geode].oreQ, b.robots[geode].obsidianQ}
			q := recone(costs, 0, 0, 0, 0, 1, 0, 0, 0, 32)
			total *= q
		}
	}
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func buildBluePrints() map[int]*blueprint {
	lines := utils.ReadFileLinesForDay(day, test)

	_blueprints := make(map[int]*blueprint)
	for _, line := range lines {
		if len(line) > 0 {
			rest := strings.Split(line, ":")
			id := utils.ParseInt(strings.Split(rest[0], " ")[1])
			_blueprints[id] = &blueprint{id: id, robots: make(map[kind]robot), upcosts: make(map[kind]int)}
			rest = strings.Split(rest[1], ".")
			for _, r := range rest {
				if len(r) > 1 {
					_robot := robot{}
					data := strings.Split(strings.TrimSpace(r)[5:], "robot")
					_robot.kind = kindOf(strings.TrimSpace(data[0]))
					for _, c := range strings.Split(strings.TrimSpace(data[1])[6:], "and") {
						cdata := strings.Split(strings.TrimSpace(strings.Trim(c, ".")), " ")
						many := utils.ParseInt(cdata[0])
						of := kindOf(cdata[1])
						switch of {
						case "ore":
							_robot.oreQ = many
						case "clay":
							_robot.clayQ = many
						case "obsidian":
							_robot.obsidianQ = many
						}
					}
					_blueprints[id].robots[_robot.kind] = _robot
				}
			}
		}
	}
	return _blueprints
}

var bestGeodes = 0

func recone(costs [6]int, or, cl, ob, ge, orr, clr, obr, ger, t int) int {
	if t == 0 {
		return ge
	}

	thMax := ge
	for i := 0; i < t; i++ {
		thMax += ger + i
	}
	if thMax < bestGeodes {
		return 0
	}

	nor := or + orr
	ncl := cl + clr
	nob := ob + obr
	nge := ge + ger

	if or >= costs[4] && ob >= costs[5] {
		return recone(costs, nor-costs[4], ncl, nob-costs[5], nge, orr, clr, obr, ger+1, t-1)
	}
	if clr >= costs[3] && obr < costs[5] && or >= costs[2] && cl >= costs[3] {
		return recone(costs, nor-costs[2], ncl-costs[3], nob, nge, orr, clr, obr+1, ger, t-1)
	}

	best := 0
	if obr < costs[5] && or >= costs[2] && cl >= costs[3] {
		best = utils.Max(best, recone(costs, nor-costs[2], ncl-costs[3], nob, nge, orr, clr, obr+1, ger, t-1))
	}
	if clr < costs[3] && or >= costs[1] {
		best = utils.Max(best, recone(costs, nor-costs[1], ncl, nob, nge, orr, clr+1, obr, ger, t-1))
	}
	if orr < 4 && or >= costs[0] {
		best = utils.Max(best, recone(costs, nor-costs[0], ncl, nob, nge, orr+1, clr, obr, ger, t-1))
	}
	if or <= 4 {
		best = utils.Max(best, recone(costs, nor, ncl, nob, nge, orr, clr, obr, ger, t-1))
	}
	return best
}

type blueprint struct {
	id      int
	robots  map[kind]robot
	upcosts map[kind]int
}

type robot struct {
	kind                   kind
	oreQ, clayQ, obsidianQ int
}

type kind string

const (
	ore      kind = "ore"
	clay     kind = "clay"
	obsidian kind = "obsidian"
	geode    kind = "geode"
)

var kinds = [4]kind{geode, obsidian, clay, ore}

func kindOf(s string) kind {
	for _, k := range kinds {
		if s == string(k) {
			return k
		}
	}
	panic("unknown")
}
