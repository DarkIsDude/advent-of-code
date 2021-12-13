package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const FILE_LOCATION = "./input.txt"

type Path struct {
	caveA string
	caveB string
}

func main() {
	world := readFile()
	toExplore := [][]string{{"start"}}
	explored := true

	for explored {
		explored = false

		for i := len(toExplore) - 1; i >= 0; i-- {
			pathToExplore := toExplore[i]
			fmt.Printf("Exploring %d %v\n", i, pathToExplore)
			lastElement := pathToExplore[len(pathToExplore)-1]

			if lastElement == "end" {
				fmt.Println("Already an endpath")
				continue
			}

			toExplore = append(toExplore[:i], toExplore[i+1:]...)
			explored = true
			endpoints := explore(world, lastElement)

			fmt.Printf("Found new element %v\n", endpoints)

			for _, endpoint := range endpoints {
				newPath := append([]string{}, pathToExplore...)
				newPath = append(newPath, endpoint)

				if isValidPath(newPath) {
					toExplore = append(toExplore, newPath)
				} else {
					fmt.Printf("This new path is not valid %v\n", newPath)
				}
			}
		}
	}

	fmt.Println(len(toExplore))
}

func readFile() []Path {
	var world []Path

	file, err := os.Open(FILE_LOCATION)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		points := strings.Split(text, "-")

		path := Path{
			caveA: points[0],
			caveB: points[1],
		}

		world = append(world, path)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return world
}

func explore(world []Path, cave string) []string {
	var nextPoint []string

	for _, path := range world {
		if path.caveA == cave {
			nextPoint = append(nextPoint, path.caveB)
		}

		if path.caveB == cave {
			nextPoint = append(nextPoint, path.caveA)
		}
	}

	return nextPoint
}

func isValidPath(path []string) bool {
	var smallCave []string

	for _, cave := range path {
		if unicode.IsLower([]rune(cave)[0]) {
			smallCave = append(smallCave, cave)
		}
	}

	numberOfSmallCaseInDouble := 0
	for posA, caveA := range smallCave {
		if posA > 0 && caveA == "start" {
			return false
		}

		for posB, caveB := range smallCave {
			if posA != posB && caveA == caveB {
				numberOfSmallCaseInDouble++
			}
		}
	}

	fmt.Printf("Element in double %d\n", numberOfSmallCaseInDouble)
	return numberOfSmallCaseInDouble/2 <= 1
}
