package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type tile struct {
	bottom bool `default:"false"`
	top bool `default:"false"`
	left bool `default:"false"`
	right bool `default:"false"`

	value string
}

type path struct {
	x int
	y int
	direction string
}

var contraception = [][]tile{}

func main() {
	args := os.Args
	file, err := os.Open(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		tiles := []tile{}

		for _, char := range line {
			tiles = append(tiles, tile{
				value: string(char),
			})
		}

		contraception = append(contraception, tiles)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, line := range contraception {
		for _, tile := range line {
			fmt.Print(tile.value)
		}

		fmt.Println()
	}

	explore()

	sum := 0
	for _, line := range contraception {
		for _, tile := range line {
			if tile.bottom || tile.top || tile.left || tile.right {
				fmt.Print("#")
				sum++
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}

	fmt.Println(sum)
}

func explore() {
	paths := []path{ { x: 0, y: 0, direction: "right" } }
	const MAX_ITERATIONS = 10000000000000
	iterations := 0

	for len(paths) > 0 && iterations < MAX_ITERATIONS {
		iterations++
		currentPath := paths[0]
		paths = paths[1:]

		if currentPath.x < 0 || currentPath.y < 0 || currentPath.x >= len(contraception) || currentPath.y >= len(contraception[0]) {
			continue
		}

		if alreadyExplored(currentPath) {
			continue
		}

		currentTile := contraception[currentPath.x][currentPath.y]
		switch currentTile.value {
		case "|":
			if currentPath.direction == "right" || currentPath.direction == "left" {
				paths = append(paths,
					path{ x: currentPath.x - 1, y: currentPath.y, direction: "top" },
					path{ x: currentPath.x + 1, y: currentPath.y, direction: "bottom" },
				)
			} else if currentPath.direction == "top" {
				paths = append(paths, path{ x: currentPath.x - 1, y: currentPath.y, direction: "top" })
			} else if currentPath.direction == "bottom" {
				paths = append(paths, path{ x: currentPath.x + 1, y: currentPath.y, direction: "bottom" })
			} else {
				panic("Unknown direction")
			}
		case "-":
			if currentPath.direction == "top" || currentPath.direction == "bottom" {
				paths = append(paths,
					path{ x: currentPath.x, y: currentPath.y - 1, direction: "left" },
					path{ x: currentPath.x, y: currentPath.y + 1, direction: "right" },
				)
			} else if currentPath.direction == "left" {
				paths = append(paths, path{ x: currentPath.x, y: currentPath.y - 1, direction: "left" })
			} else if currentPath.direction == "right" {
				paths = append(paths, path{ x: currentPath.x, y: currentPath.y + 1, direction: "right" })
			} else {
				panic("Unknown direction")
			}
		case ".":
			switch currentPath.direction {
			case "right":
				paths = append(paths, path{ x: currentPath.x, y: currentPath.y + 1, direction: "right" })
			case "left":
				paths = append(paths, path{ x: currentPath.x, y: currentPath.y - 1, direction: "left" })
			case "top":
				paths = append(paths, path{ x: currentPath.x - 1, y: currentPath.y, direction: "top" })
			case "bottom":
				paths = append(paths, path{ x: currentPath.x + 1, y: currentPath.y, direction: "bottom" })
			default:
				panic("Unknown direction")
			}
		case "\\":
			switch currentPath.direction {
			case "right":
				paths = append(paths, path{ x: currentPath.x + 1, y: currentPath.y, direction: "bottom" })
			case "left":
				paths = append(paths, path{ x: currentPath.x - 1, y: currentPath.y, direction: "top" })
			case "top":
				paths = append(paths, path{ x: currentPath.x, y: currentPath.y - 1, direction: "left" })
			case "bottom":
				paths = append(paths, path{ x: currentPath.x, y: currentPath.y + 1, direction: "right" })
			default:
				panic("Unknown direction")
			}
		case "/":
			switch currentPath.direction {
			case "right":
				paths = append(paths, path{ x: currentPath.x - 1, y: currentPath.y, direction: "top" })
			case "left":
				paths = append(paths, path{ x: currentPath.x + 1, y: currentPath.y, direction: "bottom" })
			case "top":
				paths = append(paths, path{ x: currentPath.x, y: currentPath.y + 1, direction: "right" })
			case "bottom":
				paths = append(paths, path{ x: currentPath.x, y: currentPath.y - 1, direction: "left" })
			default:
				panic("Unknown direction")
			}
		default:
			fmt.Println(currentTile.value)
			panic("Unknown tile")
		}

	}
}

func alreadyExplored(path path) bool {
	if path.direction == "right" && contraception[path.x][path.y].right {
		return true
	} else if path.direction == "right" { 
		contraception[path.x][path.y].right = true

		return false
	}

	if path.direction == "left" && contraception[path.x][path.y].left {
		return true
	} else if path.direction == "left" {
		contraception[path.x][path.y].left = true

		return false
	}

	if path.direction == "top" && contraception[path.x][path.y].top {
		return true
	} else if path.direction == "top" {
		contraception[path.x][path.y].top = true

		return false
	}

	if path.direction == "bottom" && contraception[path.x][path.y].bottom {
		return true
	} else if path.direction == "bottom" {
		contraception[path.x][path.y].bottom = true

		return false
	}

	panic("Unknown direction")
}
