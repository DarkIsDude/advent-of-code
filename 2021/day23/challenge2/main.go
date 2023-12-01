package main

import (
	"fmt"
	"math"
	"time"
)

const DEBUG = false

//const WINNER = "...........AABBCCDD"
const WINNER = "...........AAAABBBBCCCCDDDD"

const OPEN_SPACE_LENGTH = 11
const ROOM_LENGTH = 4

var ROOMS_POS []int = []int{2, 4, 6, 8}
var ROOMS_LETTER []rune = []rune{'A', 'B', 'C', 'D'}

type Diagram struct {
	elements string

	weight   int
	explored bool
}

func (d Diagram) addWeigth(element rune, distance int) Diagram {
	switch element {
	case 'A':
		d.weight += 1 * distance
	case 'B':
		d.weight += 10 * distance
	case 'C':
		d.weight += 100 * distance
	case 'D':
		d.weight += 1000 * distance
	default:
		panic("Unknown element")
	}

	return d
}

func (d Diagram) equal(other Diagram) bool {
	return d.elements == other.elements
}

func (diagram Diagram) isComplete() bool {
	return diagram.elements == WINNER
}

func (diagram Diagram) display() {
	fmt.Println("#############")

	fmt.Printf("#")
	for i := 0; i < OPEN_SPACE_LENGTH; i++ {
		fmt.Printf("%s", string(diagram.runeAt(true, 0, i)))
	}

	fmt.Println("#")

	for i := 0; i < ROOM_LENGTH; i++ {
		if i == 0 {
			fmt.Printf("###")
		} else {
			fmt.Printf("  #")
		}

		for j := range ROOMS_LETTER {
			fmt.Printf("%s#", string(diagram.runeAt(false, j, i)))
		}

		if i == 0 {
			fmt.Printf("##")
		}

		fmt.Println(" ")
	}

	fmt.Println("  #########")
}

func (d Diagram) duplicate() Diagram {
	return Diagram{
		elements: d.elements,
		weight:   d.weight,
		explored: d.explored,
	}
}

func (d Diagram) canMoveOnOpenSpace(start int, end int, checkStayAtEnd bool) (bool, int) {
	if checkStayAtEnd {
		for _, position := range ROOMS_POS {
			if end == position {
				return false, 0
			}
		}
	}

	if d.runeAt(true, 0, end) != '.' {
		return false, 0
	}

	if start > end {
		temp := end
		end = start
		start = temp
	}

	for i := start + 1; i <= end-1; i++ {
		if d.runeAt(true, 0, i) != '.' {
			return false, 0
		}
	}

	return true, int(math.Abs(float64(start - end)))
}

func (d Diagram) runeAt(openSpace bool, letter int, index int) rune {
	out := []rune(d.elements)

	if openSpace {
		return out[index]
	} else {
		return out[OPEN_SPACE_LENGTH+(letter*ROOM_LENGTH)+index]
	}
}

// Start

func main() {
	var diagrams []Diagram = []Diagram{
		{
			// Example
			//elements: "...........BACDBCDA",

			// Simple test
			//elements: ".....D......ABBCCDA",

			// Input
			//elements: "...........CDACBADB",

			// Complex example
			//elements: "...........BDDACCBDBBACDACA",

			// Complex input
			elements: "...........CDDDACDCBBAADACB",

			weight:   0,
			explored: false,
		},
	}

	completedOne := false
	for !completedOne {
		smallestPos := smallestWeight(diagrams)

		exploreDiagrams := exploreDiagram(diagrams[smallestPos])

		if DEBUG {
			time.Sleep(1 * time.Second)
			fmt.Printf("Exploring %d\n", smallestPos)
			diagrams[smallestPos].display()
		}

		if diagrams[smallestPos].isComplete() {
			diagrams[smallestPos].display()

			fmt.Printf("Minimum weight is %d\n", diagrams[smallestPos].weight)
			completedOne = true

			break
		}

		diagrams[smallestPos].explored = true
		if DEBUG {
			fmt.Printf("%d new diagram to explore\n", len(exploreDiagrams))

		}

		for _, exploreDiagram := range exploreDiagrams {
			found := false

			for pos, existingDiagram := range diagrams {
				if found {
					continue
				}

				if existingDiagram.equal(exploreDiagram) {
					found = true

					if exploreDiagram.weight < existingDiagram.weight {
						diagrams[pos] = exploreDiagram
					}
				}
			}

			if !found {
				diagrams = append(diagrams, exploreDiagram)
			}
		}
	}
}

