package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const FILE_LOCATION = "./input.txt"
const DEBUG_PLUS = false
const DEBUG = true

type Beacon struct {
	x int
	y int
	z int
}

type Scanner struct {
	beacons []Beacon
	name    int

	x int
	y int
	z int

	matched bool
}

func main() {
	scanners := readFile()

	if DEBUG_PLUS {
		for _, scanner := range scanners {
			fmt.Printf("Scanner %d (size %d)\n", scanner.name, len(scanner.beacons))

			for _, beacon := range scanner.beacons {
				fmt.Printf("%d:%d:%d\n", beacon.x, beacon.y, beacon.z)
			}
		}
	}

	scanners[0].matched = true
	scanners[0].x = 0
	scanners[0].y = 0
	scanners[0].z = 0
	atLeastOneFound := true
	for atLeastOneFound {
		atLeastOneFound = false

		for i := 0; i < len(scanners); i++ {
			if scanners[i].matched {
				if DEBUG_PLUS {
					fmt.Printf("Scanner %d already matched\n", scanners[i].name)
				}

				continue
			}

			if DEBUG_PLUS {
				fmt.Printf("Testing target %d with origin %d\n", scanners[i].name, scanners[0].name)
			}

			ok, beacons, xScanner, yScanner, zScanner := rotateScanner2AndCompare(scanners[0].beacons, scanners[i].beacons)
			if ok {
				atLeastOneFound = true
				scanners[i].matched = true
				scanners[i].x = xScanner
				scanners[i].y = yScanner
				scanners[i].z = zScanner

				fmt.Printf(
					"New scanner position detected %d (%d %d %d)\n",
					scanners[i].name,
					scanners[i].x,
					scanners[i].y,
					scanners[i].z,
				)

				for _, beacon := range beacons {
					alreadyFound := false
					for _, knowBeacon := range scanners[0].beacons {
						if equalBeacon(beacon, knowBeacon) {
							alreadyFound = true
						}
					}

					if DEBUG_PLUS {
						fmt.Printf("Testing new beacon %d %d %d : %v\n", beacon.x, beacon.y, beacon.z, alreadyFound)
					}

					if !alreadyFound {
						scanners[0].beacons = append(scanners[0].beacons, beacon)
					}
				}
			}
		}
	}

	fmt.Printf("Total elements : %d\n", len(scanners[0].beacons))

	maxDistance := 0
	for _, scanner1 := range scanners {
		for _, scanner2 := range scanners {
			distance := manhattanDistance(scanner1, scanner2)

			if distance > maxDistance {
				if DEBUG_PLUS {
					fmt.Printf("New max distance detected for %d %d : %d\n", scanner1.name, scanner2.name, distance)
				}

				maxDistance = distance
			}
		}
	}

	fmt.Printf("Max distance is %d\n", maxDistance)
}

// ### READING FILE

func readFile() []Scanner {
	var scanners []Scanner

	file, err := os.Open(FILE_LOCATION)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentScanner := Scanner{}

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			if DEBUG_PLUS {
				fmt.Println("Empty line")
			}

			scanners = append(scanners, currentScanner)
		} else if strings.HasPrefix(text, "---") {
			if DEBUG_PLUS {
				fmt.Println("A new scanner")
				fmt.Println(text)
			}

			r := regexp.MustCompile(`--- scanner (\d+) ---`)
			matches := r.FindStringSubmatch(text)

			number, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}

			currentScanner = Scanner{
				matched: false,
			}
			currentScanner.name = number
			currentScanner.beacons = make([]Beacon, 0)
		} else {
			if DEBUG_PLUS {
				fmt.Printf("A new position to %d (size %d)\n", currentScanner.name, len(currentScanner.beacons))
			}

			positions := strings.Split(text, ",")
			x, err := strconv.Atoi(positions[0])
			if err != nil {
				panic(err)
			}

			y, err := strconv.Atoi(positions[1])
			if err != nil {
				panic(err)
			}

			z, err := strconv.Atoi(positions[2])
			if err != nil {
				panic(err)
			}

			beacon := Beacon{
				x: x,
				y: y,
				z: z,
			}

			currentScanner.beacons = append(currentScanner.beacons, beacon)
		}
	}

	scanners = append(scanners, currentScanner)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return scanners
}

// Scanner operation

func rotateScanner2AndCompare(beacons1 []Beacon, beacons2 []Beacon) (bool, []Beacon, int, int, int) {
	for decal := 0; decal < 6; decal++ {
		for _, xSign := range []int{1, -1} {
			for _, ySign := range []int{1, -1} {
				for _, zSign := range []int{1, -1} {
					if DEBUG_PLUS {
						fmt.Printf("Decalage de %d avec sign %d %d %d\n", decal, xSign, ySign, zSign)
					}

					newBeacons2 := rotateBeacons(decal, xSign, ySign, zSign, beacons2)
					ok, diffX, diffY, diffZ := compareScanner(beacons1, newBeacons2)

					if ok {
						if DEBUG_PLUS {
							fmt.Printf("The final rotation is %d with %d %d %d - %d %d %d\n", decal, xSign, ySign, zSign, diffX, diffY, diffZ)
						}

						return true, moveBeacons(diffX, diffY, diffZ, newBeacons2, true), diffX, diffY, diffZ
					}
				}
			}
		}
	}

	return false, []Beacon{}, 0, 0, 0
}

