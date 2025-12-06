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

func count_pickable(lines [3][]byte) int {
	pickable := len(lines[1])

	for i, v := range lines[1] {
		if v == '@' {
			obstacle_count := 0

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
						pickable -= 1
						break outer
					}
				}
			}
		} else {
			pickable -= 1
		}
	}

	return pickable
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("bad filename: %s", err)
	}

	lines := [3][]byte{}
	total_pickable := 0

	scanner := bufio.NewScanner(input)

	scanner.Scan()
	lines[1] = append([]byte(nil), scanner.Bytes()...)
	lines[0] = make([]byte, len(lines[1]))

	for scanner.Scan() {
		lines[2] = append([]byte(nil), scanner.Bytes()...)
		total_pickable += count_pickable(lines)

		lines[0] = lines[1]
		lines[1] = lines[2]
	}

	lines[2] = make([]byte, len(lines[1]))
	total_pickable += count_pickable(lines)

	log.Printf("total pickables: %d", total_pickable)
}
