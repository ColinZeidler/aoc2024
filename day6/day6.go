package main

import (
	"fmt"
	"os"
	"strings"
)

var wall rune = '#'
var visited rune = 'X'
var empty rune = '.'
var guards string = "<^>v"

type GRow []rune
type GMap []GRow

type TouchDir map[string]bool
type TouchMap map[string]TouchDir

func main() {
	part1()
	part2()
}

func initMap(emptyMap GMap, inputstr string) GMap {
	for _, row := range strings.Split(inputstr, "\n") {
		if len(row) == 0 {
			continue
		}
		maprow := make(GRow, 0, len(row))

		for _, tile := range row {
			maprow = append(maprow, tile)
		}
		emptyMap = append(emptyMap, maprow)
	}
	return emptyMap
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
	guardMap = initMap(guardMap, inputstr)

	for x_pos, y_pos := find_guard(guardMap); x_pos != -1; x_pos, y_pos = find_guard(guardMap) {
		if guard_blocked(guardMap, x_pos, y_pos) {
			guardMap = turn_guard(guardMap, x_pos, y_pos)
		} else {
			guardMap = move_forward(guardMap, x_pos, y_pos)
		}
		// println("")
		// printMap(guardMap)
	}

	// println("")
	// printMap(guardMap)
	count := count_positions(guardMap)
	println(count)

}

func part2() {
	println("Part2")
	input, ok := os.ReadFile("day6-input.txt")
	// input, ok := os.ReadFile("small.txt")
	if ok != nil {
		panic("Error reading input")
	}
	inputstr := string(input)

	// guardMap = initMap(guardMap, inputstr)
	/*
		placing an Obstruction
		Only need to check positions the original path would travel

		find a spot where two paths cross, and place Ob one position past it.
	*/

	c := make(chan bool)

	sMap := make(GMap, 0)
	sMap = initMap(sMap, inputstr)

	for x_pos, y_pos := find_guard(sMap); x_pos != -1; x_pos, y_pos = find_guard(sMap) {
		if guard_blocked(sMap, x_pos, y_pos) {
			sMap = turn_guard(sMap, x_pos, y_pos)
		} else {
			sMap = move_forward(sMap, x_pos, y_pos)
		}
	}

	count := 0
	channel_count := 0
	for y, row := range sMap {
		for x := range row {
			run := false
			tile := sMap[y][x]
			if tile == wall {
				continue
			}
			if tile == visited {
				run = true
			}

			if wall_nearby(x, y, sMap) {
				run = true
			}

			if run {
				channel_count += 1
				go run_check(x, y, inputstr, c)
			}
		}
	}

	println("Threads:", channel_count)
	checked := 0
	for range channel_count {
		res := <-c
		checked += 1
		if res {
			count += 1
		}
	}
	println("Checked", checked)
	println(count)
}

func wall_nearby(x int, y int, sMap GMap) bool {
	var new_x, new_y int
	// up
	new_y = y - 1
	new_x = x

	if 0 <= new_x && new_x < len(sMap[y]) &&
		0 <= new_y && new_y < len(sMap) {
		if sMap[new_y][new_x] == wall {
			return true
		}
	}

	// down
	new_y = y + 1

	if 0 <= new_x && new_x < len(sMap[y]) &&
		0 <= new_y && new_y < len(sMap) {
		if sMap[new_y][new_x] == wall {
			return true
		}
	}
	// left
	new_x = x - 1
	new_y = y

	if 0 <= new_x && new_x < len(sMap[y]) &&
		0 <= new_y && new_y < len(sMap) {
		if sMap[new_y][new_x] == wall {
			return true
		}
	}
	// right
	new_x = x + 1

	if 0 <= new_x && new_x < len(sMap[y]) &&
		0 <= new_y && new_y < len(sMap) {
		if sMap[new_y][new_x] == wall {
			return true
		}
	}

	return false
}

func run_check(x int, y int, inputstr string, output chan bool) {

	guardMap := make(GMap, 0)
	guardMap = initMap(guardMap, inputstr)
	for _, guarddir := range guards {
		if guardMap[y][x] == guarddir {
			output <- false
			return
		}
	}
	guardMap[y][x] = wall
	// println(" ------ ")
	// printMap(guardMap)
	output <- check_loop(guardMap)
}

func check_loop(guardMap GMap) bool {
	// if you hit a wall from the same position twice you are in a loop
	touched := make(TouchMap)
	looped := false

	// Walk the map
	for x_pos, y_pos := find_guard(guardMap); x_pos != -1; x_pos, y_pos = find_guard(guardMap) {
		if guard_blocked(guardMap, x_pos, y_pos) {
			// Track current pos + dir as wall.
			guardMap, touched, looped = touched_wall(guardMap, x_pos, y_pos, touched)
			if looped {
				break
			}
		} else {
			guardMap = move_forward(guardMap, x_pos, y_pos)
		}
	}
	// println("")
	// printMap(guardMap)
	return looped
}

func touched_wall(guardMap GMap, x_pos int, y_pos int, touched TouchMap) (GMap, TouchMap, bool) {
	//track position
	looped := false
	guard := guardMap[y_pos][x_pos]
	x_dir, y_dir := get_direction(guard)

	pos_str := fmt.Sprintf("%d,%d", x_pos, y_pos)
	dir_str := fmt.Sprintf("%d,%d", x_dir, y_dir)

	dirmap, exists := touched[pos_str]
	if !exists {
		// We have not hit a wall at this position before
		dirmap = make(TouchDir)
		dirmap[dir_str] = true
		touched[pos_str] = dirmap
	} else {
		// we have previously hit a wall while standing at this position
		// Check which direction we have faced
		_, exists := dirmap[dir_str]
		if exists {
			// Facing the same direction as before, loop
			looped = true
		} else {
			// facing new direction
			dirmap[dir_str] = true
		}
	}

	guardMap = turn_guard(guardMap, x_pos, y_pos)

	return guardMap, touched, looped
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
