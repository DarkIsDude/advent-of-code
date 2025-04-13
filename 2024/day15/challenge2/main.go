package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var movements = [][]int{}

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

	warehouse = migrateWarehouse(warehouse)
	displayMoves(moves)
	displayWarehouse(warehouse)

	for _, m := range moves {
		warehouse = move(warehouse, m)
	}

	displayWarehouse(warehouse)
	fmt.Println("Distance:", calculateDistance(warehouse))
}

func displayWarehouse(warehouse []string) {
	fmt.Print("   ")
	for i := range warehouse[0] {
		fmt.Print(i % 10)
	}
	fmt.Println()

	for i, warehouseLine := range warehouse {
		fmt.Printf("%d: %s\n", i, warehouseLine)
	}
}

func displayMoves(moves []rune) {
	for _, move := range moves {
		fmt.Print(string(move))
	}

	fmt.Println()
}

func migrateWarehouse(warehouse []string) []string {
	newWarehouse := []string{}

	for _, warehouseLine := range warehouse {
		newLine := ""

		for _, char := range warehouseLine {
			switch char {
			case '#':
				newLine += "##"
			case 'O':
				newLine += "[]"
			case '@':
				newLine += "@."
			case '.':
				newLine += ".."
			}
		}

		newWarehouse = append(newWarehouse, newLine)
	}

	return newWarehouse
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

func move(warehouse []string, direction rune) []string {
	x, y := findStart(warehouse)
	directionX, directionY := 0, 0
	switch direction {
	case '>':
		directionX = 1
		directionY = 0
	case '<':
		directionX = -1
		directionY = 0
	case '^':
		directionX = 0
		directionY = -1
	case 'v':
		directionX = 0
		directionY = 1
	}

	movements = [][]int{}
	canMove := canMoveTo(warehouse, x, y, directionX, directionY)
	if canMove {
		fmt.Printf("Moved %c to %d,%d with direction %d,%d\n", direction, x+directionX, y+directionY, directionX, directionY)

		newWarehouse := duplicateWarehouse(warehouse)

		for _, move := range movements {
			newWarehouse[move[1]] = newWarehouse[move[1]][:move[0]] + "." + newWarehouse[move[1]][move[0]+1:]
		}

		for _, m := range movements {
			currentRune := warehouse[m[1]][m[0]]
			computedY := m[1] + directionY
			computedX := m[0] + directionX
			newWarehouse[computedY] = newWarehouse[computedY][:computedX] + string(currentRune) + newWarehouse[computedY][computedX+1:]

			if currentRune == '@' {
				newWarehouse[m[1]] = newWarehouse[m[1]][:m[0]] + "." + newWarehouse[m[1]][m[0]+1:]
			}
		}

		warehouse = newWarehouse
	} else {
		fmt.Printf("Can't move %c to %d,%d with direction %d,%d\n", direction, x+directionX, y+directionY, directionX, directionY)
	}

	return warehouse
}

func duplicateWarehouse(warehouse []string) []string {
	newWarehouse := make([]string, len(warehouse))

	for i, warehouseLine := range warehouse {
		newWarehouse[i] = warehouseLine
	}

	return newWarehouse
}

func canMoveTo(warehouse []string, x, y, directionX, directionY int) bool {
	movements = append(movements, []int{x, y})

	if canMoveToWall(warehouse, x, y, directionX, directionY) {
		return false
	}

	if canMoveToFreespace(warehouse, x, y, directionX, directionY) {
		return true
	}

	return canMoveToBox(warehouse, x, y, directionX, directionY)
}

func canMoveToWall(warehouse []string, x, y, directionX, directionY int) bool {
	return warehouse[y+directionY][x+directionX] == '#'
}

func canMoveToFreespace(warehouse []string, x, y, directionX, directionY int) bool {
	return warehouse[y+directionY][x+directionX] == '.'
}

func canMoveToBox(warehouse []string, x, y, directionX, directionY int) bool {
	targetRune := warehouse[y+directionY][x+directionX]

	if targetRune != '[' && targetRune != ']' {
		return false
	}

	xOffset := 0
	if targetRune == '[' {
		xOffset = 1
	} else {
		xOffset = -1
	}

	canMove := canMoveTo(warehouse, x+directionX, y+directionY, directionX, directionY)
	canMoveToOffset := true

	if directionY != 0 {
		canMoveToOffset = canMoveTo(warehouse, x+directionX+xOffset, y+directionY, directionX, directionY)
	}

	return canMove && canMoveToOffset
}

func isAlreadyRegistered(x, y int) bool {
	for _, move := range movements {
		if move[0] == x && move[1] == y {
			return true
		}
	}

	return false
}

func calculateDistance(warehouse []string) int {
	sum := 0

	for i, warehouseLine := range warehouse {
		for j, char := range warehouseLine {
			if char == '[' {
				sum += 100*i + j
			}
		}
	}

	return sum
}
