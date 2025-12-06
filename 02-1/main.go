package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

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
	if id_length%2 == 1 {
		return true
	}
	digits_mask := int(math.Pow10(id_length / 2))
	lower := number % digits_mask
	upper := number / digits_mask

	return lower != upper
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
