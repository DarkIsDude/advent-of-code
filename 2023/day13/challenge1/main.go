package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	args := os.Args
	file, err := os.Open(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	pattern := []string{}
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if line == "" {
			sum += findReflexion(pattern)
			pattern = []string{}
			continue
		}

		pattern = append(pattern, line)
	}

	sum += findReflexion(pattern)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func findReflexion(pattern []string) int {
	byColumn := findReflexionByColumn(pattern)

	if byColumn != -1 {
		fmt.Printf("Found by column: %d\n", byColumn)
		return byColumn
	}

	byWRow := findReflexionByRow(pattern)

	if byWRow != -1 {
		fmt.Printf("Found by row: %d\n", byWRow)
		return byWRow * 100
	}

	log.Panic("No reflexion found")
	return -1
}

func findReflexionByRow(pattern []string) int {
	for i := 1; i < len(pattern); i++ {
		numberOfRows := int(math.Min(float64(i), float64(len(pattern) - i)))

		rowMatch := true
		fmt.Printf("Checking row %d with %d rows\n", i, numberOfRows)

		for a := 0; a < numberOfRows; a++ {
			for j := 0; j < len(pattern[0]); j++ {
				if pattern[i - a - 1][j] != pattern[i + a][j] {
					rowMatch = false
				}
			}
		}

		if rowMatch {
			return i
		}
	}

	return -1
}

func findReflexionByColumn(pattern []string) int {
	for j := 1; j < len(pattern[0]); j++ {
		numberOfColumns := int(math.Min(float64(j), float64(len(pattern[0]) - j)))

		columnMatch := true
		fmt.Printf("Checking column %d with %d columns\n", j, numberOfColumns)

		for a := 0; a < numberOfColumns; a++ {
			for i := 0; i < len(pattern); i++ {
				if pattern[i][j - a - 1] != pattern[i][j + a] {
					columnMatch = false
				}
			}
		}

		if columnMatch {
			return j
		}
	}

	return -1
}
