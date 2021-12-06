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
	fishes := readFile()

	for day := 1; day <= 256; day++ {
		newFishes := map[int]int{
			0: fishes[1],
			1: fishes[2],
			2: fishes[3],
			3: fishes[4],
			4: fishes[5],
			5: fishes[6],
			6: fishes[7],
			7: fishes[8],
		}

		newFishes[8] += fishes[0]
		newFishes[6] += fishes[0]

		fishes = newFishes
		sum := 0
		for _, value := range newFishes {
			sum += value
		}

		fmt.Printf("After %d days, we have %d fish\n", day, sum)
	}

}

func readFile() map[int]int {
	fishes := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	fishS := strings.Split(scanner.Text(), ",")

	for _, s := range fishS {
		fish, err := strconv.Atoi(s)

		if err != nil {
			panic(err)
		}

		fishes[fish]++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return fishes
}
