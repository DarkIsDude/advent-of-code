package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	countOfXMAS := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'X' {
				if hasXMASInDirection(input, i, j, -1, -1) {
					countOfXMAS++
				}

				if hasXMASInDirection(input, i, j, 0, -1) {
					countOfXMAS++
				}

				if hasXMASInDirection(input, i, j, 1, -1) {
					countOfXMAS++
				}

				if hasXMASInDirection(input, i, j, 1, 0) {
					countOfXMAS++
				}

				if hasXMASInDirection(input, i, j, 1, 1) {
					countOfXMAS++
				}

				if hasXMASInDirection(input, i, j, 0, 1) {
					countOfXMAS++
				}

				if hasXMASInDirection(input, i, j, -1, 1) {
					countOfXMAS++
				}

				if hasXMASInDirection(input, i, j, -1, 0) {
					countOfXMAS++
				}
			}
		}
	}

	fmt.Println(countOfXMAS)
}

func hasXMASInDirection(input []string, i, j, x, y int) bool {
	if i+x*3 >= len(input) || j+y*3 >= len(input[i]) {
		return false
	}

	if i+x*3 < 0 || j+y*3 < 0 {
		return false
	}

	if input[i][j] == 'X' {
		if input[i+x][j+y] == 'M' {
			if input[i+x*2][j+y*2] == 'A' {
				if input[i+x*3][j+y*3] == 'S' {
					return true
				}
			}
		}
	}

	return false
}
