package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func factor(number int) []int {
	factors := make([]int, 0)

	for i := 1; i < number; i += 1 {
		if number%i == 0 {
			factors = append(factors, i)
		}
	}

	return factors
}

func check_valid_sized(number int, size int) bool {
	digits_mask := int(math.Pow10(size))
	pattern := number % digits_mask
	number = number / digits_mask

	for number > 0 {
		if number%digits_mask != pattern {
			return true
		}
		number = number / digits_mask
	}

	return false
}

func get_range(str string) (int, int) {
	ranges := strings.Split(str, "-")
	start, err := strconv.Atoi(ranges[0])
	if err != nil {
		log.Fatalf("wrong range (start): %s, %s", str, err)
	}
	end, err := strconv.Atoi(ranges[1])
	if err != nil {
		log.Fatalf("wrong range (end): %s, %s", str, err)
	}

	return start, end
}

func check_valid(number int) bool {
	id_length := int(math.Floor(math.Log10(float64(number)))) + 1

	for _, factor := range factor(id_length) {
		if !check_valid_sized(number, factor) {
			return false
		}
	}
	return true
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("bad filename: %s", err)
	}
	str := string(b)
	ranges := strings.Split(str, ",")
	invalid_id_sum := 0

	for _, token := range ranges {
		token = strings.TrimSpace(token)
		// log.Printf("Checking %s", token)
		start, end := get_range(token)
		for i := start; i <= end; i++ {
			if !check_valid(i) {
				invalid_id_sum += i
			}
		}
	}
	log.Printf("id sum: %d", invalid_id_sum)
}
