package main

import (
	"bufio"
	"log"
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

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("bad filename: %s", err)
	}

	reading_ranges := true
	ranges := [][2]int{}
	fresh_count := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			reading_ranges = false
			continue
		} else if reading_ranges {
			new_range := [2]int{}
			new_range[0], new_range[1] = get_range(scanner.Text())
			ranges = append(ranges, new_range)
		} else {
			ingridient, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatalf("evil ingrindient: %s", err)
			}

			ingridient_fresh := false
			for _, r := range ranges {
				if ingridient >= r[0] && ingridient <= r[1] {
					ingridient_fresh = true
					break
				}
			}
			if ingridient_fresh {
				fresh_count += 1
			}
		}
	}

	log.Printf("Fresh score: %d", fresh_count)
}
