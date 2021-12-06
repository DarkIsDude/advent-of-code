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

	for day := 1; day <= 80; day++ {
		totalFishes := len(fishes)

		for i := 0; i < totalFishes; i++ {
			if fishes[i] == 0 {
				fishes[i] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[i]--
			}
		}

		fmt.Println((fishes))
		fmt.Printf("After %d days, we have %d fish\n", day, len(fishes))
	}

}

func readFile() []int {
	var fishes []int

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

		fishes = append(fishes, fish)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return fishes
}
