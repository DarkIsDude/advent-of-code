package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const FILE_LOCATION = "./input.txt"
const DEBUG = false

type Instruction struct {
	xStart int
	xEnd   int
	yStart int
	yEnd   int
	zStart int
	zEnd   int

	on bool
}

func main() {
	instructions := readFile()
	grid := initGrid(instructions)

	for pos, instruction := range instructions {
		fmt.Printf("Apply instruction %d\n", pos+1)
		grid = applyInstructions(grid, instruction)
		fmt.Printf("%d cubes are on\n", cubeOn(grid))
	}
}

func readFile() []Instruction {
	var instructions []Instruction

	file, err := os.Open(FILE_LOCATION)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		if DEBUG {
			fmt.Printf("New instruction : %s\n", text)
		}

		r := regexp.MustCompile(`(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)`)
		matches := r.FindStringSubmatch(text)

		on := matches[1] == "on"

		startX, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}

		endX, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}

		startY, err := strconv.Atoi(matches[4])
		if err != nil {
			panic(err)
		}

		endY, err := strconv.Atoi(matches[5])
		if err != nil {
			panic(err)
		}

		startZ, err := strconv.Atoi(matches[6])
		if err != nil {
			panic(err)
		}

		endZ, err := strconv.Atoi(matches[7])
		if err != nil {
			panic(err)
		}

		instructions = append(instructions, Instruction{
			xStart: startX,
			xEnd:   endX,
			yStart: startY,
			yEnd:   endY,
			zStart: startZ,
			zEnd:   endZ,

			on: on,
		})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return instructions
}

func initGrid(instructions []Instruction) [][][]bool {
	var grid [][][]bool

	minX := 0
	minY := 0
	minZ := 0
	maxX := 0
	maxY := 0
	maxZ := 0

	for _, instruction := range instructions {
		// MAX
		if instruction.xEnd > maxX {
			maxX = instruction.xEnd
		}

		if instruction.yEnd > maxY {
			maxY = instruction.yEnd
		}

		if instruction.zEnd > maxZ {
			maxZ = instruction.zEnd
		}

		// MIN
		if instruction.xStart < minX {
			minX = instruction.xStart
		}

		if instruction.yStart > minY {
			minY = instruction.yStart
		}

		if instruction.zStart > minZ {
			minZ = instruction.zStart
		}
	}

	sizeX := 50
	sizeY := 50
	sizeZ := 50

	for x := 0; x <= sizeX*2; x++ {
		grid = append(grid, [][]bool{})

		for y := 0; y <= sizeY*2; y++ {
			grid[x] = append(grid[x], []bool{})

			for z := 0; z <= sizeZ*2; z++ {
				grid[x][y] = append(grid[x][y], false)
			}
		}
	}

	if DEBUG {
		fmt.Printf("Size of the grid : %d %d %d\n", sizeX, sizeY, sizeZ)
	}

	return grid
}

func applyInstructions(grid [][][]bool, instruction Instruction) [][][]bool {
	if DEBUG {
		fmt.Printf("Instruction x(%d,%d) y(%d,%d) z(%d,%d)\n", instruction.xStart, instruction.xEnd, instruction.yStart, instruction.yEnd, instruction.zStart, instruction.zEnd)
	}

	if instruction.xStart < -50 || instruction.xEnd > 50 || instruction.yStart < -50 || instruction.yEnd > 50 || instruction.zStart < -50 || instruction.zEnd > 50 {
		if DEBUG {
			fmt.Printf("Ignore instruction x(%d,%d) y(%d,%d) z(%d,%d)\n", instruction.xStart, instruction.xEnd, instruction.yStart, instruction.yEnd, instruction.zStart, instruction.zEnd)
		}

		return grid
	}

	countCubeOperation := 0

	for x := instruction.xStart; x <= instruction.xEnd; x++ {
		for y := instruction.yStart; y <= instruction.yEnd; y++ {
			for z := instruction.zStart; z <= instruction.zEnd; z++ {
				xPos := len(grid)/2 + x
				yPos := len(grid[0])/2 + y
				zPos := len(grid[0][0])/2 + z

				if DEBUG {
					fmt.Printf("Impact %d,%d,%d\n", xPos, yPos, zPos)
				}

				grid[xPos][yPos][zPos] = instruction.on

				countCubeOperation++
			}
		}
	}

	fmt.Printf("Instruction x(%d,%d) y(%d,%d) z(%d,%d) touch %d cubes with %v\n", instruction.xStart, instruction.xEnd, instruction.yStart, instruction.yEnd, instruction.zStart, instruction.zEnd, countCubeOperation, instruction.on)

	return grid
}

func cubeOn(grid [][][]bool) int {
	counter := 0

	for _, x := range grid {
		for _, y := range x {
			for _, z := range y {
				if z {
					counter++
				}
			}
		}
	}

	return counter
}
