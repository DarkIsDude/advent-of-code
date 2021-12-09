package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := readFile()
	lowestPointsX, lowestPointsY := lowestPoints(numbers)
	var bassinSize []int

	for i := 0; i < len(lowestPointsX); i++ {
		fmt.Printf("%d:%d\n", lowestPointsX[i], lowestPointsY[i])

		size := sizeOfBasin(numbers, lowestPointsX[i], lowestPointsY[i])

		fmt.Printf("%d:%d => %d\n", lowestPointsX[i], lowestPointsY[i], size)

		bassinSize = append(bassinSize, size)
	}

	bassinSize, max1 := extractMax(bassinSize)
	bassinSize, max2 := extractMax(bassinSize)
	//lint:ignore SA4006 the variable is not recreated, only max3 is created
	bassinSize, max3 := extractMax(bassinSize)

	fmt.Println(max1 * max2 * max3)
}

func extractMax(array []int) ([]int, int) {
	maxElement := 0
	maxPos := 0

	for pos, elem := range array {
		if elem > maxElement {
			maxElement = elem
			maxPos = pos
		}
	}

	return append(array[:maxPos], array[maxPos+1:]...), maxElement
}

func sizeOfBasin(numbers [][]int, x int, y int) int {
	if x > len(numbers[0])-1 || x < 0 {
		return 0
	}

	if y > len(numbers)-1 || y < 0 {
		return 0
	}

	if numbers[y][x] == 9 {
		return 0
	}

	numbers[y][x] = 9

	return 1 +
		sizeOfBasin(numbers, x+1, y) +
		sizeOfBasin(numbers, x-1, y) +
		sizeOfBasin(numbers, x, y+1) +
		sizeOfBasin(numbers, x, y-1)
}

func lowestPoints(numbers [][]int) ([]int, []int) {
	maxX := len(numbers[0]) - 1
	maxY := len(numbers) - 1

	var lowestPointX []int
	var lowestPointY []int

	for y, line := range numbers {
		for x, number := range line {
			lowest := true

			// Left
			if x-1 >= 0 && number >= numbers[y][x-1] {
				lowest = false
			}

			// Top
			if y-1 >= 0 && number >= numbers[y-1][x] {
				lowest = false
			}

			// Right
			if x+1 <= maxX && number >= numbers[y][x+1] {
				lowest = false
			}

			// Bottom
			if y+1 <= maxY && number >= numbers[y+1][x] {
				lowest = false
			}

			if lowest {
				lowestPointX = append(lowestPointX, x)
				lowestPointY = append(lowestPointY, y)
			}
		}
	}

	return lowestPointX, lowestPointY
}

func readFile() [][]int {
	var numbers [][]int

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbers = append(numbers, []int{})

		for _, character := range line {
			integer, err := strconv.Atoi(string(character))
			if err != nil {
				panic(err)
			}

			numbers[lineNumber] = append(numbers[lineNumber], integer)
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return numbers
}
