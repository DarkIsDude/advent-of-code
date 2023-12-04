package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	file, err := os.Open(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumOfGamePossible := 0

	for scanner.Scan() {
		line := scanner.Text()

    split := strings.Split(line, ": ")
    gameAsString := strings.Replace(split[0], "Game ", "", 1)
		game, err := strconv.Atoi(gameAsString)

		if err != nil {
			log.Fatal(err)
		}

		gamePossible := true
		rounds := strings.Split(split[1], ";")

		for _, round := range rounds {
			grabs := strings.Split(round, ",")

			for _, grab := range grabs {
				grabSplitted := strings.Split(strings.TrimSpace(grab), " ")
				unit, err := strconv.Atoi(grabSplitted[0])
				color := grabSplitted[1]

				if err != nil {
					log.Fatal(err)
				}

				switch color {
				case "red":
					if unit > 12 {
						gamePossible = false
						fmt.Printf("Game %d is not possible because of red with %d\n", game, unit)
					}
					break
				case "green":
					if unit > 13 {
						gamePossible = false
						fmt.Printf("Game %d is not possible because of green with %d\n", game, unit)
					}
					break
				case "blue":
					if unit > 14 {
						gamePossible = false
						fmt.Printf("Game %d is not possible because of blue with %d\n", game, unit)
					}
					break
				default:
					log.Fatal("Invalid color")
				}
			}
		}

		if gamePossible {
			sumOfGamePossible += game
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum of game possible:", sumOfGamePossible)
}
