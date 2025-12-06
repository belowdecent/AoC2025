package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func rotate(pos int, amount int) int {
	new_pos := pos + amount
	for new_pos < 0 {
		new_pos += 100
	}
	return new_pos % 100
}

func main() {
	dial_pos := 50
	zero_count := 0

	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("bad filename: %s", err)
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		turn := scanner.Text()

		if turn, ok := strings.CutPrefix(turn, "L"); ok {
			turn_amount, err := strconv.Atoi(turn)
			if err != nil {
				log.Fatalf("bad turn: %s", turn)
			}
			dial_pos = rotate(dial_pos, -turn_amount)
		} else if turn, ok := strings.CutPrefix(turn, "R"); ok {
			turn_amount, err := strconv.Atoi(turn)
			if err != nil {
				log.Fatalf("bad turn: %s", turn)
			}
			dial_pos = rotate(dial_pos, turn_amount)
		} else {
			log.Fatalf("illegal command: %s", turn)
		}

		if dial_pos == 0 {
			zero_count += 1
		}
	}

	log.Println(zero_count)
}
