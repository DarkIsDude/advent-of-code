package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	args := os.Args
	file, err := os.Open(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	universe := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		universe = append(universe, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	displayUniverse(universe)
	rowsNonEmpty, colsNonEmpty := findRowsAndColsWithGalaxy(universe)
	galaxies := findGalaxies(universe)
	sumOfDistances := sumDistance(galaxies, rowsNonEmpty, colsNonEmpty)

	fmt.Println(sumOfDistances)
}

func sumDistance(galaxies [][]int, rowsNonEmpty map[int]bool, colsNonEmpty map[int]bool) int {
	sumOfDistances := 0

	for i, galaxy := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			distance := calculateDistance(galaxy, galaxies[j], rowsNonEmpty, colsNonEmpty)
			sumOfDistances += distance
		}
	}

	return sumOfDistances
}

func calculateDistance(galaxy1 []int, galaxy2 []int, rowsNonEmpty map[int]bool, colsNonEmpty map[int]bool) int {
	scalableValue := 1_000_000

	minI := math.Min(float64(galaxy1[0]), float64(galaxy2[0]))
	maxI := math.Max(float64(galaxy1[0]), float64(galaxy2[0]))

	minJ := math.Min(float64(galaxy1[1]), float64(galaxy2[1]))
	maxJ := math.Max(float64(galaxy1[1]), float64(galaxy2[1]))

	distanceI := 0
	for i := int(minI) + 1; i <= int(maxI); i++ {
		if !rowsNonEmpty[i] {
			distanceI += scalableValue
		} else {
			distanceI++
		}
	}

	distanceJ := 0
	for j := int(minJ) + 1; j <= int(maxJ); j++ {
		if !colsNonEmpty[j] {
			distanceJ += scalableValue
		} else {
			distanceJ++
		}
	}

	return int(distanceI + distanceJ)
}

func findGalaxies(universe []string) [][]int {
	var galaxies [][]int

	for i, line := range universe {
		for j, char := range line {
			if char == '#' {
				galaxies = append(galaxies, []int{i, j})
			}
		}
	}

	return galaxies
}

func displayUniverse(universe []string) {
	for _, line := range universe {
		fmt.Println(line)
	}
}

func findRowsAndColsWithGalaxy(universe []string) (map[int]bool, map[int]bool) {
	cols := map[int]bool{}
	rows := map[int]bool{}

	for i, line := range universe {
		for j, char := range line {
			if char == '#' {
				cols[j] = true
				rows[i] = true
			}
		}
	}

	return rows, cols
}
