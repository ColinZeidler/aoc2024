package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()

	part2()

}

func part1() {
	println("Part1")

	input, err := os.ReadFile("day5-input.txt")
	if err != nil {
		panic("Error reading input")
	}
	inputStr := string(input)

	// map, page -> before
	// have a map of printed pages
	// if current page is in  map, check that its before is not in printed
	orderRules := make(map[int]map[int]bool, 0)
	updates := make([]string, 0)
	sectionEnd := false
	sum := 0
	for _, line := range strings.Split(inputStr, "\n") {
		if len(line) == 0 {
			sectionEnd = true
			continue
		}
		if sectionEnd {
			updates = append(updates, line)
		} else {
			lineData := strings.Split(line, "|")
			page, err := strconv.Atoi(lineData[0])
			if err != nil {
				panic("Error parsing int")
			}
			before, err := strconv.Atoi(lineData[1])
			if err != nil {
				panic("Error parsing int")
			}
			before_map, exists := orderRules[page]
			if !exists {
				before_map = make(map[int]bool)
				orderRules[page] = before_map
			}
			before_map[before] = true
		}
	}

	for _, update := range updates {
		pages := strings.Split(update, ",")
		printed := make(map[int]bool, 0)
		ok := true
		for _, page := range pages {
			ipage, parsed := strconv.Atoi(page)
			if parsed != nil {
				panic("Error parsing page number")
			}
			before_map, exists := orderRules[ipage]
			if exists {
				for before := range before_map {
					_, outoforder := printed[before]
					if outoforder {
						ok = false
						// fmt.Printf("Pages out of order: %s, %d\n", pages, ipage)
						break
					}
				}
			}
			if !ok {
				break
			}
			printed[ipage] = true
		}
		if !ok {
			continue
		}
		middle := len(pages) / 2
		val, parsed := strconv.Atoi(pages[middle])
		if parsed != nil {
			panic("Error parsing int")
		}
		sum += val
	}
	println(sum)

}

func part2() {
	println("Part2")

}
