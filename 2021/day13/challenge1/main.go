package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const FILE_LOCATION = "./input.txt"

func main() {
	grid, instructions := readFile()

	display(grid)

	for pos, instruction := range instructions {
		fmt.Printf("Applying instruction %d : %s\n", pos, instruction)
		grid = applyInstruction(grid, instruction)
		fmt.Printf("Couting %d dots\n", countDots(grid))
	}
}

func readFile() ([][]int, []string) {
	var grid [][]int
	var instructions []string

	maxX := 0
	maxY := 0
	var pointX []int
	var pointY []int

	file, err := os.Open(FILE_LOCATION)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		}

		if strings.HasPrefix(text, "fold") {
			instructions = append(instructions, text)
			continue
		}

		points := strings.Split(text, ",")
		x, err := strconv.Atoi(points[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(points[1])
		if err != nil {
			panic(err)
		}

		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}

		pointX = append(pointX, x)
		pointY = append(pointY, y)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for y := 0; y <= maxY; y++ {
		grid = append(grid, []int{})

		for x := 0; x <= maxX; x++ {
			grid[y] = append(grid[y], 0)
		}
	}

	for pos := range pointX {
		grid[pointY[pos]][pointX[pos]] = 1
	}

	return grid, instructions
}

func applyInstruction(grid [][]int, instruction string) [][]int {
	data := strings.Split(instruction, "=")
	value, err := strconv.Atoi(data[1])
	if err != nil {
		panic(err)
	}

	if data[0] == "fold along y" {
		return foldY(grid, value)
	} else {
		return foldX(grid, value)
	}
}

func countDots(grid [][]int) int {
	dots := 0

	for _, line := range grid {
		for _, number := range line {
			if number == 1 {
				dots++

			}
		}
	}

	return dots
}

func display(grid [][]int) {
	for _, line := range grid {
		for _, number := range line {
			if number == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}

		fmt.Println(" ")
	}
}

func foldY(grid [][]int, foldY int) [][]int {
	maxY := len(grid) - 1
	var newGrid [][]int

	for y := 0; y < foldY; y++ {
		newGrid = append(newGrid, []int{})

		for x := 0; x < len(grid[0]); x++ {
			newGrid[y] = append(newGrid[y], 0)

			if grid[maxY-y][x] == 1 || grid[y][x] == 1 {
				newGrid[y][x] = 1
			}
		}
	}

	return newGrid
}

func foldX(grid [][]int, foldX int) [][]int {
	maxX := len(grid[0]) - 1
	var newGrid [][]int

	for y := 0; y < len(grid); y++ {
		newGrid = append(newGrid, []int{})

		for x := 0; x < foldX; x++ {
			newGrid[y] = append(newGrid[y], 0)

			if grid[y][maxX-x] == 1 || grid[y][x] == 1 {
				newGrid[y][x] = 1
			}
		}
	}

	return newGrid
}
