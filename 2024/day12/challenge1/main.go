package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	garden := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		garden = append(garden, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	exploreGarden(garden)
}

func exploreGarden(garden []string) {
	displayGarden(garden)
	seen := createEmptySeen(garden)
	displayGarden(seen)
	sizeByType := map[byte]int{}
	perimeterByType := map[byte]int{}

	sum := 0

	for y := range garden {
		for x := range garden[y] {
			if seen[y][x] != '.' {
				continue
			}

			currentType := garden[y][x]
			fmt.Printf("Exploring plot at %d, %d of type %c\n", x, y, currentType)
			newSeen, size, perimeter := getSizeAndPerimeter(garden, seen, currentType, x, y)

			seen = newSeen
			sizeByType[currentType] += size
			perimeterByType[currentType] += perimeter

			sum += size * perimeter

			displayGarden(seen)
		}
	}

	displayGarden(seen)

	for plot, size := range sizeByType {
		fmt.Printf("Plot of type %c has size %d and perimeter %d\n", plot, size, perimeterByType[plot])
	}

	fmt.Printf("Total sum is %d\n", sum)
}

func createEmptySeen(garden []string) []string {
	seen := []string{}

	for i := range garden {
		seen = append(seen, strings.Repeat(".", len(garden[i])))
	}

	return seen
}

func displayGarden(garden []string) {
	for _, row := range garden {
		fmt.Println(row)
	}

	fmt.Println()
}

func getSizeAndPerimeter(garden []string, seen []string, plot byte, x int, y int) ([]string, int, int) {
	if x < 0 || x >= len(garden[0]) || y < 0 || y >= len(garden) {
		return seen, 0, 1
	}

	if garden[y][x] != plot {
		return seen, 0, 1
	}

	if seen[y][x] != '.' {
		return seen, 0, 0
	}

	seen[y] = seen[y][:x] + "#" + seen[y][x+1:]
	size := 1
	perimeter := 0

	for _, diff := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		newX := x + diff[0]
		newY := y + diff[1]
		newSeen, newSize, newPerimeter := getSizeAndPerimeter(garden, seen, plot, newX, newY)
		size += newSize
		perimeter += newPerimeter
		seen = newSeen
	}

	return seen, size, perimeter
}
