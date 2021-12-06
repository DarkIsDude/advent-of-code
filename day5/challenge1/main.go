package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Grid struct {
	grid  [][]int
	lines []Line
}

type Line struct {
	x1 int
	y1 int

	x2 int
	y2 int
}

func main() {
	grid := readFile()

	for _, x := range grid.grid {
		fmt.Println(x)
	}

	for _, line := range grid.lines {
		fmt.Printf("Apply line x1 %d y1 %d and x2 %d y2 %d\n", line.x1, line.y1, line.x2, line.y2)
		grid = applyLine(grid, line)
	}

	for _, x := range grid.grid {
		fmt.Println(x)
	}

	fmt.Printf("Dangerous area %d\n", countDangerousArea(grid))
}

func readFile() Grid {
	grid := Grid{}
	maxSize := 0

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		positions := strings.Split(scanner.Text(), " -> ")
		points1 := strings.Split(positions[0], ",")
		x1, err := strconv.Atoi(points1[0])
		if err != nil {
			panic(err)
		}

		y1, err := strconv.Atoi(points1[1])
		if err != nil {
			panic(err)
		}

		points2 := strings.Split(positions[1], ",")
		x2, err := strconv.Atoi(points2[0])
		if err != nil {
			panic(err)
		}

		y2, err := strconv.Atoi(points2[1])
		if err != nil {
			panic(err)
		}

		if x1 > maxSize {
			maxSize = x1
		}

		if y1 > maxSize {
			maxSize = y1
		}

		if x2 > maxSize {
			maxSize = x2
		}

		if y2 > maxSize {
			maxSize = y2
		}

		line := Line{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		}

		grid.lines = append(grid.lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for y := 0; y < maxSize+1; y++ {
		grid.grid = append(grid.grid, []int{})

		for x := 0; x < maxSize+1; x++ {
			grid.grid[y] = append(grid.grid[y], 0)
		}
	}

	return grid
}

func applyLine(grid Grid, line Line) Grid {
	if line.x1 != line.x2 && line.y1 != line.y2 {
		fmt.Println("Not good line")
		return grid
	}

	y := 0
	stopY := false
	lengthY := int(math.Abs(float64(line.y1 - line.y2)))
	startY := line.y1
	if line.y1 > line.y2 {
		startY = line.y2
	}

	for !stopY {
		x := 0
		stopX := false
		lengthX := int(math.Abs(float64(line.x1 - line.x2)))
		startX := line.x1
		if line.x1 > line.x2 {
			startX = line.x2
		}

		for !stopX {
			fmt.Printf("Add 1 to %d, %d\n", startX+x, startY+y)
			grid.grid[startY+y][startX+x]++

			if x >= lengthX {
				stopX = true
			}

			x++
		}

		if y >= lengthY {
			stopY = true
		}

		y++
	}

	return grid
}

func countDangerousArea(grid Grid) int {
	dangerousArea := 0

	for _, y := range grid.grid {
		for _, x := range y {
			if x >= 2 {
				dangerousArea++
			}
		}
	}

	return dangerousArea
}
