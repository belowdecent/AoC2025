package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func rotate_left(pos int, amount int) (int, int) {
	norm_amount := amount % 100
	zero_count := amount / 100

	new_pos := pos - norm_amount
	if new_pos < 0 {
		new_pos += 100
		zero_count += 1
	}

	return new_pos, zero_count
}

func rotate_right(pos int, amount int) (int, int) {
	norm_amount := amount % 100
	zero_count := amount / 100

	if pos == 0 {
		zero_count += 1
	}

	new_pos := pos + norm_amount
	if new_pos > 100 {
		new_pos -= 100
		zero_count += 1
	}
	new_pos = new_pos % 100

	return new_pos, zero_count
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
		zero_per_turn := 0

		if turn, ok := strings.CutPrefix(turn, "L"); ok {
			turn_amount, err := strconv.Atoi(turn)
			if err != nil {
				log.Fatalf("bad turn: %s", turn)
			}
			dial_pos, zero_per_turn = rotate_left(dial_pos, turn_amount)
			zero_count += zero_per_turn
		} else if turn, ok := strings.CutPrefix(turn, "R"); ok {
			turn_amount, err := strconv.Atoi(turn)
			if err != nil {
				log.Fatalf("bad turn: %s", turn)
			}
			dial_pos, zero_per_turn = rotate_right(dial_pos, turn_amount)
			zero_count += zero_per_turn
		} else {
			log.Fatalf("illegal command: %s", turn)
		}
	}

	if dial_pos == 0 {
		zero_count += 1
	}

	log.Println(zero_count, dial_pos)
}
