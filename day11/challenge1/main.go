package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const NUMBER_OF_ITERATION = 100
const FILE_LOCATION = "./input.txt"

func main() {
	grid := readFile()

	fmt.Println("We found the following grid, let's play")
	display(grid)

	totalExplosion := play(grid)

	fmt.Printf("Total number of explosion %d\n", totalExplosion)
}

func readFile() [][]int {
	var grid [][]int

	file, err := os.Open(FILE_LOCATION)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		var line []int

		for _, c := range text {
			number, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}

			line = append(line, number)
		}

		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid
}

func display(grid [][]int) {
	for _, line := range grid {
		for _, number := range line {
			fmt.Print(number)
		}

		fmt.Println(" ")
	}
}

func play(grid [][]int) int {
	totalExplosion := 0

	for i := 0; i < NUMBER_OF_ITERATION; i++ {
		fmt.Printf("-------- %d --------\n", i+1)

		totalExplosion += iterate(grid)

		display(grid)
	}

	return totalExplosion
}

func iterate(grid [][]int) int {
	totalExplosion := 0

	// Every element growth
	for y, line := range grid {
		for x := range line {
			grid[y][x]++
		}
	}

	// Explode everyone who have 9 as value
	elementFound := true
	for elementFound {
		elementFound = false

		for y, line := range grid {
			for x, number := range line {
				if number > 9 {
					elementFound = true
					grid[y][x] = 0
					explode(grid, x, y)

					totalExplosion++
				}
			}
		}
	}

	return totalExplosion
}

func explode(grid [][]int, xRoot int, yRoot int) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			explodeX := x + xRoot
			isValidX := explodeX >= 0 && explodeX < len(grid[0])

			explodeY := y + yRoot
			isValidY := explodeY >= 0 && explodeY < len(grid)

			if isValidX && isValidY && grid[explodeY][explodeX] > 0 {
				grid[explodeY][explodeX]++
			}
		}
	}
}
