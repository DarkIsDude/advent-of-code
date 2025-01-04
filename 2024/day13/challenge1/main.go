package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type game struct {
	buttonAX int
	buttonAY int
	buttonBX int
	buttonBY int
	prizeX   int
	prizeY   int
}

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	tokenCost := 0

	for scanner.Scan() {
		lineButtonA := scanner.Text()
		scanner.Scan()
		lineButtonB := scanner.Text()
		scanner.Scan()
		linePrize := scanner.Text()
		scanner.Scan()

		regexpButtonA := regexp.MustCompile(`^Button A: X\+(?<X>\d+), Y\+(?<Y>\d+)$`)
		regexpButtonB := regexp.MustCompile(`^Button B: X\+(?<X>\d+), Y\+(?<Y>\d+)$`)
		regexpPrize := regexp.MustCompile(`^Prize: X=(?<X>\d+), Y=(?<Y>\d+)$`)

		matchButtonA := regexpButtonA.FindStringSubmatch(lineButtonA)
		matchButtonB := regexpButtonB.FindStringSubmatch(lineButtonB)
		matchPrize := regexpPrize.FindStringSubmatch(linePrize)

		game := createGame(matchButtonA, matchButtonB, matchPrize)
		fmt.Println("----- Game:")
		fmt.Println(game.String())
		party := game.play()

		if party[0] == -1 {
			fmt.Println("The game is impossible")
		} else {
			fmt.Printf("Button A: %d | Button B: %d\n", party[0], party[1])
			tokenCost += party[0]*3 + party[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total token cost: %d\n", tokenCost)
}

func createGame(buttonA []string, buttonB []string, prize []string) game {
	buttonAInt, err := strconv.Atoi(buttonA[1])
	if err != nil {
		log.Fatal(err)
	}

	buttonAYInt, err := strconv.Atoi(buttonA[2])
	if err != nil {
		log.Fatal(err)
	}

	buttonBInt, err := strconv.Atoi(buttonB[1])
	if err != nil {
		log.Fatal(err)
	}

	buttonBYInt, err := strconv.Atoi(buttonB[2])
	if err != nil {
		log.Fatal(err)
	}

	prizeXInt, err := strconv.Atoi(prize[1])
	if err != nil {
		log.Fatal(err)
	}

	prizeYInt, err := strconv.Atoi(prize[2])
	if err != nil {
		log.Fatal(err)
	}

	return game{
		buttonAX: buttonAInt,
		buttonAY: buttonAYInt,
		buttonBX: buttonBInt,
		buttonBY: buttonBYInt,
		prizeX:   prizeXInt,
		prizeY:   prizeYInt,
	}
}

func (g game) String() string {
	return fmt.Sprintf("Button A: X=%d, Y=%d\nButton B: X=%d, Y=%d\nPrize: X=%d, Y=%d", g.buttonAX, g.buttonAY, g.buttonBX, g.buttonBY, g.prizeX, g.prizeY)
}

func (g game) play() []int {
	pressACount := 0
	pressBCount := g.prizeY/g.buttonBY + 1

	if pressBCount > 100 {
		pressBCount = 100
	}

	endOfGame := false

	for !endOfGame {
		xResult := pressACount*g.buttonAX + pressBCount*g.buttonBX
		yResult := pressACount*g.buttonAY + pressBCount*g.buttonBY

		if xResult == g.prizeX && yResult == g.prizeY {
			endOfGame = true

			return []int{pressACount, pressBCount}
		}

		if pressACount == 100 || pressBCount == 0 {
			endOfGame = true
			continue
		}

		if xResult > g.prizeX || yResult > g.prizeY {
			pressBCount--
		} else {
			pressACount++
		}
	}

	return []int{-1, -1}
}
