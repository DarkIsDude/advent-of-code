package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	destination  int
	source		   int
	length       int
}

type mapSeed struct {
	from string
	to	 string
	rules []rule
}

func main() {
	args := os.Args
	file, err := os.Open(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	seedsAsString := parseSeeds(line)
	seeds := []int{}

	for _, seedAsString := range seedsAsString {
		seed, err := strconv.Atoi(seedAsString)

		if err != nil {
			log.Fatal(err)
		}

		seeds = append(seeds, seed)
	}

	scanner.Scan()

	maps := []mapSeed{}
	endOfMapsToParse := false

	for !endOfMapsToParse {
		rule, continueParsing := parseMap(scanner)
		maps = append(maps, rule)

		if !continueParsing {
			endOfMapsToParse = true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	minLocation := 0
	for {
		seed := locationToSeed(minLocation, maps)

		for i := 0; i < len(seeds) / 2; i++ {
			startSeed := seeds[i * 2]
			endSeed := startSeed + seeds[i * 2 + 1]

			if seed >= startSeed && seed <= endSeed {
				fmt.Printf("%d\n", minLocation)
				os.Exit(0)
			}
		}

		minLocation++
	}
}

func locationToSeed(location int, maps []mapSeed) int {
	currentUnit := "location"

	for currentUnit != "seed" {
		for _, mapSeed := range maps {
			if currentUnit == mapSeed.to {
				findRule := false

				for _, rule := range mapSeed.rules {
					if rule.destination <= location && rule.destination + rule.length >= location {
						location = location + rule.source - rule.destination
						currentUnit = mapSeed.from
						findRule = true

						break
					}
				}

				if !findRule {
					currentUnit = mapSeed.from
				}

				break
			}
		}
	}

	return location
}

func seedToLocation(seed int, maps []mapSeed) int {
	currentUnit := "seed"

	for currentUnit != "location" {
		for _, mapSeed := range maps {
			if currentUnit == mapSeed.from {
				findRule := false

				for _, rule := range mapSeed.rules {
					if rule.source <= seed && rule.source + rule.length >= seed {
						seed = seed + rule.destination - rule.source
						currentUnit = mapSeed.to
						findRule = true

						break
					}
				}

				if !findRule {
					currentUnit = mapSeed.to
				}

				break
			}
		}
	}

	return seed
}

func parseMap(scanner *bufio.Scanner) (mapSeed, bool) {
	scanner.Scan()
	line := scanner.Text()
	line = strings.TrimSuffix(line, " map:")

	directions := strings.Split(line, "-")

	mapSeed := mapSeed{
		from: directions[0],
		to: directions[2],
	}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			return mapSeed, true
		}

		numbers := strings.Split(line, " ")

		rule := rule{}
		rule.destination, _ = strconv.Atoi(numbers[0])
		rule.source, _ = strconv.Atoi(numbers[1])
		rule.length, _ = strconv.Atoi(numbers[2])

		mapSeed.rules = append(mapSeed.rules, rule)
	}

	return mapSeed, false
}

func parseSeeds(line string) []string {
	line = strings.TrimPrefix(line, "seeds: ")
	seeds := strings.Split(line, " ")

	return seeds
}
