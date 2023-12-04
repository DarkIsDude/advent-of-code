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
	finalResult := 0

	for scanner.Scan() {
		line := scanner.Text()

    split := strings.Split(line, ": ")
    // gameAsString := strings.Replace(split[0], "Game ", "", 1)
		// game, err := strconv.Atoi(gameAsString)

		if err != nil {
			log.Fatal(err)
		}

		rounds := strings.Split(split[1], ";")
		fewestBalls := map[string]int {
			"red": 0,
			"blue": 0,
			"green": 0,
		}

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
					if unit > fewestBalls["red"] {
						fewestBalls["red"] = unit
					}
					break
				case "green":
					if unit > fewestBalls["green"] {
						fewestBalls["green"] = unit
					}
					break
				case "blue":
					if unit > fewestBalls["blue"] {
						fewestBalls["blue"] = unit
					}
					break
				default:
					log.Fatal("Invalid color")
				}
			}
		}

		finalResult += fewestBalls["red"] * fewestBalls["green"] * fewestBalls["blue"]
}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(finalResult)
}
