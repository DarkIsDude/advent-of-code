package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// const SIZE_X = 11
// const SIZE_Y = 7

const SIZE_X = 101
const SIZE_Y = 103

type point struct {
	x  int
	y  int
	vX int
	vY int
}

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	points := []*point{}

	for scanner.Scan() {
		line := scanner.Text()
		regexp := regexp.MustCompile(`^p=(?<X>\d+),(?<Y>\d+) v=(?<vX>-?\d+),(?<vY>-?\d+)$`)
		match := regexp.FindStringSubmatch(line)
		point := &point{}
		point.init(match)
		points = append(points, point)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10000000; i++ {
		for _, p := range points {
			p.move()
		}

		fmt.Printf("After %d seconds:\n", i+1)
		countPointsWithNeighbours(points)
	}
}

func (p *point) init(init []string) {
	var err error

	p.x, err = strconv.Atoi(init[1])
	if err != nil {
		log.Fatal(err)
	}

	p.y, err = strconv.Atoi(init[2])
	if err != nil {
		log.Fatal(err)
	}

	p.vX, err = strconv.Atoi(init[3])
	if err != nil {
		log.Fatal(err)
	}

	p.vY, err = strconv.Atoi(init[4])
	if err != nil {
		log.Fatal(err)
	}
}

func (p *point) display() {
	fmt.Printf("x: %d, y: %d, vX: %d, vY: %d\n", p.x, p.y, p.vX, p.vY)
}

func displayPoints(points []*point, hide bool) {
	space := []string{}
	for i := 0; i < SIZE_Y; i++ {
		space = append(space, strings.Repeat(".", SIZE_X))
	}

	for _, p := range points {
		space[p.y] = space[p.y][:p.x] + "#" + space[p.y][p.x+1:]
	}

	for i, line := range space {
		if !hide {
			fmt.Println(line)
			continue
		}

		if i == SIZE_Y/2 {
			fmt.Println(strings.Repeat(" ", SIZE_X))
			continue
		}

		fmt.Println(line[:SIZE_X/2] + " " + line[SIZE_X/2+1:])
	}

	fmt.Println()
}

func (p *point) move() {
	p.x += p.vX
	p.y += p.vY

	if p.x < 0 {
		p.x = SIZE_X + p.x
	}

	if p.y < 0 {
		p.y = SIZE_Y + p.y
	}

	if p.x >= SIZE_X {
		p.x = p.x - SIZE_X
	}

	if p.y >= SIZE_Y {
		p.y = p.y - SIZE_Y
	}
}

func countPointsWithNeighbours(points []*point) {
	pointsWithNeighboutsCount := 0

	for _, p := range points {
		maxX := p.x + 1
		maxY := p.y + 1
		minX := p.x - 1
		minY := p.y - 1

		found := false

		for _, p2 := range points {
			if p2 == p {
				continue
			}

			if p2.x >= minX && p2.x <= maxX && p2.y >= minY && p2.y <= maxY {
				found = true
				break
			}
		}

		if found {
			pointsWithNeighboutsCount++
		}
	}

	fmt.Printf("Count: %d\n", pointsWithNeighboutsCount)

	if pointsWithNeighboutsCount > 350 {
		displayPoints(points, false)
		os.Exit(0)
	}
}
