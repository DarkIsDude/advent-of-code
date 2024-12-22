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
			if input[i][j] == 'A' {
				if count2MASInX(input, i, j) {
					countOfXMAS++
				}
			}
		}
	}

	fmt.Println(countOfXMAS)
}

func count2MASInX(input []string, i, j int) bool {
	fmt.Println("Trying", i, j)
	count := 0

	if i-1 < 0 || i+1 >= len(input) || j-1 < 0 || j+1 >= len(input[i]) {
		fmt.Println("Out of bounds")
		return false
	}

	if input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S' {
		fmt.Println("Found 1")
		count++
	}

	if input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M' {
		fmt.Println("Found 2")
		count++
	}

	if input[i+1][j-1] == 'M' && input[i-1][j+1] == 'S' {
		fmt.Println("Found 3")
		count++
	}

	if input[i+1][j-1] == 'S' && input[i-1][j+1] == 'M' {
		fmt.Println("Found 4")
		count++
	}

	return count == 2
}
