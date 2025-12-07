package main

import (
	"bufio"
	"log"
	"os"
)

type Node struct {
	left        *Node
	right       *Node
	total_paths int
}

func (node *Node) get_total_paths() int {
	if node.total_paths != -1 {
		return node.total_paths
	}

	left_paths, right_paths := 0, 0
	if node.left != nil {
		left_paths = node.left.get_total_paths()
	}
	if node.right != nil {
		right_paths = node.right.get_total_paths()
	}

	node.total_paths = left_paths + right_paths
	return node.total_paths
}

func calculate_line(line []byte, above []byte) ([]byte, int) {
	new_line := make([]byte, len(line))
	times_split := 0

	for i := range new_line {
		new_line[i] = '.'
	}

	for i := range line {
		switch above[i] {
		case 'S', '|':
			if line[i] == '^' {
				new_line[i] = '^'
				times_split += 1
			} else {
				new_line[i] = '|'
			}
		case '^':
			if i > 0 && new_line[i-1] != '^' {
				new_line[i-1] = '|'
			}
			if i < len(line)-1 && new_line[i+1] != '^' {
				new_line[i+1] = '|'
			}
		}
	}

	return new_line, times_split
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("bad filename: %s", err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Scan()

	grid := [][]byte{}
	grid = append(grid, append([]byte{}, scanner.Bytes()...))

	total_split := 0
	for scanner.Scan() {
		new_line, split := calculate_line(scanner.Bytes(), grid[len(grid)-1])
		grid = append(grid, new_line)
		total_split += split
	}

	node_grid := make([][]*Node, len(grid))
	for i := range node_grid {
		node_grid[i] = make([]*Node, len(grid[0]))
	}

	for i, v := range grid[len(grid)-1] {
		if v == '|' {
			node_grid[len(grid)-1][i] = &Node{
				left:        nil,
				right:       nil,
				total_paths: 1,
			}
		}
	}

	for row := len(grid) - 2; row > 1; row -= 1 { // we ignore the first two lines
		for col := range grid[row] {
			switch grid[row][col] {
			case '|':
				node_grid[row][col] = &Node{
					left:        node_grid[row+1][col],
					right:       nil,
					total_paths: -1,
				}
			case '^':
				node_grid[row][col] = &Node{
					left:        node_grid[row+1][col-1],
					right:       node_grid[row+1][col+1],
					total_paths: -1,
				}

				if row == 2 {
					log.Printf("total paths: %d", node_grid[row][col].get_total_paths())
					return
				}
			}
		}
	}
}
