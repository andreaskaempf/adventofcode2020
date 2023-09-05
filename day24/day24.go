// Advent of Code 2020, Day 24
//
// Part 1: Count the number of black tiles in a hexagonal grid
// Part 2: Simulate 100 days of flipping tiles in a hexagonal grid

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

	// Convert to string and split lines
	lines := strings.Split(string(data), "\n")

	// Part 1: simulate one day, ignorning neighbors
	// S/b 10 for sample, 266 for input
	part1 := simulateDays(1, lines, false)
	fmt.Println("Part 1:", part1)
}

func simulateDays(days int, lines []string, considerNeighbors bool) int {

	// Create an empty map of coordinates
	coords := map[point]int{} // empty map of coords

	// Do each day
	for day := 0; day < days; day++ {

		// Process each line of instructions
		for _, line := range lines {

			// Skip empty lines
			if len(line) == 0 {
				continue
			}

			// Start at 0,0 and move the point according to the instructions
			//fmt.Println(line)
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
	}

	// Count up the flipped tiles
	count := 0
	for _, v := range coords {
		count += v
	}
	return count
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

// Count the number of neighbors of a point that are black
func countNeighbors(coords map[point]int, p point) int {
	count := 0
	for _, dir := range []string{"e", "w", "ne", "nw", "se", "sw"} {
		n := move(p, dir)
		count += coords[n]
	}
	return count
}