func smallestWeight(diagrams []Diagram) int {
	smallestDiagramPos := -1

	for pos, diagram := range diagrams {
		if smallestDiagramPos < 0 && !diagram.explored {
			smallestDiagramPos = pos
			continue
		}

		if !diagram.explored && diagram.weight < diagrams[smallestDiagramPos].weight {
			smallestDiagramPos = pos
		}
	}

	return smallestDiagramPos
}

// Explore

func exploreDiagram(d Diagram) []Diagram {
	newDiagram := exploreRoomToHallway(d)
	newDiagram = append(newDiagram, exploreHallwayToRoom(d)...)

	return newDiagram
}

func exploreRoomToHallway(d Diagram) []Diagram {
	var newDiagrams []Diagram

	for letter := 0; letter < len(ROOMS_LETTER); letter++ {
		foundTheFirstLetter := false

		for i := 0; i < ROOM_LENGTH; i++ {
			if foundTheFirstLetter || d.runeAt(false, letter, i) == '.' {
				continue
			}

			foundTheFirstLetter = true

			allFollowingLetterAlreadyOk := true
			for j := i; j < ROOM_LENGTH; j++ {
				if d.runeAt(false, letter, j) != ROOMS_LETTER[letter] {
					allFollowingLetterAlreadyOk = false
				}
			}

			if !allFollowingLetterAlreadyOk || d.runeAt(false, letter, i) != ROOMS_LETTER[letter] {
				for pos := 0; pos < OPEN_SPACE_LENGTH; pos++ {
					if ok, distance := d.canMoveOnOpenSpace(ROOMS_POS[letter], pos, true); ok {
						diagram := d.duplicate()
						diagram.elements = replaceAtIndex(diagram.elements, d.runeAt(false, letter, i), pos)
						diagram.elements = replaceAtIndex(diagram.elements, '.', OPEN_SPACE_LENGTH+(ROOM_LENGTH*letter)+i)
						diagram = diagram.addWeigth(d.runeAt(false, letter, i), distance+i+1)
						newDiagrams = append(newDiagrams, diagram)
					}
				}
			}

		}
	}

	return newDiagrams
}

func exploreHallwayToRoom(d Diagram) []Diagram {
	var newDiagrams []Diagram

	for pos := 0; pos < OPEN_SPACE_LENGTH; pos++ {
		c := d.runeAt(true, 0, pos)

		if c != '.' {
			var indexC int
			switch c {
			case 'A':
				indexC = 0
			case 'B':
				indexC = 1
			case 'C':
				indexC = 2
			case 'D':
				indexC = 3
			}

			if ok, distance := d.canMoveOnOpenSpace(pos, ROOMS_POS[indexC], false); ok {
				diagram := d.duplicate()
				diagram.elements = replaceAtIndex(diagram.elements, '.', pos)

				assigned := false
				for i := ROOM_LENGTH - 1; i >= 0; i-- {
					if assigned {
						continue
					}

					if diagram.runeAt(false, indexC, i) == '.' {
						assigned = true

						diagram.elements = replaceAtIndex(diagram.elements, c, OPEN_SPACE_LENGTH+(indexC*ROOM_LENGTH)+i)
						diagram = diagram.addWeigth(c, distance+i+1)
						newDiagrams = append(newDiagrams, diagram)
					}
				}

			}
		}
	}

	return newDiagrams
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
