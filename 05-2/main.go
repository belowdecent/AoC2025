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

func ranges_overlap(r1, r2 [2]int) bool {
	return r1[0] >= r2[0] && r1[0] <= r2[1] || r2[0] >= r1[0] && r2[0] <= r1[1]
}

func join_ranges(r1, r2 [2]int) [2]int {
	return [2]int{min(r1[0], r2[0]), max(r1[1], r2[1])}
}

func reduce_existing(ranges [][2]int, index int) (new_index int, new_ranges [][2]int) {
	for i, r := range ranges {
		if i == index {
			continue
		}

		if ranges_overlap(ranges[index], r) {
			new_ranges = append([][2]int{}, ranges[:index]...)
			new_ranges = append(new_ranges, ranges[index+1:]...)
			if i < index {
				new_ranges[i] = join_ranges(r, ranges[index])
				return i, new_ranges
			} else {
				new_ranges[i-1] = join_ranges(r, ranges[index])
				return i - 1, new_ranges
			}
		}
	}

	return -1, ranges
}

func add_range(ranges [][2]int, new_r [2]int) (new_index int, new_ranges [][2]int) {
	new_ranges = append([][2]int{}, ranges...)

	for i, r := range ranges {
		if ranges_overlap(new_r, r) {
			new_ranges[i] = join_ranges(new_r, r)
			return i, new_ranges
		}
	}

	return -1, append(new_ranges, new_r)
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("bad filename: %s", err)
	}

	ranges := [][2]int{}
	new_index := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			break
		}

		new_range := [2]int{}
		new_range[0], new_range[1] = get_range(scanner.Text())
		new_index, ranges = add_range(ranges, new_range)
		for new_index != -1 {
			new_index, ranges = reduce_existing(ranges, new_index)
		}
	}

	total_fresh := 0
	for _, r := range ranges {
		total_fresh += r[1] - r[0] + 1
	}
	log.Printf("total fresh: %d", total_fresh)
}
