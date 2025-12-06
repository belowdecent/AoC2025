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

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		batteries := scanner.Bytes()
		index1, joltage1 := find_largest(batteries[:len(batteries)-1])
		_, joltage2 := find_largest(batteries[index1+1:])

		total_joltage += joltage1*10 + joltage2
	}

	log.Printf("total joltage: %d", total_joltage)
}
