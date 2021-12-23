package main

import (
	"fmt"
	"math"
	"time"
)

const FILE_LOCATION = "./input.txt"
const DEBUG = false

// Diagram struct

const ROOM_A = 2
const ROOM_B = 4
const ROOM_C = 6
const ROOM_D = 8

type Diagram struct {
	slotA1 rune
	slotA2 rune
	slotB1 rune
	slotB2 rune
	slotC1 rune
	slotC2 rune
	slotD1 rune
	slotD2 rune

	openSpace []rune

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
	if d.slotA1 != other.slotA1 {
		return false
	}

	if d.slotA2 != other.slotA2 {
		return false
	}

	if d.slotB1 != other.slotB1 {
		return false
	}

	if d.slotB2 != other.slotB2 {
		return false
	}

	if d.slotC1 != other.slotC1 {
		return false
	}

	if d.slotC2 != other.slotC2 {
		return false
	}

	if d.slotD1 != other.slotD1 {
		return false
	}

	if d.slotD2 != other.slotD2 {
		return false
	}

	for pos := range d.openSpace {
		if d.openSpace[pos] != other.openSpace[pos] {
			return false
		}
	}

	return true
}

func (diagram Diagram) isComplete() bool {
	return diagram.slotA1 == 'A' && diagram.slotA2 == 'A' &&
		diagram.slotB1 == 'B' && diagram.slotB2 == 'B' &&
		diagram.slotC1 == 'C' && diagram.slotC2 == 'C' &&
		diagram.slotD1 == 'D' && diagram.slotD2 == 'D'
}

func (diagram Diagram) display() {
	fmt.Printf("## %d ########\n", diagram.weight)
	fmt.Println("#############")

	fmt.Printf("#")
	for _, c := range diagram.openSpace {
		fmt.Printf("%s", string(c))
	}
	fmt.Println("#")

	fmt.Printf("###%s#%s#%s#%s###\n", string(diagram.slotA1), string(diagram.slotB1), string(diagram.slotC1), string(diagram.slotD1))
	fmt.Printf("  #%s#%s#%s#%s#\n", string(diagram.slotA2), string(diagram.slotB2), string(diagram.slotC2), string(diagram.slotD2))
	fmt.Println("  #########")
}

func (d Diagram) duplicate() Diagram {
	diagram := Diagram{
		slotA1: d.slotA1,
		slotA2: d.slotA2,
		slotB1: d.slotB1,
		slotB2: d.slotB2,
		slotC1: d.slotC1,
		slotC2: d.slotC2,
		slotD1: d.slotD1,
		slotD2: d.slotD2,

		weight:   d.weight,
		explored: d.explored,
	}

	diagram.openSpace = append(diagram.openSpace, d.openSpace...)

	return diagram
}

func (d Diagram) canMoveOnOpenSpace(start int, end int, checkStayAtEnd bool) (bool, int) {
	if checkStayAtEnd && (end == ROOM_A || end == ROOM_B || end == ROOM_C || end == ROOM_D) {
		return false, 0
	}

	if d.openSpace[end] != '.' {
		return false, 0
	}

	if start > end {
		temp := end
		end = start
		start = temp
	}

	for i := start + 1; i <= end-1; i++ {
		if d.openSpace[i] != '.' {
			return false, 0
		}
	}

	return true, int(math.Abs(float64(start - end)))
}

// Start