func manhattanDistance(scanner1 Scanner, scanner2 Scanner) int {
	x := manhattanDistancePoint(scanner1.x, scanner2.x)
	y := manhattanDistancePoint(scanner1.y, scanner2.y)
	z := manhattanDistancePoint(scanner1.z, scanner2.z)

	if DEBUG_PLUS {
		fmt.Printf("Distance between %d %d : %d (%d %d %d)\n", scanner1.name, scanner2.name, x+y+z, x, y, z)
	}

	return x + y + z
}

func manhattanDistancePoint(a int, b int) int {
	if (a > 0 && b > 0) || (a < 0 && b < 0) {
		if a > b {
			return a - b
		} else {
			return b - a
		}

	} else {
		absA := int(math.Abs(float64(a)))
		absB := int(math.Abs(float64(b)))

		return absA + absB
	}

}

// Return true or false if beacons2 overlappse beacons1
// If yes, return the (x, y, z) needed to add to all beacons of 2 to match 1
func compareScanner(beacons1 []Beacon, beacons2 []Beacon) (bool, int, int, int) {
	for _, beacon1 := range beacons1 {
		newBeaconsRelativeTo1 := moveBeacons(beacon1.x, beacon1.y, beacon1.z, beacons1, false)

		for _, beacon2 := range beacons2 {
			newBeaconsRelativeTo2 := moveBeacons(beacon2.x, beacon2.y, beacon2.z, beacons2, false)

			matchingFound := 0

			for _, newBeaconRelativeTo1 := range newBeaconsRelativeTo1 {
				for _, newBeaconRelativeTo2 := range newBeaconsRelativeTo2 {
					if DEBUG_PLUS {
						fmt.Printf(
							"Compare %d:%d:%d (%d:%d:%d) with %d:%d:%d (%d:%d:%d)\n",
							newBeaconRelativeTo1.x,
							newBeaconRelativeTo1.y,
							newBeaconRelativeTo2.z,
							beacon1.x,
							beacon1.y,
							beacon1.z,
							newBeaconRelativeTo2.x,
							newBeaconRelativeTo2.y,
							newBeaconRelativeTo2.z,
							beacon2.x,
							beacon2.y,
							beacon2.z,
						)
					}

					if equalBeacon(newBeaconRelativeTo1, newBeaconRelativeTo2) {
						matchingFound++

						if DEBUG_PLUS {
							fmt.Printf("New match %d\n", matchingFound)
						}
					}
				}
			}

			if matchingFound >= 12 {
				xDifference := beacon1.x - beacon2.x
				yDifference := beacon1.y - beacon2.y
				zDifference := beacon1.z - beacon2.z

				if DEBUG_PLUS {
					fmt.Printf("Compensation found %d:%d:%d / %d:%d:%d (for %d elements)\n", beacon1.x, beacon1.y, beacon1.z, beacon2.x, beacon2.y, beacon2.z, matchingFound)
					fmt.Printf("Need to add %d and %d and %d to match both\n", xDifference, yDifference, zDifference)
				}

				return true, xDifference, yDifference, zDifference
			}
		}
	}

	return false, 0, 0, 0
}

// Beacon operation

func equalBeacon(beacon1 Beacon, beacon2 Beacon) bool {
	return beacon1.x == beacon2.x && beacon1.y == beacon2.y && beacon1.z == beacon2.z
}

func moveBeacons(originX int, originY int, originZ int, beacons []Beacon, addition bool) []Beacon {
	var newBeacons []Beacon

	for _, beacon := range beacons {
		if addition {
			newBeacons = append(newBeacons, Beacon{
				x: beacon.x + originX,
				y: beacon.y + originY,
				z: beacon.z + originZ,
			})
		} else {
			newBeacons = append(newBeacons, Beacon{
				x: beacon.x - originX,
				y: beacon.y - originY,
				z: beacon.z - originZ,
			})
		}
	}

	return newBeacons
}

func rotateBeacons(decal int, xSign int, ySign, zSign int, beacons []Beacon) []Beacon {
	var newBeacons []Beacon

	for _, beacon := range beacons {
		newBeacon := Beacon{}

		switch decal {
		case 0:
			newBeacon.x = beacon.x
			newBeacon.y = beacon.y
			newBeacon.z = beacon.z
		case 1:
			newBeacon.x = beacon.x
			newBeacon.y = beacon.z
			newBeacon.z = beacon.y
		case 2:
			newBeacon.x = beacon.y
			newBeacon.y = beacon.x
			newBeacon.z = beacon.z
		case 3:
			newBeacon.x = beacon.y
			newBeacon.y = beacon.z
			newBeacon.z = beacon.x
		case 4:
			newBeacon.x = beacon.z
			newBeacon.y = beacon.x
			newBeacon.z = beacon.y
		case 5:
			newBeacon.x = beacon.z
			newBeacon.y = beacon.y
			newBeacon.z = beacon.x
		}

		newBeacon.x = newBeacon.x * xSign
		newBeacon.y = newBeacon.y * ySign
		newBeacon.z = newBeacon.z * zSign

		newBeacons = append(newBeacons, newBeacon)
	}

	return newBeacons
}
