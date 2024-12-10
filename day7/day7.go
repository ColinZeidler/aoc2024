package main

import (
	"fmt"
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
	input, ok := os.ReadFile("small.txt")
	// input, ok := os.ReadFile("day7-input.txt")
	if ok != nil {
		panic("Error reading input")
	}
	inputstr := string(input)
	lines := strings.Split(inputstr, "\n")

	output := make(chan int)

	vals := make([]int, 2)
	vals[0] = 0
	vals[1] = 1

	started := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		started += 1
		go handle_row(line, output, vals)
	}

	total := 0
	for range started {
		result := <-output
		total += result
	}
	println(total)
}

func part2() {
	println("Part2")
	// input, ok := os.ReadFile("small.txt")
	input, ok := os.ReadFile("day7-input.txt")
	if ok != nil {
		panic("Error reading input")
	}
	inputstr := string(input)
	lines := strings.Split(inputstr, "\n")

	output := make(chan int)

	vals := make([]int, 3)
	vals[0] = 0
	vals[1] = 1
	vals[2] = 2

	started := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		started += 1
		go handle_row(line, output, vals)
	}

	total := 0
	for range started {
		result := <-output
		total += result
	}
	println(total)
}

func handle_row(row string, output chan int, vals []int) {
	data := strings.Split(row, ": ")
	answer := parseint(data[0])
	numbers := data[1]
	numstr_list := strings.Split(numbers, " ")

	number_list := make([]int, len(numstr_list))

	for index, numstr := range numstr_list {
		number := parseint(numstr)
		number_list[index] = number
	}

	// 3 for example
	length := len(number_list) - 1

	idx := make([]int, length)

	perm := make([]int, length)
	for index := range perm {
		perm[index] = vals[0]
	}

	for {
		// fmt.Printf("perm %d\n", perm)
		// test value
		test_result := number_list[0]
		for index := range perm {
			left := test_result
			right := number_list[index+1]
			if perm[index] == 0 {
				test_result = left + right
			} else if perm[index] == 1 {
				test_result = left * right
			} else {
				test_result = parseint(fmt.Sprintf("%d%d", left, right))
			}
		}
		if test_result == answer {
			output <- answer
			return
		}

		k := len(idx) - 1
		for ; k >= 0; k-- {
			idx[k] += 1
			if idx[k] < len(vals) {
				perm[k] = vals[idx[k]]
				break
			}
			idx[k] = 0
			perm[k] = vals[idx[k]]
		}
		if k < 0 {
			break
		}
	}

	output <- 0
}

func parseint(num string) int {
	n, ok := strconv.Atoi(num)
	if ok != nil {
		panic("Error parsing int")
	}
	return n
}
