package main

import (
	"os"
	"strings"
)

type MapRow []rune
type TowerMap []MapRow

type Tower struct {
	x_pos int
	y_pos int
}

type Antinode Tower

func main() {
	part1()
	part2()
}

func part1() {
	println("Part1")
	input, err := os.ReadFile("day8-input.txt")
	if err != nil {
		panic("Error reading input")
	}
	inputstr := string(input)

	tower_map := make(TowerMap, 0)
	tower_map = initmap(tower_map, inputstr)

	towers := make(map[rune][]Tower)
	var max_y, max_x int
	for y, row := range tower_map {
		for x, tile := range row {
			if tile != '.' {
				new_tower := Tower{
					x_pos: x,
					y_pos: y,
				}
				towerlist, exists := towers[tile]
				if !exists {
					towerlist = make([]Tower, 0)
				}
				towerlist = append(towerlist, new_tower)
				towers[tile] = towerlist
			}
			max_x = x
		}
		max_y = y
	}

	// for each signal get all pairs
	antiNodes := make([]Antinode, 0)
	for _, list := range towers {
		temp_list := list[:]
		for len(temp_list) > 1 {
			pairSource := temp_list[0]
			temp_list = temp_list[1:]
			for _, second := range temp_list {
				anti1, anti2 := pairSource.getAntiNodes(second)
				if anti1.validPos(max_x, max_y) && !nodeExists(antiNodes, anti1) {
					antiNodes = append(antiNodes, anti1)
				}
				if anti2.validPos(max_x, max_y) && !nodeExists(antiNodes, anti2) {
					antiNodes = append(antiNodes, anti2)
				}
			}
		}
	}

	println(len(antiNodes))
}

func part2() {
	println("Part2")
	// input, err := os.ReadFile("small.txt")
	// if err != nil {
	// 	panic("Error reading input")
	// }
	// inputstr := string(input)
}

func initmap(tower_map TowerMap, inputstr string) TowerMap {
	for _, line := range strings.Split(inputstr, "\n") {
		if len(line) == 0 {
			continue
		}
		row := make(MapRow, len(line))
		for index, tile := range line {
			row[index] = tile
		}
		tower_map = append(tower_map, row)
	}
	return tower_map
}

func (tower *Tower) getAntiNodes(towerpair Tower) (Antinode, Antinode) {
	deltax := tower.x_pos - towerpair.x_pos
	deltay := tower.y_pos - towerpair.y_pos
	node_a := Antinode{
		x_pos: tower.x_pos + deltax,
		y_pos: tower.y_pos + deltay,
	}
	node_b := Antinode{
		x_pos: towerpair.x_pos + (deltax * -1),
		y_pos: towerpair.y_pos + (deltay * -1),
	}
	// fmt.Println(tower, deltax, deltay, towerpair, node_a, node_b)
	return node_a, node_b
}

func (antinode *Antinode) validPos(max_x, max_y int) bool {
	return antinode.x_pos >= 0 && antinode.x_pos <= max_x &&
		antinode.y_pos >= 0 && antinode.y_pos <= max_y
}

func (antinode *Antinode) Equals(other Antinode) bool {
	return antinode.x_pos == other.x_pos && antinode.y_pos == other.y_pos
}

func nodeExists(nodeList []Antinode, node Antinode) bool {
	for _, antiNode := range nodeList {
		if node.Equals(antiNode) {
			return true
		}
	}
	return false
}
