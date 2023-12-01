package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	crabsPosition := readFile()

	minPosition := 1
	minPositionFuel := calculateFuelNeedTo(crabsPosition, minPosition)

	maxPosition := maxPositions(crabsPosition)
	maxPositionFuel := calculateFuelNeedTo(crabsPosition, maxPosition)

	for maxPosition != minPosition {
		positionToTest := minPosition + int((maxPosition-minPosition)/2)
		if positionToTest == minPosition {
			positionToTest++
		}

		fmt.Printf("We are testing the position %d (min %d (%d), max %d (%d))\n", positionToTest, minPosition, minPositionFuel, maxPosition, maxPositionFuel)

		positionToTestFuel := calculateFuelNeedTo(crabsPosition, positionToTest)

		if minPositionFuel > maxPositionFuel {
			minPosition = positionToTest
			minPositionFuel = positionToTestFuel
		} else {
			maxPosition = positionToTest
			maxPositionFuel = positionToTestFuel
		}
	}

	fmt.Printf("Fuel needed for position %d is %d\n", maxPosition, maxPositionFuel)

}

func readFile() []int {
	var crabsPosition []int

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	crabsPositionS := strings.Split(scanner.Text(), ",")

	for _, crabPositionS := range crabsPositionS {
		crabPosition, err := strconv.Atoi(crabPositionS)

		if err != nil {
			panic(err)
		}

		crabsPosition = append(crabsPosition, crabPosition)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return crabsPosition
}

func maxPositions(crabsPosition []int) int {
	maxPosition := 0

	for _, crabPosition := range crabsPosition {
		if maxPosition < crabPosition {
			maxPosition = crabPosition
		}
	}

	return maxPosition
}

func calculateFuelNeedTo(crabsPosition []int, desiredPosition int) int {
	fuelNeeded := 0

	for _, crabPosition := range crabsPosition {
		distance := int(math.Abs(float64(desiredPosition - crabPosition)))
		localFuelNeeded := 0

		for i := 1; i <= distance; i++ {
			localFuelNeeded += i
		}

		fuelNeeded += int(localFuelNeeded)
	}

	return fuelNeeded
}
