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
		point.display()

		points = append(points, point)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		for _, p := range points {
			p.move()
		}

		displayPoints(points)
	}

	result := countPointsByQuadrant(points)
	fmt.Println(result)
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

func displayPoints(points []*point) {
	space := []string{}
	for i := 0; i < SIZE_Y; i++ {
		space = append(space, strings.Repeat(".", SIZE_X))
	}

	for _, p := range points {
		space[p.y] = space[p.y][:p.x] + "#" + space[p.y][p.x+1:]
	}

	for _, line := range space {
		fmt.Println(line)
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

func countPointsByQuadrant(points []*point) int {
	quadrants := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
	}

	for _, p := range points {
		if p.x == SIZE_X/2 || p.y == SIZE_Y/2 {
			continue
		}

		if p.x < SIZE_X/2 && p.y < SIZE_Y/2 {
			quadrants[1]++
		}

		if p.x >= SIZE_X/2 && p.y < SIZE_Y/2 {
			quadrants[2]++
		}

		if p.x < SIZE_X/2 && p.y >= SIZE_Y/2 {
			quadrants[3]++
		}

		if p.x >= SIZE_X/2 && p.y >= SIZE_Y/2 {
			quadrants[4]++
		}
	}

	fmt.Println(quadrants)

	return quadrants[1] * quadrants[2] * quadrants[3] * quadrants[4]
}
