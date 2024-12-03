package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	levels      []int
	safe        bool
	error_count int
	bad_index   int
}

func main() {
	// input, err := os.ReadFile("day2-input.txt")
	input, err := os.ReadFile("test_data.txt")
	if err != nil {
		panic("Failed to read input")
	}
	input_str := string(input)
	lines := strings.Split(input_str, "\n")
	reports := make([]Report, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		new_report := parse_report(line)
		reports = append(reports, new_report)
	}

	safe_count := 0
	for _, report := range reports {
		report.CheckSafe(false)
		if report.safe {
			safe_count += 1
		} else {
			if report.error_count == 1 {
				report.safe = true
				report.CheckSafe(true)
				if report.safe {
					safe_count += 1
				}
			}
			// report.Print()
		}
	}
	println("Safe reports")
	println(safe_count)
}

func parse_report(report_data string) Report {
	report_array := strings.Split(report_data, " ")

	new_report := Report{
		levels:      make([]int, 0, len(report_array)),
		safe:        true,
		error_count: 0,
		bad_index:   0,
	}

	for _, item := range report_array {
		value, err := strconv.Atoi(item)
		if err != nil {
			println(item)
			panic("Error parsing int")
		}
		new_report.levels = append(new_report.levels, value)
	}

	return new_report
}

func (report *Report) CheckSafe(skip_bad bool) {
	previous_value := 0
	direction_up := false
	for index, value := range report.levels {
		if skip_bad && index == report.bad_index {
			continue
		}
		if index == 0 {
			previous_value = value
			continue
		}

		// get diff as absolute value
		diff := previous_value - value
		if diff < 0 {
			diff = diff * -1
		}

		// Change by 1 or more, and 3 or less
		if diff < 1 || diff > 3 {
			// println("Change out of bounds", previous_value, value, diff)
			report.safe = false
			report.error_count += 1
		}

		if index == 1 {
			// don't check direction change for unsafe between first and second as there is no direction initially
			direction_up = value > previous_value
			previous_value = value
			continue
		}

		current_direction_up := value > previous_value
		if current_direction_up != direction_up {
			// println("Direction changed", previous_value, value, direction_up)
			report.safe = false
			report.error_count += 1
		}

		previous_value = value
	}
}

func (report *Report) Print() {
	for _, value := range report.levels {
		fmt.Printf(" %d", value)
	}
	fmt.Print("\n")
}
