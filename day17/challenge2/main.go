package main

import (
	"fmt"
	"math"
)

const DEBUG = false

const LEFT_X = 102
const RIGHT_X = 157
const BOTTOM_Y = -146
const TOP_Y = -90

/*
const LEFT_X = 20
const RIGHT_X = 30
const BOTTOM_Y = -10
const TOP_Y = -5
*/

func main() {
	maxY := 0
	velocityAvailable := 0

	for xVelocity := -500; xVelocity < 500; xVelocity++ {
		for yVelocity := -500; yVelocity < 500; yVelocity++ {
			ok, maxYLocal := launchProbe(xVelocity, yVelocity)

			if ok {
				fmt.Printf("Valid velocity for %d and %d\n", xVelocity, yVelocity)
				velocityAvailable++

				if maxYLocal > maxY {
					maxY = maxYLocal
				}
			}
		}
	}

	fmt.Println(maxY)
	fmt.Println(velocityAvailable)
}

func launchProbe(xVelocity int, yVelocity int) (bool, int) {
	maxY := 0
	x := 0
	y := 0

	for x <= RIGHT_X && y >= BOTTOM_Y {
		x += xVelocity
		y += yVelocity

		if DEBUG {
			fmt.Printf("%d:%d | %d\n", x, y, yVelocity)
		}

		if y > maxY {
			maxY = y
		}

		if x >= LEFT_X && x <= RIGHT_X && y <= TOP_Y && y >= BOTTOM_Y {
			return true, maxY
		}

		xVelocity = int(math.Max(float64(xVelocity)-1, 0))
		yVelocity--
	}

	return false, 0
}
