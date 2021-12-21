package main

import "fmt"

const POSITION_START_PLAYER1 = 4
const POSITION_START_PLAYER2 = 5

const DEBUG = true

var rollTheDiceCounter int = 0

func main() {
	positionPlayer1 := POSITION_START_PLAYER1
	scorePlayer1 := 0

	positionPlayer2 := POSITION_START_PLAYER2
	scorePlayer2 := 0

	theDice := 1

	currentPlayer := 1

	for scorePlayer1 < 1000 && scorePlayer2 < 1000 {
		if currentPlayer == 1 {
			fmt.Println("Player 1")

			theDice, positionPlayer1, scorePlayer1 = play(theDice, positionPlayer1, scorePlayer1)
			currentPlayer = 2

		} else {
			fmt.Println("Player 2")
			theDice, positionPlayer2, scorePlayer2 = play(theDice, positionPlayer2, scorePlayer2)
			currentPlayer = 1
		}
	}

	if currentPlayer == 1 {
		fmt.Printf("Final result is %d*%d = %d\n", rollTheDiceCounter, scorePlayer1, rollTheDiceCounter*scorePlayer1)
	} else {
		fmt.Printf("Final result is %d*%d = %d\n", rollTheDiceCounter, scorePlayer2, rollTheDiceCounter*scorePlayer2)
	}

}

func rollTheDice(theDice int) (int, int) {
	roll := theDice
	theDice++
	rollTheDiceCounter++

	if theDice > 100 {
		theDice = 1
	}

	return roll, theDice
}

func play(theDice int, position int, score int) (int, int, int) {
	dice1, theDice := rollTheDice(theDice)
	dice2, theDice := rollTheDice(theDice)
	dice3, theDice := rollTheDice(theDice)

	diceSum := (dice1 + dice2 + dice3) % 10
	position += diceSum

	if position > 10 {
		position = position % 10
	}

	score = score + position

	fmt.Printf("Player rolls %d+%d+%d moves to space %d for a total of %d\n", dice1, dice2, dice3, position, score)

	return theDice, position, score
}
