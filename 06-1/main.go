package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func apply_op(page [][]int, index int, op byte) int {
	switch op {
	case '+':
		result := 0
		for row := range len(page) {
			result += page[row][index]
		}
		return result
	case '*':
		result := 1
		for row := range len(page) {
			result *= page[row][index]
		}
		return result
	}

	log.Fatalln("illegal op:", op)
	return -1
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("bad filename: %s", err)
	}

	page := [][]int{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if scanner.Bytes()[0] == '+' || scanner.Bytes()[0] == '*' {
			break
		}

		numbers := []int{}
		for token := range strings.SplitSeq(scanner.Text(), " ") {
			if len(token) != 0 {
				n, err := strconv.Atoi(token)
				if err != nil {
					log.Fatalf("bad number %s: %s", token, err)
				}
				numbers = append(numbers, n)
			}
		}
		page = append(page, numbers)
	}

	index := 0
	total := 0
	for _, b := range scanner.Bytes() {
		if b == '+' || b == '*' {
			total += apply_op(page, index, b)
			index += 1
		}

		if index >= len(page[0]) {
			break
		}
	}

	log.Printf("total: %d", total)
}
