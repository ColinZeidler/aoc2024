package main

import (
	"os"
	"strings"
)

var wall rune = '#'
var visited rune = 'X'
var empty rune = '.'
var guards string = "<^>v"

type GRow []rune
type GMap []GRow

func main() {
	part1()
	part2()
}

func part1() {
	println("Part1")
	input, ok := os.ReadFile("day6-input.txt")
	// input, ok := os.ReadFile("small.txt")
	if ok != nil {
		panic("Error reading input")
	}
	inputstr := string(input)
	/*
		while guard exists (lookup pos)
		if blocked turn right
		else move forward
	*/

	guardMap := make(GMap, 0)
	for _, row := range strings.Split(inputstr, "\n") {
		if len(row) == 0 {
			continue
		}
		maprow := make(GRow, 0, len(row))

		for _, tile := range row {
			maprow = append(maprow, tile)
		}
		guardMap = append(guardMap, maprow)
	}

	for x_pos, y_pos := find_guard(guardMap); x_pos != -1; x_pos, y_pos = find_guard(guardMap) {
		if guard_blocked(guardMap, x_pos, y_pos) {
			guardMap = turn_guard(guardMap, x_pos, y_pos)
		} else {
			guardMap = move_forward(guardMap, x_pos, y_pos)
		}
		// println("")
		// printMap(guardMap)
	}

	println("")
	printMap(guardMap)
	count := count_positions(guardMap)
	println(count)

}

func part2() {
	println("Part2")
	/*
		placing an Obstruction
		Only need to check positions the original path would travel

		find a spot where two paths cross, and place Ob one position past it.
	*/
}

func printMap(guardMap GMap) {
	for _, row := range guardMap {
		println(string(row))
	}
}

func find_guard(guardMap GMap) (int, int) {
	found := false
	y_pos, x_pos := -1, -1
	for y, row := range guardMap {
		y_pos = y
		for x, tile := range row {
			x_pos = x
			for _, guard := range guards {
				if tile == guard {
					found = true
					break
				}
			}
			if found {
				break
			}

		}
		if found {
			break
		}
	}
	if found {
		return x_pos, y_pos
	}
	return -1, -1
}

func guard_blocked(guardMap GMap, x_pos int, y_pos int) bool {
	// check icon to determine direction
	guard := guardMap[y_pos][x_pos]
	x_dir, y_dir := get_direction(guard)

	new_x := x_pos + x_dir
	new_y := y_pos + y_dir
	if 0 <= new_x && new_x < len(guardMap[y_pos]) &&
		0 <= new_y && new_y < len(guardMap) {
		if guardMap[new_y][new_x] == wall {
			return true
		}
	}
	return false
}

func turn_guard(guardMap GMap, x_pos int, y_pos int) GMap {
	guard := guardMap[y_pos][x_pos]
	var newGuard rune
	for index, icon := range guards {
		next_index := index + 1
		if next_index >= len(guards) {
			next_index = 0
		}
		if icon == guard {

			newGuard = rune(guards[next_index])
			break
		}
	}
	// replace guard with next icon
	guardMap[y_pos][x_pos] = newGuard
	return guardMap
}

func move_forward(guardMap GMap, x_pos int, y_pos int) GMap {
	// check icon to determine direction
	guard := guardMap[y_pos][x_pos]
	x_dir, y_dir := get_direction(guard)

	new_x := x_pos + x_dir
	new_y := y_pos + y_dir
	if 0 <= new_x && new_x < len(guardMap[y_pos]) &&
		0 <= new_y && new_y < len(guardMap) {
		guardMap[new_y][new_x] = guard
	}
	// move guard
	//mark old pos with X
	guardMap[y_pos][x_pos] = visited
	return guardMap
}

func get_direction(guard rune) (int, int) {
	// returns x_dir,y_dir
	switch guard {
	case '>':
		return 1, 0
	case 'v':
		return 0, 1
	case '<':
		return -1, 0
	case '^':
		return 0, -1
	}
	return 0, 0
}

func count_positions(guardMap GMap) int {
	count := 0
	for _, row := range guardMap {
		for _, tile := range row {
			if tile == visited {
				count += 1
			}
		}
	}
	return count
}
