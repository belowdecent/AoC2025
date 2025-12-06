package main

import (
	"bufio"
	"log"
	"os"
)

func is_blocked(line []byte, i int) bool {
	if i < 0 || i >= len(line) {
		return false
	}

	return line[i] == '@'
}

type Field = [][]byte

func remove_pickable_from_line(lines [3][]byte) int {
	removed := 0

	for i, v := range lines[1] {
		if v == '@' {
			obstacle_count := 0
			pickable := true

		outer:
			for offset := -1; offset < 2; offset += 1 {
				for line_n, line := range lines {
					if line_n == 1 && offset == 0 {
						continue
					}

					if is_blocked(line, i+offset) {
						obstacle_count += 1
					}
					if obstacle_count >= 4 {
						pickable = false
						break outer
					}
				}
			}

			if pickable {
				lines[1][i] = '.'
				removed += 1
			}
		}
	}

	return removed
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("bad filename: %s", err)
	}

	scanner := bufio.NewScanner(input)
	field := [][]byte{}

	for scanner.Scan() {
		field = append(field, append([]byte(nil), scanner.Bytes()...))
	}

	total_removed := 0
	cycle_start := -1
	row := 0
	empty_row := make([]byte, len(field[row]))

	for cycle_start != row {
		removed := 0

		if row < 1 {
			removed = remove_pickable_from_line([3][]byte{empty_row, field[row], field[row+1]})
		} else if row > len(field)-2 {
			removed = remove_pickable_from_line([3][]byte{field[row-1], field[row], empty_row})
		} else {
			removed = remove_pickable_from_line([3][]byte{field[row-1], field[row], field[row+1]})
		}

		if removed > 0 {
			cycle_start = row
			total_removed += removed
		}
		row = (row + 1) % len(field)
	}

	log.Printf("total removed: %d", total_removed)
}
