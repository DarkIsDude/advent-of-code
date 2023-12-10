package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type exploration struct {
	area []string
	exploredArea [][]bool
	currentI int
	currentJ int
	over bool
	iteration int
}

func main() {
	args := os.Args
	file, err := os.Open(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	area := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		area = append(area, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	exploreArea(area)
}

func displayArea(area []string) {
	fmt.Println("-------- Area -------")

	for _, line := range area {
		fmt.Println(line)
	}

	fmt.Println("---------------------")
}

func displayExploredArea(exploredArea [][]bool) {
	fmt.Println("--- Explored Area ---")
	for _, line := range exploredArea {
		for _, char := range line {
			if char {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("---------------------")
}

func displayExploration(e *exploration) {
	fmt.Println("------------- Exploration ---")
	displayArea(e.area)
	displayExploredArea(e.exploredArea)
	fmt.Println("Current position:", e.currentI, e.currentJ, string(e.area[e.currentI][e.currentJ]))
	fmt.Println("-----------------------------")
}

func exploreArea(area []string) {
	explorations := []*exploration{}

	for i, line := range area {
		for j, char := range line {
			if char == 'S' {
				explorations = exploreStartingPoint(area, i, j)
			}
		}
	}

	fmt.Println("Explorations:", len(explorations))
	endAtSExplorations := []*exploration{}
	for _, e := range explorations {
		if e.area[e.currentI][e.currentJ] == 'S' {
			endAtSExplorations = append(endAtSExplorations, e)
		}
	}

	fmt.Println("End at S:", len(endAtSExplorations))
	fmt.Println("Farthest:", endAtSExplorations[0].iteration / 2)

	displayExploration(endAtSExplorations[0])
	scaleExploredAreaByTwo(endAtSExplorations[0])
	displayExploration(endAtSExplorations[0])
	propagateExploration(endAtSExplorations[0])
	displayExploration(endAtSExplorations[0])

	sumRemainingUnexplored := 0
	for i, line := range endAtSExplorations[0].exploredArea {
		for j, tile := range line {
			if !tile && i%2 == 0 && j%2 == 0 {
				sumRemainingUnexplored++
			}
		}
	}

	fmt.Println("Remaining unexplored:", sumRemainingUnexplored)
}

func scaleExploredAreaByTwo(e *exploration) {
	newExploredArea := [][]bool{}

	for i, line := range e.exploredArea {
		newLine := []bool{}
		newNextLine := []bool{}

		for j, tile := range line {
			newLine = append(newLine, tile)
			if j > 0 && (e.area[i][j] == 'S' || e.area[i][j] == 'L' || e.area[i][j] == 'F' || e.area[i][j] == '-') {
				newLine = append(newLine, tile)
			} else {
				newLine = append(newLine, false)
			}

			if i > 0 && (e.area[i][j] == 'S' || e.area[i][j] == '7' || e.area[i][j] == 'F' || e.area[i][j] == '|') {
				newNextLine = append(newNextLine, tile)
			} else {
				newNextLine = append(newNextLine, false)
			}
			newNextLine = append(newNextLine, false)
		}

		newExploredArea = append(newExploredArea, newLine)
		newExploredArea = append(newExploredArea, newNextLine)
	}

	e.exploredArea = newExploredArea
}

func propagateExploration(e *exploration) {
	queue := [][]int{}
	queue = append(queue, []int{0, 0})

	for len(queue) > 0 {
		i := queue[0][0]
		j := queue[0][1]

		queue = queue[1:]

		if i >= 0 && i < len(e.exploredArea) && j >= 0 && j < len(e.exploredArea[0]) && !e.exploredArea[i][j] {
			e.exploredArea[i][j] = true
			queue = append(queue, []int{i + 1, j})
			queue = append(queue, []int{i - 1, j})
			queue = append(queue, []int{i, j + 1})
			queue = append(queue, []int{i, j - 1})
		}
	}
}

func createEmptyArea(area []string) [][]bool {
	emtpyArea := [][]bool{}

	for _, line := range area {
		newLine := []bool{}

		for range line {
			newLine = append(newLine, false)
		}

		emtpyArea = append(emtpyArea, newLine)
	}

	return emtpyArea
}

func exploreStartingPoint(area []string, startI int, startJ int) []*exploration {
	toExplore := []*exploration{}

	for _, i := range []int{1, -1} {
		if startI + i >= 0 && startI + i < len(area) {
			exploration := &exploration{
				area: area,
				exploredArea: createEmptyArea(area),
				currentI: startI + i,
				currentJ: startJ,
				over: false,
			}
			exploration.exploredArea[startI][startJ] = true
			toExplore = append(toExplore, exploration)
		}

		if startJ + i >= 0 && startJ + i < len(area[0]) {
			exploration := &exploration{
				area: area,
				exploredArea: createEmptyArea(area),
				currentI: startI,
				currentJ: startJ + i,
				over: false,
			}
			exploration.exploredArea[startI][startJ] = true
			toExplore = append(toExplore, exploration)
		}
	}

	for !explorationAreOver(toExplore) {
		for _, e := range toExplore {
			if !e.over {
				exploreNextPoint(e)
			}
		}
	}

	return toExplore
}

func explorationAreOver(toExplore []*exploration) bool {
	for _, e := range toExplore {
		if !e.over {
			return false
		}
	}

	return true
}

func exploreNextPoint(e *exploration) bool {
	e.iteration++

	if e.iteration > 10000000 {
		displayExploration(e)
		e.over = true

		return false
	}

	if e.area[e.currentI][e.currentJ] == '.' {
		e.over = true

		return false
	} else if e.area[e.currentI][e.currentJ] == 'S' {
		e.over = true

		return false
	} else if e.area[e.currentI][e.currentJ] == '|' {
		if moveExplorationTo(e, e.currentI + 1, e.currentJ, e.currentI - 1, e.currentJ) {
			return true
		}

		moveExplorationTo(e, e.currentI - 1, e.currentJ, e.currentI + 1, e.currentJ)

		return true
	} else if e.area[e.currentI][e.currentJ] == '-' {
		if moveExplorationTo(e, e.currentI, e.currentJ + 1, e.currentI, e.currentJ - 1) {
			return true
		}

		moveExplorationTo(e, e.currentI, e.currentJ - 1, e.currentI, e.currentJ + 1)
		return true
	} else if e.area[e.currentI][e.currentJ] == 'L' {
		if moveExplorationTo(e, e.currentI - 1, e.currentJ, e.currentI, e.currentJ + 1) {
			return true
		}

		moveExplorationTo(e, e.currentI, e.currentJ + 1, e.currentI - 1, e.currentJ)
		return true
	} else if e.area[e.currentI][e.currentJ] == 'J' {
		if moveExplorationTo(e, e.currentI - 1, e.currentJ, e.currentI, e.currentJ - 1) {
			return true
		}

		moveExplorationTo(e, e.currentI, e.currentJ - 1, e.currentI - 1, e.currentJ)
		return true
	} else if e.area[e.currentI][e.currentJ] == '7' {
		if moveExplorationTo(e, e.currentI + 1, e.currentJ, e.currentI, e.currentJ - 1) {
			return true
		}

		moveExplorationTo(e, e.currentI, e.currentJ - 1, e.currentI + 1, e.currentJ)
		return true
	} else if e.area[e.currentI][e.currentJ] == 'F' {
		if (moveExplorationTo(e, e.currentI + 1, e.currentJ, e.currentI, e.currentJ + 1)) {
			return true
		}

		moveExplorationTo(e, e.currentI, e.currentJ + 1, e.currentI + 1, e.currentJ)
		return true
	} else {
		log.Fatal("Unknown character", string(e.area[e.currentI][e.currentJ]))
	}

	return true
}

func moveExplorationTo(e *exploration, i int, j int, toCheckI int, toCheckJ int) bool {
	if i < 0 || i > len(e.area) - 1 {
		return false
	}

	if j < 0 || j > len(e.area[0]) - 1 {
		return false
	}

	if toCheckI < 0 || toCheckI > len(e.area) - 1 {
		return false
	}

	if toCheckJ < 0 || toCheckJ > len(e.area[0]) - 1 {
		return false
	}

	if !e.exploredArea[toCheckI][toCheckJ] {
		return false
	}

	if e.area[i][j] != 'S' && e.exploredArea[i][j] {
		return false
	}

	e.exploredArea[e.currentI][e.currentJ] = true
	e.currentI = i
	e.currentJ = j

	return true
}
