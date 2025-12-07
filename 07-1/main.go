package main

import (
	"bufio"
	"log"
	"os"
)

func calculate_line(line []byte, above []byte) ([]byte, int) {
	new_line := make([]byte, len(line))
	times_split := 0

	for i := range new_line {
		new_line[i] = '.'
	}

	for i := range line {
		switch above[i] {
		case 'S', '|':
			if line[i] == '^' {
				new_line[i] = '^'
				times_split += 1
			} else {
				new_line[i] = '|'
			}
		case '^':
			if i > 0 && new_line[i-1] != '^' {
				new_line[i-1] = '|'
			}
			if i < len(line)-1 && new_line[i+1] != '^' {
				new_line[i+1] = '|'
			}
		}
	}

	return new_line, times_split
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("bad filename: %s", err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Scan()

	grid := [][]byte{}
	grid = append(grid, append([]byte{}, scanner.Bytes()...))

	total_split := 0
	for scanner.Scan() {
		new_line, split := calculate_line(scanner.Bytes(), grid[len(grid)-1])
		grid = append(grid, new_line)
		total_split += split
	}

	for _, v := range grid {
		log.Println(string(v))
	}
	log.Println(total_split)
}
