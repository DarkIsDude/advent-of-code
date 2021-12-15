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
	grid = reproduceGridBy5(grid)
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
				value:   number,
				visited: false,
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

func reproduceGridBy5(grid [][]point) [][]point {
	originalXSize := len(grid[0])
	originalYSize := len(grid)

	for y := 0; y < originalYSize*5; y++ {
		if y >= originalYSize {
			grid = append(grid, []point{})
		}

		for x := 0; x < originalXSize*5; x++ {
			if y < originalYSize && x < originalXSize {
				continue
			}

			if DEBUG {
				fmt.Printf("Adding a new point %d %d\n", x, y)
			}

			grid[y] = append(grid[y], point{
				value:   0,
				visited: false,
				weight:  int(^uint(0) >> 1),
			})
		}
	}

	if DEBUG {
		for _, line := range grid {
			for _, point := range line {
				fmt.Print(point.value)
			}

			fmt.Println(" ")
		}
	}

	for i := 0; i <= 4; i++ {
		for j := 0; j <= 4; j++ {

			for y := 0; y < originalYSize; y++ {
				for x := 0; x < originalXSize; x++ {
					newY := originalYSize*j + y
					newX := originalXSize*i + x

					if newY < originalYSize && newX < originalXSize {
						continue
					}

					originalValue := grid[y][x].value
					newValue := (originalValue + i + j)

					if newValue > 9 {
						newValue = newValue%10 + 1
					}

					if DEBUG {
						fmt.Printf("New value for %d*%d+%d %d*%d+%d %d\n", originalXSize, i, x, originalYSize, j, y, newValue)
					}

					grid[newY][newX].value = newValue
				}
			}
		}
	}

	if DEBUG {
		for _, line := range grid {
			for _, point := range line {
				fmt.Printf("%d", point.value)
			}

			fmt.Println(" ")
		}
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
