package main

import (
	"bufio"
	"log"
	"os"
)

func find_largest(batteries []byte) (index int, value int) {
	max_value := -1
	max_index := -1
	for i, v := range batteries {
		if int(v) > max_value {
			max_index = i
			max_value = int(v)
		}

		if v == '9' {
			break
		}
	}

	max_value -= int('0')
	return max_index, max_value
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("bad filename: %s", err)
	}

	total_joltage := 0
	battery_count := 12

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		batteries := scanner.Bytes()
		bank_joltage := 0
		index := 0

		for i := range battery_count {
			new_index, joltage := find_largest(batteries[index : len(batteries)-(battery_count-1-i)])
			index += 1 + new_index
			bank_joltage = bank_joltage*10 + joltage
		}

		total_joltage += bank_joltage
	}

	log.Printf("total joltage: %d", total_joltage)
}
