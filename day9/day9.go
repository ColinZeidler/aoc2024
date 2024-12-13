package main

import (
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	println("Part1")
	// input, ok := os.ReadFile("small.txt")
	input, ok := os.ReadFile("day9-input.txt")
	if ok != nil {
		panic("Error reading input")
	}
	inputstr := string(input)

	next_id := 0
	is_file := true
	blocklist := make([]Block, 0, 100)
	for _, blocktile := range inputstr {
		if blocktile == '\n' {
			break
		}
		block_len := parseint(blocktile)
		for range block_len {
			id := next_id
			if !is_file {
				id = -1
			}
			newBlock := Block{
				file: is_file,
				id:   id,
			}
			blocklist = append(blocklist, newBlock)
		}
		if is_file {
			next_id += 1
		}
		is_file = !is_file
	}

	right_pos := len(blocklist) - 1
leftloop:
	for left_pos, block := range blocklist {
		if block.file {
			continue
		}

		for !blocklist[right_pos].file {
			right_pos = right_pos - 1
			if right_pos < 0 {
				break leftloop
			}
		}
		if !blocklist[right_pos].file {
			break leftloop
		}
		if right_pos < left_pos {
			break
		}
		temp := blocklist[right_pos]
		blocklist[right_pos] = blocklist[left_pos]
		blocklist[left_pos] = temp

		// printblocks(blocklist)
	}
	total := checksum(blocklist)
	println(total)
}

func part2() {
	println("Part2")
	input, ok := os.ReadFile("small.txt")
	// input, ok := os.ReadFile("day9-input.txt")
	if ok != nil {
		panic("Error reading input")
	}
	inputstr := string(input)

	next_id := 0
	is_file := true
	blocklist := make([]Block, 0, 10)
	for _, blocktile := range inputstr {
		if blocktile == '\n' {
			break
		}
		block_len := parseint(blocktile)
		id := next_id
		if !is_file {
			id = -1
		}
		newBlock := Block{
			file:   is_file,
			id:     id,
			length: block_len,
		}
		blocklist = append(blocklist, newBlock)
		if is_file {
			next_id += 1
		}
		is_file = !is_file
	}
	printblocks(blocklist)

	last_id := next_id - 1
	for right_pos := len(blocklist) - 1; right_pos >= 0; right_pos-- {
		move_block := blocklist[right_pos]
		if !move_block.file {
			continue
		}

		if last_id > move_block.id {
			break
		}

		for left_pos, freespace := range blocklist {
			if freespace.file {
				continue
			}
			if freespace.length >= move_block.length {
				temp := blocklist[left_pos]
				blocklist[left_pos] = blocklist[right_pos]
				blocklist[right_pos] = temp
				break
			}
		}
		last_id -= 1
	}
	printblocks(blocklist)
}

type Block struct {
	file   bool
	id     int
	length int
}

func parseint(num rune) int {
	n, ok := strconv.Atoi(string(num))
	if ok != nil {
		panic("Error parsing int")
	}
	return n
}

func printblocks(blocklist []Block) {
	for _, b := range blocklist {
		for range b.length {
			if b.file {
				print(b.id)
			} else {
				print(".")
			}
		}
	}
	println()
}

func checksum(blocklist []Block) int {

	total := 0
	for index, block := range blocklist {
		if !block.file {
			break
		}
		total = total + (index * block.id)
	}
	return total
}
