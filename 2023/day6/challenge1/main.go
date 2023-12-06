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
	times := []int{}
	distances := []int{}

	scanner.Scan()
	line := scanner.Text()
	line = strings.Replace(line, "Time:", "", -1)

	for _, v := range strings.Split(line, " ") {
		if v == "" {
			continue
		}

		t, _ := strconv.Atoi(v)
		times = append(times, t)
	}

	scanner.Scan()
	line = scanner.Text()
	line = strings.Replace(line, "Distance:", "", -1)

	for _, v := range strings.Split(line, " ") {
		if v == "" {
			continue
		}

		t, _ := strconv.Atoi(v)
		distances = append(distances, t)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	multiplyPossibilities := 1

	for i := 0; i < len(times); i++ {
		fmt.Printf("-- %d\t%d\n", times[i], distances[i])
		possibilitiesCount := 0

		for t := 0; t < times[i]; t++ {
			distance := calculateDistance(t, times[i])

			if distance > distances[i] {
				possibilitiesCount++
			}
		}

		multiplyPossibilities *= possibilitiesCount
	}

	fmt.Printf("Possibilities: %d\n", multiplyPossibilities)
}

func calculateDistance(hold int, times int) int {
	return (times - hold) * hold
}
