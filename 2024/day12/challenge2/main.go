package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
	sum := 0

	for y := range garden {
		for x := range garden[y] {
			if seen[y][x] != '.' {
				continue
			}

			uniqueSeen := createEmptySeen(garden)

			currentType := garden[y][x]
			fmt.Printf("Exploring plot at %d, %d of type %c\n", x, y, currentType)
			uniqueSeen, newSeen, size, sidesPoint := getSizeAndSides(garden, uniqueSeen, seen, currentType, x, y)

			sides := countSides(uniqueSeen, sidesPoint)
			fmt.Printf("Plot at %d, %d of type %c has size %d and %d sides\n", x, y, currentType, size, sides)

			seen = newSeen
			sum += size * sides
		}
	}

	fmt.Printf("Total sum is %d\n", sum)
}

func countSides(garden []string, sides [][]int) int {
	sidesByDirection := map[string][][]int{
		"up":    {},
		"down":  {},
		"left":  {},
		"right": {},
	}

	for _, side := range sides {
		if isInTheGardenAndMatches(garden, side[0], side[1]-1) {
			sidesByDirection["up"] = append(sidesByDirection["up"], side)
		}

		if isInTheGardenAndMatches(garden, side[0], side[1]+1) {
			sidesByDirection["down"] = append(sidesByDirection["down"], side)
		}

		if isInTheGardenAndMatches(garden, side[0]-1, side[1]) {
			sidesByDirection["left"] = append(sidesByDirection["left"], side)
		}

		if isInTheGardenAndMatches(garden, side[0]+1, side[1]) {
			sidesByDirection["right"] = append(sidesByDirection["right"], side)
		}
	}

	slices.SortFunc(sidesByDirection["up"], func(a []int, b []int) int {
		return (a[1]*10000 + a[0]) - (b[1]*10000 + b[0])
	})

	slices.SortFunc(sidesByDirection["down"], func(a []int, b []int) int {
		return (a[1]*10000 + a[0]) - (b[1]*10000 + b[0])
	})

	slices.SortFunc(sidesByDirection["left"], func(a []int, b []int) int {
		return (a[0]*10000 + a[1]) - (b[0]*10000 + b[1])
	})

	slices.SortFunc(sidesByDirection["right"], func(a []int, b []int) int {
		return (a[0]*10000 + a[1]) - (b[0]*10000 + b[1])
	})

	count := 0

	for _, direction := range []string{"up", "down", "left", "right"} {
		fmt.Println(direction, sidesByDirection[direction])
		if len(sidesByDirection[direction]) > 0 {
			count++
		}

		for i := 1; i < len(sidesByDirection[direction]); i++ {
			diffX := sidesByDirection[direction][i][0] - sidesByDirection[direction][i-1][0]
			diffY := sidesByDirection[direction][i][1] - sidesByDirection[direction][i-1][1]

			if direction == "up" || direction == "down" {
				if diffX > 1 || diffY != 0 {
					fmt.Println("Adding 1 for", direction)
					count++
				}
			} else {
				if diffY > 1 || diffX != 0 {
					fmt.Println("Adding 1 for", direction)
					count++
				}
			}
		}
	}

	return count
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

func getSizeAndSides(garden []string, uniqueSeen []string, seen []string, plot byte, x int, y int) ([]string, []string, int, [][]int) {
	if !isInTheGarden(garden, x, y) {
		return uniqueSeen, seen, 0, [][]int{{x, y}}
	}

	if garden[y][x] != plot {
		return uniqueSeen, seen, 0, [][]int{{x, y}}
	}

	if seen[y][x] != '.' {
		return uniqueSeen, seen, 0, [][]int{}
	}

	seen[y] = seen[y][:x] + "#" + seen[y][x+1:]
	uniqueSeen[y] = uniqueSeen[y][:x] + "#" + uniqueSeen[y][x+1:]
	size := 1
	sides := [][]int{}

	for _, diff := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		newX := x + diff[0]
		newY := y + diff[1]
		newUniqueSeen, newSeen, newSize, newSides := getSizeAndSides(garden, uniqueSeen, seen, plot, newX, newY)
		size += newSize
		sides = append(sides, newSides...)
		uniqueSeen = newUniqueSeen
		seen = newSeen
	}

	return uniqueSeen, seen, size, sides
}

func isInTheGardenAndMatches(garden []string, x int, y int) bool {
	return isInTheGarden(garden, x, y) && garden[y][x] == '#'
}

func isInTheGarden(garden []string, x int, y int) bool {
	return x >= 0 && x < len(garden[0]) && y >= 0 && y < len(garden)
}
