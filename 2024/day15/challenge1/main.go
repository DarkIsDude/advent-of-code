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
	warehouse := []string{}
	moves := []rune{}
	warehouseDone := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			warehouseDone = true
			continue
		}

		if !warehouseDone {
			warehouse = append(warehouse, line)
			continue
		}

		if warehouseDone {
			moves = append(moves, []rune(line)...)
			continue
		}
	}

	displayWarehouse(warehouse)
	displayMoves(moves)

	for _, m := range moves {
		move(warehouse, m)
	}

	displayWarehouse(warehouse)

	fmt.Println("Distance:", calculateDistance(warehouse))
}

func displayWarehouse(warehouse []string) {
	for _, warehouseLine := range warehouse {
		fmt.Println(warehouseLine)
	}
}

func displayMoves(moves []rune) {
	for _, move := range moves {
		fmt.Print(string(move))
	}

	fmt.Println()
}

func findStart(warehouse []string) (int, int) {
	for i, warehouseLine := range warehouse {
		for j, char := range warehouseLine {
			if char == '@' {
				return j, i
			}
		}
	}

	panic("No start position found")
}

func move(warehouse []string, direction rune) {
	x, y := findStart(warehouse)

	switch direction {
	case '>':
		moveTo(warehouse, x, y, 1, 0)
	case '<':
		moveTo(warehouse, x, y, -1, 0)
	case '^':
		moveTo(warehouse, x, y, 0, -1)
	case 'v':
		moveTo(warehouse, x, y, 0, 1)
	}
}

func moveTo(warehouse []string, x, y, directionX, directionY int) {
	if moveToWall(warehouse, x, y, directionX, directionY) {
		return
	}

	if moveToFreespace(warehouse, x, y, directionX, directionY) {
		return
	}

	moveToBox(warehouse, x, y, directionX, directionY)
}

func moveToWall(warehouse []string, x, y, directionX, directionY int) bool {
	return warehouse[y+directionY][x+directionX] == '#'
}

func moveToFreespace(warehouse []string, x, y, directionX, directionY int) bool {
	currentRune := warehouse[y][x]

	if warehouse[y+directionY][x+directionX] == '.' {
		warehouse[y+directionY] = warehouse[y+directionY][:x+directionX] + string(currentRune) + warehouse[y+directionY][x+directionX+1:]
		warehouse[y] = warehouse[y][:x] + "." + warehouse[y][x+1:]

		return true
	}

	return false
}

func moveToBox(warehouse []string, x, y, directionX, directionY int) bool {
	currentRune := warehouse[y][x]
	moveTo(warehouse, x+directionX, y+directionY, directionX, directionY)

	if warehouse[y+directionY][x+directionX] == '.' {
		warehouse[y+directionY] = warehouse[y+directionY][:x+directionX] + string(currentRune) + warehouse[y+directionY][x+directionX+1:]
		warehouse[y] = warehouse[y][:x] + "." + warehouse[y][x+1:]
	}

	return true
}

func calculateDistance(warehouse []string) int {
	sum := 0
	for i, warehouseLine := range warehouse {
		for j, char := range warehouseLine {
			if char == 'O' {
				sum += 100*i + j
			}
		}
	}

	return sum
}
