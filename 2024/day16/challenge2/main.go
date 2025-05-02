package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const DISPLAY_RESULT = false

type Direction rune
type DirectionRune rune

const (
	NORTH Direction = 'N'
	EAST            = 'E'
	SOUTH           = 'S'
	WEST            = 'W'
)

var DirectionToRune = map[Direction]DirectionRune{
	NORTH: '^',
	EAST:  '>',
	SOUTH: 'v',
	WEST:  '<',
}

type Path struct {
	x         int
	y         int
	cost      int
	direction Direction
	steps     []Step
}

type Step struct {
	x         int
	y         int
	direction Direction
}

type Kiosk struct {
	kiosk     [][]KioskPoint
	toExplore []Path
	cost      int
	finals    []Path
}

type Best struct {
	cost  int
	paths []Path
}

type KioskPoint struct {
	value rune
	bests map[Direction]Best
}

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	kiosk := Kiosk{
		cost: -1,
	}

	for scanner.Scan() {
		line := scanner.Text()
		kioskLine := []KioskPoint{}

		for _, c := range line {
			kioskLine = append(kioskLine, KioskPoint{
				value: c,
				bests: map[Direction]Best{
					NORTH: {cost: -1, paths: []Path{}},
					EAST:  {cost: -1, paths: []Path{}},
					SOUTH: {cost: -1, paths: []Path{}},
					WEST:  {cost: -1, paths: []Path{}},
				},
			})
		}

		kiosk.kiosk = append(kiosk.kiosk, kioskLine)
	}

	kiosk.display()
	kiosk.explore()

	fmt.Println("Found", len(kiosk.finals), "final paths", "with cost", kiosk.cost)

	if DISPLAY_RESULT {
		for _, path := range kiosk.finals {
			fmt.Println("Found exit at", path.x, path.y, "with cost", path.cost)

			for _, p := range path.steps {
				stepPoint := kiosk.kiosk[p.y][p.x]
				kiosk.displayWithDirection(p.x, p.y, p.direction)
				fmt.Println("Path", p.x, p.y, "direction", p.direction, "bestPaths", len(stepPoint.bests[p.direction].paths), "bestCost", stepPoint.bests[p.direction].cost)

				fmt.Scanf("%s", nil)
			}
		}
	}

	steps := kiosk.computeFinals(kiosk.finals)

	kiosk.displayWithSteps(steps)
}

func (kiosk Kiosk) computeFinals(paths []Path) []Step {
	pathToExplore := paths
	stepsVisited := map[string]bool{}

	steps := []Step{}

	for len(pathToExplore) > 0 {
		path := pathToExplore[0]
		pathToExplore = pathToExplore[1:]

		for _, step := range path.steps {
			stepKey := fmt.Sprintf("%d,%d,%c", step.x, step.y, step.direction)

			if stepsVisited[stepKey] {
				continue
			}

			stepsVisited[stepKey] = true
			steps = append(steps, Step{
				x:         step.x,
				y:         step.y,
				direction: step.direction,
			})
			pathToExplore = append(pathToExplore, kiosk.kiosk[step.y][step.x].bests[step.direction].paths...)
		}
	}

	return steps
}

func hasUniqueStep(steps []Step, x, y int, d Direction, skipDirection bool) bool {
	for _, step := range steps {
		if step.x == x && step.y == y {
			if skipDirection {
				return true
			}

			return step.direction == d
		}
	}

	return false
}

func (kiosk Kiosk) display() {
	kiosk.displayWithDirection(0, 0, EAST)
}

func (kiosk Kiosk) displayWithSteps(steps []Step) {
	counter := 0
	for y, k := range kiosk.kiosk {
		for x, point := range k {
			if hasUniqueStep(steps, x, y, NORTH, true) {
				fmt.Print("0")
				counter++
			} else {
				fmt.Print(string(point.value))
			}
		}
		fmt.Println()
	}

	fmt.Println("Found", counter, "unique points")
}

func (kiosk Kiosk) displayWithDirection(dX, dY int, direction Direction) {
	for y, k := range kiosk.kiosk {
		for x, point := range k {
			if x == dX && y == dY {
				fmt.Print(string(DirectionToRune[direction]))
			} else {
				fmt.Print(string(point.value))
			}
		}

		fmt.Println()
	}
}

func (kiosk *Kiosk) explore() {
	startX, startY := kiosk.findStart()
	kiosk.toExplore = []Path{
		{x: startX, y: startY, cost: 0, direction: EAST},
	}

	fmt.Println("Starting exploration at", startX, startY)

	for len(kiosk.toExplore) > 0 {
		path := kiosk.cheapestPathToExplore()

		if kiosk.cost != -1 && path.cost > kiosk.cost {
			fmt.Println("Ending exploration")

			return
		}

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
	if shouldExplore := kiosk.shouldExplorePath(path); !shouldExplore {
		return
	}

	if kiosk.kiosk[path.y][path.x].value == 'E' {
		path.steps = append(path.steps, Step{
			x:         path.x,
			y:         path.y,
			direction: path.direction,
		})

		kiosk.finals = append(kiosk.finals, path)
		kiosk.cost = path.cost
		fmt.Println("Found exit at", path.x, path.y, "with cost", path.cost)

		return
	}

	kiosk.kiosk[path.y][path.x].bests[path.direction] = Best{
		cost:  path.cost,
		paths: append(kiosk.kiosk[path.y][path.x].bests[path.direction].paths, path),
	}

	kiosk.explorePathToDirection(path, NORTH)
	kiosk.explorePathToDirection(path, EAST)
	kiosk.explorePathToDirection(path, SOUTH)
	kiosk.explorePathToDirection(path, WEST)
}

func (kiosk *Kiosk) shouldExplorePath(path Path) bool {
	best := kiosk.kiosk[path.y][path.x].bests[path.direction]

	if best.cost == -1 {
		return true
	}

	if path.cost < best.cost {
		return true
	}

	if path.cost == best.cost {
		fmt.Println("Found equal cost path at", path.x, path.y, "with cost", path.cost)
		path.steps = append(path.steps, Step{
			x:         path.x,
			y:         path.y,
			direction: path.direction,
		})

		kiosk.kiosk[path.y][path.x].bests[path.direction] = Best{
			cost:  path.cost,
			paths: append(best.paths, path),
		}
	}

	return false
}

func (kiosk *Kiosk) explorePathToDirection(path Path, direction Direction) {
	newSteps := []Step{}

	for _, step := range path.steps {
		newSteps = append(newSteps, Step{
			x:         step.x,
			y:         step.y,
			direction: step.direction,
		})
	}

	newSteps = append(newSteps, Step{
		x:         path.x,
		y:         path.y,
		direction: path.direction,
	})

	if path.direction != direction {
		kiosk.toExplore = append(kiosk.toExplore, Path{
			x:         path.x,
			y:         path.y,
			cost:      path.cost + 1000,
			direction: direction,
			steps:     newSteps,
		})

		return
	}

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

	targetPoint := kiosk.kiosk[path.y+directionY][path.x+directionX]
	targetX := path.x + directionX
	targetY := path.y + directionY

	if targetPoint.value == '#' {
		return
	}

	newPath := Path{
		x:         targetX,
		y:         targetY,
		cost:      path.cost + 1,
		direction: direction,
		steps:     newSteps,
	}

	kiosk.toExplore = append(kiosk.toExplore, newPath)
}

func (kiosk *Kiosk) findStart() (int, int) {
	for i, k := range kiosk.kiosk {
		for j, c := range k {
			if c.value == 'S' {
				return j, i
			}
		}
	}

	panic("Start not found")
}
