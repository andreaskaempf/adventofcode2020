package main

import (
	"fmt"
	"os"
	"strings"
)

// A point is a coordinate in a hexagonal grid
type point struct {
	x, y int
}

func main() {

	// Read input file
	fname := "sample.txt"
	fname = "input.txt"
	data, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	// Create a map of coordinates
	coords := map[point]int{}

	// Process each line
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {

		if len(line) == 0 {
			continue
		}

		// Move the point according tot he instructions
		fmt.Println(line)
		insts := parseLine(line)
		p := point{0, 0}
		for _, inst := range insts {
			p = move(p, inst)
		}

		// Toggle the final position
		if coords[p] == 0 {
			coords[p] = 1
		} else {
			coords[p] = 0
		}
	}

	// Count up the flipped tiles
	count := 0
	for _, v := range coords {
		count += v
	}
	fmt.Println("Part 1:", count)
}

// Convert a line of instructins to a list of directions
func parseLine(line string) []string {
	var directions []string
	for i := 0; i < len(line); i++ {
		if line[i] == 'e' || line[i] == 'w' {
			directions = append(directions, line[i:i+1])
		} else {
			directions = append(directions, line[i:i+2])
			i++
		}
	}
	return directions
}

// Move a point in a hexagonal grid
func move(p point, dir string) point {
	switch dir {
	case "e":
		p.x++
	case "w":
		p.x--
	case "ne":
		p.x++
		p.y--
	case "nw":
		p.y--
	case "se":
		p.y++
	case "sw":
		p.x--
		p.y++
	}
	return p
}
