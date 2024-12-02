package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

type Lists struct {
	left  []int
	right []int
}

func main() {

	//Read input into string
	input, err := os.ReadFile("day1-input.txt")
	if err != nil {
		panic("Failed to read input")
	}
	input_str := string(input)
	lists := parseData(input_str)
	// two lists of equal length

	//sort both lists, small to big
	sort.Slice(lists.left, func(i, j int) bool {
		return lists.left[i] < lists.left[j]
	})

	sort.Slice(lists.right, func(i, j int) bool {
		return lists.right[i] < lists.right[j]
	})

	total_distance := 0
	// iterate for len of lists, diff values at current point, add to running total
	for index := range lists.left {
		diff := lists.left[index] - lists.right[index]
		if diff < 0 {
			diff = diff * -1
		}
		total_distance += diff
	}

	// create map of id counts for right
	sim_score := 0
	rightMap := make(map[int]int)
	for _, item := range lists.right {
		count, ok := rightMap[item]
		if !ok {
			rightMap[item] = 1
		} else {
			rightMap[item] = count + 1
		}
	}

	for _, item := range lists.left {
		count, ok := rightMap[item]
		if !ok {
			continue
		}
		sim_score += count * item
	}

	// return running total
	println("Difference")
	println(total_distance)

	println("Similarity")
	println(sim_score)

}

func parseData(input string) Lists {
	// Count number of lines
	lines := strings.Split(input, "\n")
	length := len(lines)
	list := Lists{
		make([]int, 0, length),
		make([]int, 0, length),
	}

	// parse lines from input string
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		values := strings.Split(line, " ")
		left, err := strconv.Atoi(values[0])
		if err != nil {
			panic("Error parsing Int")
		}
		right, err := strconv.Atoi(values[3])
		if err != nil {
			panic("Error parsing Int")
		}
		list.left = append(list.left, left)
		list.right = append(list.right, right)
	}

	return list
}
