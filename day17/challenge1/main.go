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
const MIN_X = 20
const MAX_X = 30
const MIN_Y = -5
const MAX_Y = -10
*/

func main() {
	maxY := 0
	var xVelocities []int

	for i := 0; i < RIGHT_X; i++ {
		if possibleX(i) {
			xVelocities = append(xVelocities, i)
		}
	}

	fmt.Printf("Possible X : %v\n", xVelocities)

	for _, xVelocity := range xVelocities {
		for yVelocity := 0; yVelocity < 500; yVelocity++ {
			ok, maxYLocal := launchProbe(xVelocity, yVelocity)

			if ok && maxYLocal > maxY {
				fmt.Printf("Valid velocity for %d and %d\n", xVelocity, yVelocity)
				maxY = maxYLocal
			}
		}
	}

	fmt.Println(maxY)
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

func possibleX(velocity int) bool {
	x := 0

	for i := velocity; i >= 0; i-- {
		x += i

		if x >= LEFT_X && x <= RIGHT_X {
			return true
		}
	}

	return false
}
