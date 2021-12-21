package main

import "fmt"

const POSITION_START_PLAYER1 = 4
const POSITION_START_PLAYER2 = 5

const DEBUG = true

var player1WinCounter int = 0
var player2WinCounter int = 0

type Score struct {
	player1 int
	player2 int
}

func main() {
	var gridPositions [][][][]Score

	for position1 := 0; position1 <= 10; position1++ {
		gridPositions = append(gridPositions, [][][]Score{})

		for position2 := 0; position2 <= 10; position2++ {
			gridPositions[position1] = append(gridPositions[position1], [][]Score{})

			for score1 := 0; score1 < 21; score1++ {
				gridPositions[position1][position2] = append(gridPositions[position1][position2], []Score{})

				for score2 := 0; score2 < 21; score2++ {
					gridPositions[position1][position2][score1] = append(gridPositions[position1][position2][score1], Score{
						player1: 0,
						player2: 0,
					})
				}
			}
		}
	}

	gridPositions[POSITION_START_PLAYER1][POSITION_START_PLAYER2][0][0].player1 = 1

	playing := true
	for playing {
		playing = false

		for position1 := 1; position1 <= 10; position1++ {
			for position2 := 1; position2 <= 10; position2++ {
				for score1 := 0; score1 < 21; score1++ {
					for score2 := 0; score2 < 21; score2++ {
						if gridPositions[position1][position2][score1][score2].player1 > 0 || gridPositions[position1][position2][score1][score2].player2 > 0 {
							playing = true

							player1 := gridPositions[position1][position2][score1][score2].player1
							player2 := gridPositions[position1][position2][score1][score2].player2
							gridPositions[position1][position2][score1][score2].player1 = 0
							gridPositions[position1][position2][score1][score2].player2 = 0

							for dice1 := 1; dice1 <= 3; dice1++ {
								for dice2 := 1; dice2 <= 3; dice2++ {
									for dice3 := 1; dice3 <= 3; dice3++ {
										newScorePlayer1, newPositionPlayer1 := play(score1, position1, dice1, dice2, dice3)
										if newScorePlayer1 >= 21 {
											player1WinCounter += player1
										} else {
											gridPositions[newPositionPlayer1][position2][newScorePlayer1][score2].player2 += player1
										}

										newScorePlayer2, newPositionPlayer2 := play(score2, position2, dice1, dice2, dice3)
										if newScorePlayer2 >= 21 {
											player2WinCounter += player2
										} else {
											gridPositions[position1][newPositionPlayer2][score1][newScorePlayer2].player1 += player2
										}

									}
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("Player 1 won %d and player 2 won %d\n", player1WinCounter, player2WinCounter)
}

func play(score int, position int, dice1 int, dice2 int, dice3 int) (int, int) {
	diceSum := (dice1 + dice2 + dice3) % 10
	position += diceSum

	if position > 10 {
		position = position % 10
	}

	score = score + position

	return score, position
}
