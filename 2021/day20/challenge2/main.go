package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const FILE_LOCATION = "./input.txt"
const DEBUG = true
const ITERATION = 50
const OVERSIZE = 100

func main() {
	matrice, grid := readFile()

	display(grid)

	for i := 0; i < ITERATION; i++ {
		fmt.Println(" ")
		fmt.Printf("----- Iteration %d\n", i+1)
		fmt.Println(" ")

		grid = enhance(grid, matrice)
		display(grid)
	}

	pixelOn := 0
	for y := range grid {
		for x := range grid[y] {
			xBorder := x < OVERSIZE/2 || x > len(grid[0])-(OVERSIZE/2)
			yBorder := y < OVERSIZE/2 || y > len(grid)-(OVERSIZE/2)

			if !xBorder && !yBorder && grid[y][x] {
				pixelOn++
			}
		}
	}

	fmt.Println(" ")
	fmt.Printf("Pixel on : %d\n", pixelOn)
}

func readFile() (string, [][]bool) {
	matrice := ""
	var basicGrid [][]bool
	var extensibleGrid [][]bool

	file, err := os.Open(FILE_LOCATION)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	matrice = scanner.Text()

	if DEBUG {
		fmt.Printf("Matrice found (%d) : %s\n", len(matrice), matrice)
	}

	scanner.Scan()

	for scanner.Scan() {
		text := scanner.Text()

		var line []bool
		for _, c := range text {
			if c == '#' {
				line = append(line, true)
			} else {
				line = append(line, false)
			}
		}

		basicGrid = append(basicGrid, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	oversize := ITERATION + OVERSIZE

	ySize := oversize*2 + len(basicGrid)
	xSize := oversize*2 + len(basicGrid[0])

	for y := 0; y < ySize; y++ {
		var line []bool

		for x := 0; x < xSize; x++ {
			xMatch := x >= oversize && x < oversize+len(basicGrid[0])
			yMatch := y >= oversize && y < oversize+len(basicGrid)

			if xMatch && yMatch {
				line = append(line, basicGrid[y-oversize][x-oversize])
			} else {
				line = append(line, false)
			}
		}

		extensibleGrid = append(extensibleGrid, line)
	}

	return matrice, extensibleGrid
}

func display(grid [][]bool) {
	for y := range grid {
		for x := range grid[y] {
			xBorder := x < OVERSIZE/2 || x > len(grid[0])-(OVERSIZE/2)
			yBorder := y < OVERSIZE/2 || y > len(grid)-(OVERSIZE/2)

			if xBorder || yBorder {
				continue
			}

			if grid[y][x] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}

		fmt.Println(" ")
	}
}

func enhance(grid [][]bool, matrice string) [][]bool {
	var newGrid [][]bool

	for y := 0; y < len(grid); y++ {
		var newLine []bool

		for x := 0; x < len(grid[0]); x++ {
			if x == 0 || y == 0 || x == len(grid[0])-1 || y == len(grid)-1 {
				newLine = append(newLine, false)
				continue
			}

			binary := fmt.Sprintf("%s%s%s%s%s%s%s%s%s",
				binary(grid[y-1][x-1]), binary(grid[y-1][x]), binary(grid[y-1][x+1]),
				binary(grid[y][x-1]), binary(grid[y][x]), binary(grid[y][x+1]),
				binary(grid[y+1][x-1]), binary(grid[y+1][x]), binary(grid[y+1][x+1]),
			)

			value, err := strconv.ParseInt(binary, 2, 64)
			if err != nil {
				panic(err)
			}

			if matrice[value] == '#' {
				newLine = append(newLine, true)
			} else {
				newLine = append(newLine, false)
			}
		}

		newGrid = append(newGrid, newLine)
	}

	return newGrid
}

func binary(b bool) string {
	if b {
		return "1"
	} else {
		return "0"
	}
}
