package main

import (
	"os"
	"strings"
)

var search_pattern string = "XMAS"

func main() {

	input, err := os.ReadFile("day4-input.txt")
	if err != nil {
		panic("Can't read input")
	}

	puzzle := string(input)

	xmas_count := 0

	lines := make([]string, 0)
	for _, line := range strings.Split(puzzle, "\n") {
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}

	for puzzle_row, line := range lines {
		for puzzle_column, letter := range line {
			if letter != rune(search_pattern[0]) {
				continue
			}
			for x_dir := -1; x_dir <= 1; x_dir++ {
				for y_dir := -1; y_dir <= 1; y_dir++ {
					if x_dir == 0 && y_dir == 0 {
						continue
					}
					xmas_count += checkDir(y_dir, x_dir, puzzle_row, puzzle_column, lines)
				}
			}

		}
	}

	println(xmas_count)

}

func checkDir(row_direction int, column_direction int, row_pos int, column_pos int, puzzle []string) int {
	xmas_count := 0
	found := true
	for shift, xmas := range search_pattern {
		// up
		new_row := row_pos + (shift * row_direction)
		new_col := column_pos + (shift * column_direction)
		if new_row < 0 || new_col < 0 {
			found = false
			break
		}
		if new_row >= len(puzzle) || new_col >= len(puzzle[new_row]) {
			found = false
			break
		}
		if rune(puzzle[new_row][new_col]) != xmas {
			found = false
			break
		}
	}
	if found {
		xmas_count += 1
	}
	return xmas_count
}
