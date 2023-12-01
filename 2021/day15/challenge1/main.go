package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const FILE_LOCATION = "./input.txt"
const DEBUG = false

type point struct {
	visited bool
	value   int
	weight  int
}

func main() {
	grid := readFile()
	grid[0][0].weight = grid[0][0].value

	ended := false
	for !ended {
		x, y := pointToExplore(grid)

		fmt.Printf("Exploring %d %d\n", x, y)

		grid[y][x].visited = true
		grid = explorePoint(grid, x, y, x+1, y)
		grid = explorePoint(grid, x, y, x, y+1)
		grid = explorePoint(grid, x, y, x-1, y)
		grid = explorePoint(grid, x, y, x, y-1)

		ended = x == len(grid[0])-1 && y == len(grid)-1
	}

	fmt.Println(grid[len(grid)-1][len(grid[0])-1].weight - grid[0][0].value)

}

func readFile() [][]point {
	var grid [][]point

	file, err := os.Open(FILE_LOCATION)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line []point
		text := scanner.Text()

		for _, c := range text {
			number, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}

			line = append(line, point{
				visited: false,
				value:   number,
				weight:  int(^uint(0) >> 1),
			})
		}

		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid
}

func pointToExplore(grid [][]point) (int, int) {
	exploreX := 0
	exploreY := 0
	minimal := int(^uint(0) >> 1)

	for y := range grid {
		for x := range grid[y] {
			point := grid[y][x]

			if DEBUG {
				fmt.Printf("Is %d %d is the next point to explore (visited %t) ? weight %d minimal %d\n", x, y, point.visited, point.weight, minimal)
			}

			if !point.visited && point.weight < minimal {
				exploreX = x
				exploreY = y
				minimal = point.weight
			}
		}
	}

	return exploreX, exploreY
}

func explorePoint(grid [][]point, rootX int, rootY int, x int, y int) [][]point {
	if y < 0 || y > len(grid)-1 {
		return grid
	}

	if x < 0 || x > len(grid[0])-1 {
		return grid
	}

	if grid[y][x].visited {
		return grid
	}

	if DEBUG {
		fmt.Printf("Updating weight of %d %d with %d or %d\n", x, y, grid[rootY][rootX].weight+grid[y][x].value, grid[y][x].weight)
	}

	if grid[rootY][rootX].weight+grid[y][x].value < grid[y][x].weight {
		grid[y][x].weight = grid[rootY][rootX].weight + grid[y][x].value
	}

	return grid
}
