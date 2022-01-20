// Advent of Code 2020, Day 12
//
// Simulate movement of a "ship" based on simple instructions.
//
// AK, 14/01/2022

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	// Read file and split into lines
	t, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(t), "\n")

	// Do part 1
	part1(lines)
}

// Part 1: interpret instructions as simple movement of the ship N/E/S/W, or
// changing direction left/right by x degrees, or move forward in current
// direction
func part1(lines []string) {

	// Initial position is 0,0, and ship starts by facing east
	var x, y int64     // +x is east (right), +y is up (up)
	var dir int64 = 90 // 0 = north, 90 = east, 180 = south, 270 = west

	// Process each instruction
	// N/S/E/W : move north/south/east/weset by given value
	// L/R : turn left/right the given number of degrees.
	// F : move forward by the given value in the current direction
	for _, l := range lines {

		// Skip blank lines
		if len(l) == 0 {
			continue
		}

		// Parse instruction
		inst := l[0]
		amt, _ := strconv.ParseInt(l[1:], 10, 64)

		// Execute instruction
		if inst == 'N' {
			y += amt
		} else if inst == 'S' {
			y -= amt
		} else if inst == 'E' {
			x += amt
		} else if inst == 'W' {
			x -= amt
		} else if inst == 'L' {
			dir -= amt
			for dir < 0 {
				dir += 360
			}
		} else if inst == 'R' {
			dir += amt
			for dir >= 360 {
				dir -= 360
			}
		} else if inst == 'F' {
			if dir == 0 { // north (up)
				y += amt
			} else if dir == 90 { // east (right)
				x += amt
			} else if dir == 180 { // south (down)
				y -= amt
			} else if dir == 270 { // west (left)
				x -= amt
			} else {
				fmt.Printf("** Invalid direction %d\n", dir)
			}
		}

		fmt.Printf("%c %d -> %d,%d (dir %d)\n", inst, amt, x, y, dir)
	}

	dist := abs(x) + abs(y)
	fmt.Printf("Part 1: ending distance = %d\n", dist)
}

// Simple absolute number
func abs(a int64) int64 {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
