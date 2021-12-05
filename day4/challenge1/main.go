package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	lotery []string
	grids  []Grid
}

type Grid struct {
	values [][]string
}

func main() {
	data := readFile()

	winner, numberPlayed := play(data)

	if len(winner.values) == 0 {
		panic("WTF")
	}

	fmt.Println(winner.values)
	fmt.Println(numberPlayed)

	calculateAnswer(winner, numberPlayed)
}

func readFile() Data {
	data := Data{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	data.lotery = strings.Split(scanner.Text(), ",")
	gridPosition := -1

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			gridPosition++

			fmt.Println("-------------")

			continue
		}

		if len(data.grids) < gridPosition+1 {
			fmt.Printf("-- Grid %d --\n", gridPosition)
			data.grids = append(data.grids, Grid{})
		}

		numbers := strings.Split(scanner.Text(), " ")
		numbers = deleteEmpty(numbers)

		fmt.Println(numbers)

		data.grids[gridPosition].values = append(data.grids[gridPosition].values, numbers)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return data
}

func deleteEmpty(s []string) []string {
	var r []string

	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}

	return r
}

func play(data Data) (Grid, []string) {
	var winner Grid
	var numberPlayed []string

	for _, lotery := range data.lotery {
		if len(winner.values) > 0 {
			continue
		}

		fmt.Printf("-- We'll play %s with %v\n", lotery, numberPlayed)
		numberPlayed = append(numberPlayed, lotery)

		for gridPos, grid := range data.grids {
			if len(winner.values) > 0 {
				continue
			}

			fmt.Printf("-- -- Checking grid %d\n", gridPos)
			if isWinner(grid, numberPlayed) {
				winner = grid
			}
		}
	}

	return winner, numberPlayed
}

func isWinner(grid Grid, lotery []string) bool {
	winner := false

	// Scan row
	fmt.Println("-- -- Checking row")
	for rowPos, row := range grid.values {
		fmt.Printf("-- -- -- Checking row %d\n", rowPos)
		matchElement := 0

		for _, column := range row {
			if indexOf(lotery, column) >= 0 {
				matchElement++
			}
		}

		fmt.Printf("-- -- -- Found %d elemnt\n", matchElement)
		if matchElement == len(row) {
			winner = true
		}
	}

	// Scan column
	fmt.Println("-- -- Checking column")
	for column := 0; column < len(grid.values[0]); column++ {
		fmt.Printf("-- -- -- Checking column %d\n", column)
		matchElement := 0

		for _, row := range grid.values {
			if indexOf(lotery, row[column]) >= 0 {
				matchElement++
			}
		}

		fmt.Printf("-- -- -- Found %d elemnt\n", matchElement)
		if matchElement == len(grid.values[0]) {
			winner = true
		}
	}

	return winner
}

func calculateAnswer(grid Grid, numberPlayedS []string) {
	lastNumberPlayed, err := strconv.Atoi(numberPlayedS[len(numberPlayedS)-1])
	if err != nil {
		panic(err)
	}

	sumOfElementNotMarked := 0

	for _, row := range grid.values {
		for _, element := range row {
			if indexOf(numberPlayedS, element) == -1 {
				number, err := strconv.Atoi(element)
				if err != nil {
					panic(err)
				}

				sumOfElementNotMarked += number
			}
		}
	}

	fmt.Println("Final result")
	fmt.Println(sumOfElementNotMarked * lastNumberPlayed)
}

func indexOf(array []string, element string) int {
	position := -1

	for pos, value := range array {
		if position >= 0 {
			continue
		}

		if value == element {
			position = pos
		}
	}

	return position
}