func main() {
	var diagrams []Diagram = []Diagram{
		{
			// Example
			//slotA1:    'B',
			//slotA2:    'A',
			//slotB1:    'C',
			//slotB2:    'D',
			//slotC1:    'B',
			//slotC2:    'C',
			//slotD1:    'D',
			//slotD2:    'A',
			//openSpace: []rune{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},

			// Simple test
			//slotA1:    '.',
			//slotA2:    'A',
			//slotB1:    'B',
			//slotB2:    'B',
			//slotC1:    'C',
			//slotC2:    'C',
			//slotD1:    'D',
			//slotD2:    'A',
			//openSpace: []rune{'.', '.', '.', '.', '.', 'D', '.', '.', '.', '.', '.'},

			// Input
			slotA1:    'C',
			slotA2:    'D',
			slotB1:    'A',
			slotB2:    'C',
			slotC1:    'B',
			slotC2:    'A',
			slotD1:    'D',
			slotD2:    'B',
			openSpace: []rune{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},

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

			fmt.Printf("Minimum wieght is %d\n", diagrams[smallestPos].weight)
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

	// Slot 1
	if d.slotA1 != '.' && (d.slotA1 != 'A' || d.slotA2 != 'A') {
		for pos := range d.openSpace {
			if ok, distance := d.canMoveOnOpenSpace(ROOM_A, pos, true); ok {
				diagram := d.duplicate()
				diagram.openSpace[pos] = d.slotA1
				diagram.slotA1 = '.'
				diagram = diagram.addWeigth(d.slotA1, distance+1)

				newDiagrams = append(newDiagrams, diagram)
			}
		}
	}

	if d.slotB1 != '.' && (d.slotB1 != 'B' || d.slotB2 != 'B') {
		for pos := range d.openSpace {
			if ok, distance := d.canMoveOnOpenSpace(ROOM_B, pos, true); ok {
				diagram := d.duplicate()
				diagram.openSpace[pos] = d.slotB1
				diagram.slotB1 = '.'
				diagram = diagram.addWeigth(d.slotB1, distance+1)

				newDiagrams = append(newDiagrams, diagram)
			}
		}
	}

	if d.slotC1 != '.' && (d.slotC1 != 'C' || d.slotC2 != 'C') {
		for pos := range d.openSpace {
			if ok, distance := d.canMoveOnOpenSpace(ROOM_C, pos, true); ok {
				diagram := d.duplicate()
				diagram.openSpace[pos] = d.slotC1
				diagram.slotC1 = '.'
				diagram = diagram.addWeigth(d.slotC1, distance+1)

				newDiagrams = append(newDiagrams, diagram)
			}
		}
	}

	if d.slotD1 != '.' && (d.slotD1 != 'D' || d.slotD2 != 'D') {
		for pos := range d.openSpace {
			if ok, distance := d.canMoveOnOpenSpace(ROOM_D, pos, true); ok {
				diagram := d.duplicate()
				diagram.openSpace[pos] = d.slotD1
				diagram.slotD1 = '.'
				diagram = diagram.addWeigth(d.slotD1, distance+1)

				newDiagrams = append(newDiagrams, diagram)
			}
		}
	}

	// Slot 2
	if d.slotA1 == '.' && d.slotA2 != '.' && d.slotA2 != 'A' {
		for pos := range d.openSpace {
			if ok, distance := d.canMoveOnOpenSpace(ROOM_A, pos, true); ok {
				diagram := d.duplicate()
				diagram.openSpace[pos] = d.slotA2
				diagram.slotA2 = '.'
				diagram = diagram.addWeigth(d.slotA2, distance+2)

				newDiagrams = append(newDiagrams, diagram)
			}
		}
	}

	if d.slotB1 == '.' && d.slotB2 != '.' && d.slotB2 != 'B' {
		for pos := range d.openSpace {
			if ok, distance := d.canMoveOnOpenSpace(ROOM_B, pos, true); ok {
				diagram := d.duplicate()
				diagram.openSpace[pos] = d.slotB2
				diagram.slotB2 = '.'
				diagram = diagram.addWeigth(d.slotB2, distance+2)

				newDiagrams = append(newDiagrams, diagram)
			}
		}
	}

	if d.slotC1 == '.' && d.slotC2 != '.' && d.slotC2 != 'C' {
		for pos := range d.openSpace {
			if ok, distance := d.canMoveOnOpenSpace(ROOM_C, pos, true); ok {
				diagram := d.duplicate()
				diagram.openSpace[pos] = d.slotC2
				diagram.slotC2 = '.'
				diagram = diagram.addWeigth(d.slotC2, distance+2)

				newDiagrams = append(newDiagrams, diagram)
			}
		}
	}

	if d.slotD1 == '.' && d.slotD2 != '.' && d.slotD2 != 'D' {
		for pos := range d.openSpace {
			if ok, distance := d.canMoveOnOpenSpace(ROOM_D, pos, true); ok {
				diagram := d.duplicate()
				diagram.openSpace[pos] = d.slotD2
				diagram.slotD2 = '.'
				diagram = diagram.addWeigth(d.slotD2, distance+2)

				newDiagrams = append(newDiagrams, diagram)
			}
		}
	}

	return newDiagrams
}

func exploreHallwayToRoom(d Diagram) []Diagram {
	var newDiagrams []Diagram

	for pos, c := range d.openSpace {
		if c != '.' {
			switch c {
			case 'A':
				if ok, distance := d.canMoveOnOpenSpace(pos, ROOM_A, false); ok {
					diagram := d.duplicate()
					diagram.openSpace[pos] = '.'

					if diagram.slotA2 == '.' {
						diagram.slotA2 = c
						diagram = diagram.addWeigth(c, distance+2)
						newDiagrams = append(newDiagrams, diagram)
					} else if diagram.slotA1 == '.' {
						diagram.slotA1 = c
						diagram = diagram.addWeigth(c, distance+1)
						newDiagrams = append(newDiagrams, diagram)
					}
				}
			case 'B':
				if ok, distance := d.canMoveOnOpenSpace(pos, ROOM_B, false); ok {
					diagram := d.duplicate()
					diagram.openSpace[pos] = '.'

					if diagram.slotB2 == '.' {
						diagram.slotB2 = c
						diagram = diagram.addWeigth(c, distance+2)
						newDiagrams = append(newDiagrams, diagram)
					} else if diagram.slotB1 == '.' {
						diagram.slotB1 = c
						diagram = diagram.addWeigth(c, distance+1)
						newDiagrams = append(newDiagrams, diagram)
					}
				}
			case 'C':
				if ok, distance := d.canMoveOnOpenSpace(pos, ROOM_C, false); ok {
					diagram := d.duplicate()
					diagram.openSpace[pos] = '.'

					if diagram.slotC2 == '.' {
						diagram.slotC2 = c
						diagram = diagram.addWeigth(c, distance+2)
						newDiagrams = append(newDiagrams, diagram)
					} else if diagram.slotC1 == '.' {
						diagram.slotC1 = c
						diagram = diagram.addWeigth(c, distance+1)
						newDiagrams = append(newDiagrams, diagram)
					}
				}
			case 'D':
				if ok, distance := d.canMoveOnOpenSpace(pos, ROOM_D, false); ok {
					diagram := d.duplicate()
					diagram.openSpace[pos] = '.'

					if diagram.slotD2 == '.' {
						diagram.slotD2 = c
						diagram = diagram.addWeigth(c, distance+2)
						newDiagrams = append(newDiagrams, diagram)
					} else if diagram.slotD1 == '.' {
						diagram.slotD1 = c
						diagram = diagram.addWeigth(c, distance+1)
						newDiagrams = append(newDiagrams, diagram)
					}
				}
			}
		}
	}

	return newDiagrams
}
