package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction rune

const (
	NORTH Direction = 'N'
	EAST            = 'E'
	SOUTH           = 'S'
	WEST            = 'W'
)

type Path struct {
	x         int
	y         int
	cost      int
	direction Direction
}

type Kiosk struct {
	kiosk     []string
	explored  [][]int
	toExplore []Path
}

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	kiosk := Kiosk{}

	for scanner.Scan() {
		line := scanner.Text()

		kiosk.kiosk = append(kiosk.kiosk, line)
	}

	kiosk.display()
	kiosk.explore()
}

func (kiosk Kiosk) display() {
	for _, k := range kiosk.kiosk {
		fmt.Println(k)
	}
}

func (kiosk Kiosk) explore() {
	startX, startY := kiosk.findStart()
	kiosk.toExplore = []Path{
		{x: startX, y: startY, cost: 0, direction: EAST},
	}

	for len(kiosk.toExplore) > 0 {
		path := kiosk.cheapestPathToExplore()
		kiosk.explorePath(path)
	}

	panic("What the fuck")
}

func (kiosk *Kiosk) cheapestPathToExplore() Path {
	cheapestPath := kiosk.toExplore[0]
	cheapestIndex := 0

	for i, path := range kiosk.toExplore {
		if path.cost < cheapestPath.cost {
			cheapestPath = path
			cheapestIndex = i
		}
	}

	kiosk.toExplore = append(kiosk.toExplore[:cheapestIndex], kiosk.toExplore[cheapestIndex+1:]...)

	return cheapestPath
}

func (kiosk *Kiosk) explorePath(path Path) {
	kiosk.explored = append(kiosk.explored, []int{path.x, path.y})

	if kiosk.kiosk[path.y][path.x] == 'E' {
		fmt.Println("Found exit at", path.x, path.y, "with cost", path.cost)
		os.Exit(0)

		return
	}

	kiosk.explorePathToDirection(path, NORTH)
	kiosk.explorePathToDirection(path, EAST)
	kiosk.explorePathToDirection(path, SOUTH)
	kiosk.explorePathToDirection(path, WEST)
}

func (kiosk *Kiosk) explorePathToDirection(path Path, direction Direction) {
	directionX := 0
	directionY := 0

	switch direction {
	case NORTH:
		directionY = -1
	case EAST:
		directionX = 1
	case SOUTH:
		directionY = 1
	case WEST:
		directionX = -1
	}

	targetRune := kiosk.kiosk[path.y+directionY][path.x+directionX]

	if targetRune == '#' {
		return
	}

	targetX := path.x + directionX
	targetY := path.y + directionY
	newPath := Path{
		x:         targetX,
		y:         targetY,
		cost:      path.cost + 1,
		direction: direction,
	}

	if kiosk.isAlreadyExploredPath(newPath) {
		return
	}

	if path.direction != direction {
		newPath.cost += 1000
	}

	kiosk.toExplore = append(kiosk.toExplore, newPath)
}

func (kiosk *Kiosk) isAlreadyExploredPath(path Path) bool {
	for _, p := range kiosk.explored {
		if p[0] == path.x && p[1] == path.y {
			return true
		}
	}

	return false
}

func (kiosk *Kiosk) findStart() (int, int) {
	for i, k := range kiosk.kiosk {
		for j, c := range k {
			if c == 'S' {
				return j, i
			}
		}
	}

	panic("Start not found")
}
