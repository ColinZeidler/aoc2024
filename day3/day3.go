package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("day3-input.txt")
	// input, err := os.ReadFile("test_data.txt")
	if err != nil {
		panic("Error reading input")
	}

	input_str := string(input)

	spliced_str := input_str[:]
	println(spliced_str)
	sum := 0

	enabled := true
	for len(spliced_str) > 0 {

		// Check for disable enable items
		do_index := strings.Index(spliced_str, "do()")
		if do_index == -1 {
			do_index = len(spliced_str)
		}
		dont_index := strings.Index(spliced_str, "don't()")
		if dont_index == -1 {
			dont_index = len(spliced_str)
		}

		// find the next mul start
		index := strings.Index(spliced_str, "mul(")

		if do_index < index && do_index < dont_index {
			spliced_str = spliced_str[do_index+4:]
			enabled = true
			continue
		}

		if dont_index < index && dont_index < do_index {
			spliced_str = spliced_str[dont_index+7:]
			enabled = false
			continue
		}

		if index == -1 {
			break
		}
		spliced_str = spliced_str[index:]

		end := strings.Index(spliced_str, ")")
		if end == -1 {
			break
		}

		mul_string := spliced_str[:end+1]
		if is_valid(mul_string) {
			if enabled {
				sum += mul_values(mul_string)
			}
		}
		// TODO run mul calc
		spliced_str = spliced_str[4:]
	}
	println(sum)
}

func is_valid(substr string) bool {
	matched, err := regexp.Match(`^mul\(\d+,\d+\)`, []byte(substr))
	if err != nil {
		panic("Error running regexp")
	}
	return matched
}

func mul_values(substr string) int {
	left := strings.Index(substr, "(")
	right := strings.Index(substr, ")")
	sep := strings.Index(substr, ",")
	lSide, err := strconv.Atoi(substr[left+1 : sep])
	if err != nil {
		return 0
	}

	rSide, err := strconv.Atoi(substr[sep+1 : right])
	if err != nil {
		return 0
	}
	return lSide * rSide
}
