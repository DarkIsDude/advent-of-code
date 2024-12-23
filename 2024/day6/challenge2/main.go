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
	game := []string{}
	position := []int{0, 0}
	direction := "right"

	for scanner.Scan() {
		line := scanner.Text()

		for index, char := range line {
			if char != '.' && char != '#' {
				position = []int{len(game), index}

				switch char {
				case '^':
					direction = "up"
				case 'v':
					direction = "down"
				case '<':
					direction = "left"
				case '>':
					direction = "right"
				}

				line = line[:index] + "X" + line[index+1:]
			}
		}

		game = append(game, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	exploreObstructions(game, position, direction)
}

func exploreObstructions(game []string, position []int, direction string) {
	counter := 0

	for i, line := range game {
		for j, char := range line {
			if i == position[0] && j == position[1] {
				continue
			}

			if char == '#' {
				continue
			}

			copyGame := make([]string, len(game))
			copy(copyGame, game)
			copyGame[i] = copyGame[i][:j] + "#" + copyGame[i][j+1:]

			copyPosition := make([]int, len(position))
			copy(copyPosition, position)

			copyDirection := direction

			if !play(copyGame, copyPosition, copyDirection) {
				counter++
			}
		}

		fmt.Println(i)
	}

	fmt.Println(counter)
}

func play(game []string, position []int, direction string) bool {
	counter := 0

	for position[0] != -1 && position[1] != -1 && counter < 20000 {
		game[position[0]] = game[position[0]][:position[1]] + "X" + game[position[0]][position[1]+1:]
		position, direction = move(game, position, direction)
		counter++
	}

	return counter < 20000
}

func move(game []string, position []int, direction string) ([]int, string) {
	nextPosition := calculateDirection(position, direction)

	if nextPosition[0] < 0 || nextPosition[0] >= len(game) || nextPosition[1] < 0 || nextPosition[1] >= len(game[0]) {
		return []int{-1, -1}, direction
	}

	if game[nextPosition[0]][nextPosition[1]] == '#' {
		switch direction {
		case "up":
			direction = "right"
		case "down":
			direction = "left"
		case "left":
			direction = "up"
		case "right":
			direction = "down"
		default:
			panic("Invalid direction")
		}

		return move(game, position, direction)
	}

	return nextPosition, direction
}

func calculateDirection(position []int, direction string) []int {
	switch direction {
	case "up":
		return []int{position[0] - 1, position[1]}
	case "down":
		return []int{position[0] + 1, position[1]}
	case "left":
		return []int{position[0], position[1] - 1}
	case "right":
		return []int{position[0], position[1] + 1}
	}

	panic("Invalid direction")
}

func display(game []string) {
	for _, line := range game {
		fmt.Println(line)
	}
}
