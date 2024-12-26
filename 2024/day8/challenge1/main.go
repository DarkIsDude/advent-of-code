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
	input := []string{}
	antinodes := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		input = append(input, line)
		antinodes = append(antinodes, strings.Repeat(".", len(line)))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	frequences := parseFrequencies(input)

	for _, value := range frequences {
		antinodes = findAntinode(antinodes, value)
	}

	display(antinodes)

	countAntinodes := 0
	for _, line := range antinodes {
		for _, char := range line {
			if char == 'X' {
				countAntinodes++
			}
		}
	}

	fmt.Println("Antinodes:", countAntinodes)
}

func parseFrequencies(input []string) map[rune][][]int {
	values := map[rune][][]int{}

	for y, line := range input {
		for x, char := range line {
			if char != '.' {
				values[char] = append(values[char], []int{x, y})
			}
		}
	}

	return values
}

func findAntinode(antinodes []string, antenna [][]int) []string {
	for i := 0; i < len(antenna); i++ {
		for j := 0; j < len(antenna); j++ {
			if i == j {
				continue
			}

			distanceX := (antenna[i][0] - antenna[j][0])
			distanceY := (antenna[i][1] - antenna[j][1])

			fmt.Println("Distance X:", distanceX, "Distance Y:", distanceY)

			antinodeX := antenna[i][0] + distanceX
			antinodeY := antenna[i][1] + distanceY

			if antinodeX >= 0 && antinodeX < len(antinodes[0]) && antinodeY >= 0 && antinodeY < len(antinodes) {
				antinodes[antinodeY] = antinodes[antinodeY][:antinodeX] + "X" + antinodes[antinodeY][antinodeX+1:]
			}
		}
	}

	return antinodes
}

func display(input []string) {
	for _, line := range input {
		fmt.Println(line)
	}
}
