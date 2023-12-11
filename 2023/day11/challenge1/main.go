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

	universe = scaleUniverse(universe)
	displayUniverse(universe)
	galaxies := findGalaxies(universe)
	sumOfDistances := sumDistance(galaxies)

	fmt.Println(sumOfDistances)
}

func sumDistance(galaxies [][]int) int {
	sumOfDistances := 0

	for i, galaxy := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			distance := calculateDistance(galaxy, galaxies[j])
			sumOfDistances += distance
		}
	}

	return sumOfDistances
}

func calculateDistance(galaxy1 []int, galaxy2 []int) int {
	distanceI := math.Max(float64(galaxy1[0]), float64(galaxy2[0])) - math.Min(float64(galaxy1[0]), float64(galaxy2[0]))
	distanceJ := math.Max(float64(galaxy1[1]), float64(galaxy2[1])) - math.Min(float64(galaxy1[1]), float64(galaxy2[1]))

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

func scaleUniverse(universe []string) []string {
	rows, cols := findRowsAndColsWithGalaxy(universe)
	var scaledUniverse []string

	for i, line := range universe {
		newLine := ""
		newNextLine := ""

		for j, char := range line {
			newLine += string(char)
			newNextLine += "."

			if !cols[j] {
				newLine += "."
				newNextLine += "."
			}
		}

		scaledUniverse = append(scaledUniverse, newLine)

		if !rows[i] {
			scaledUniverse = append(scaledUniverse, newNextLine)
		}
	}

	return scaledUniverse
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
