package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func apply_op(line []int, op byte) int {
	switch op {
	case '+':
		result := 0
		for _, n := range line {
			result += n
		}
		return result
	case '*':
		result := 1
		for _, n := range line {
			result *= n
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

	// reading the input and finding the ops
	scanner := bufio.NewScanner(input)
	input_buffer := bytes.Buffer{} // we need to read the last line first hehe
	for scanner.Scan() {
		if scanner.Bytes()[0] == '+' || scanner.Bytes()[0] == '*' {
			break
		}
		fmt.Fprintln(&input_buffer, scanner.Text())
	}

	// calculating column widths
	column_widths := []int{}
	ops := []byte{}
	for _, b := range scanner.Bytes() {
		if b == '+' || b == '*' {
			ops = append(ops, b)
			column_widths = append(column_widths, 0) // subtract one to account for spacing
		} else {
			column_widths[len(column_widths)-1] += 1
		}
	}
	column_widths[len(column_widths)-1] += 1 // no white space at the end here

	// reading written numbers normally
	entries := [][]string{}
	scanner = bufio.NewScanner(&input_buffer)
	for scanner.Scan() {
		line := []string{}
		from := 0
		for i, width := range column_widths {
			line = append(line, scanner.Text()[from:from+width])
			if i < len(column_widths)-1 {
				from += 1
			}
			from += width
		}
		entries = append(entries, line)
	}

	// transposing kiln from dark souls 3
	transposed := make([][]string, len(entries[0]))
	for i := range len(entries[0]) {
		transposed[i] = make([]string, len(entries))
	}
	for i := range len(entries) {
		for j := range len(entries[i]) {
			transposed[j][i] = entries[i][j]
		}
	}

	// now the real game begins
	numbers := [][]int{}
	for col_num, col := range transposed {
		numbers = append(numbers, []int{})
		for digit := range column_widths[col_num] {
			new_n := []byte{}
			for _, entry := range col {
				new_n = append(new_n, entry[digit])
			}

			new_n_int, err := strconv.Atoi(strings.TrimSpace(string(new_n)))
			if err != nil {
				log.Fatalf("oh fuck: '%s'", new_n)
			}
			numbers[len(numbers)-1] = append(numbers[len(numbers)-1], new_n_int)
		}
	}

	total := 0
	for i, line := range numbers {
		total += apply_op(line, ops[i])
	}
	log.Printf("total: %d", total)
}
