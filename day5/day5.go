package main

import (
	"os"
	"slices"
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
	sum := 0
	orderRules, updates = init_data(orderRules, updates, inputStr)

	for _, update := range updates {
		pages := strings.Split(update, ",")
		printed := make(map[int]bool, 0)
		ok := true
		for _, page := range pages {
			ipage := parseint(page)
			ok = check_page_order(ipage, orderRules, printed)
			if !ok {
				break
			}
			printed[ipage] = true
		}
		if !ok {
			continue
		}
		middle := len(pages) / 2
		val := parseint(pages[middle])
		sum += val
	}
	println(sum)

}

func part2() {
	println("Part2")
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
	sum := 0
	orderRules, updates = init_data(orderRules, updates, inputStr)

	for _, update := range updates {
		pages := strings.Split(update, ",")
		printed := make(map[int]bool, 0)
		tosort := make([]int, 0, len(pages))
		ok := true
		for _, page := range pages {
			ipage := parseint(page)
			tosort = append(tosort, ipage)
			if ok {
				ok = check_page_order(ipage, orderRules, printed)
			}
			printed[ipage] = true
		}

		// get the sum of only updates that were bad
		if ok {
			continue
		}
		sorted := make([]int, 0, len(pages))
		// prev_length := len(tosort)
		for len(tosort) > 0 {
			for index, page := range tosort {
				rules, exists := orderRules[page]
				// if page not in rules, append to sorted
				satisfied := true
				if !exists {
					sorted = append(sorted, page)
				} else {
					// if page has rules check rule exists and insert before rule entry.
					// check all before pages exist
					// only check rules that are part of the page list
					for rule := range rules {
						if rule_in_pages(rule, pages) {
							found := rule_in_output(rule, sorted)
							if !found {
								satisfied = false
								break
							}
						}
					}
					if satisfied {
						sorted = slices.Insert(sorted, 0, page)
					}
				}

				// remove sorted from tosort
				if satisfied {
					tosort = remove(tosort, index)
					break
				}
			}
		}

		middle := len(sorted) / 2
		val := sorted[middle]
		sum += val
	}
	println(sum)

}

func init_data(rules map[int]map[int]bool, updates []string, inputStr string) (map[int]map[int]bool, []string) {
	sectionEnd := false
	for _, line := range strings.Split(inputStr, "\n") {
		if len(line) == 0 {
			sectionEnd = true
			continue
		}
		if sectionEnd {
			updates = append(updates, line)
		} else {
			lineData := strings.Split(line, "|")
			page := parseint(lineData[0])
			before := parseint(lineData[1])
			before_map, exists := rules[page]
			if !exists {
				before_map = make(map[int]bool)
				rules[page] = before_map
			}
			before_map[before] = true
		}
	}

	return rules, updates
}

func check_page_order(page int, rules map[int]map[int]bool, printed map[int]bool) bool {
	ok := true
	before_map, exists := rules[page]
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
	return ok
}

func parseint(num string) int {
	value, ok := strconv.Atoi(num)
	if ok != nil {
		panic("Error parsing Int")
	}
	return value
}

func remove(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func rule_in_pages(rule int, pages []string) bool {
	found := false
	for _, page := range pages {
		ipage := parseint(page)
		if rule == ipage {
			found = true
			break
		}
	}
	return found
}

func rule_in_output(rule int, output []int) bool {
	found := false
	for _, sortedPage := range output {
		if rule == sortedPage {
			found = true
		}
	}
	return found
}
