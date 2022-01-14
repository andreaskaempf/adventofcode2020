// Part 1: 2237 too high

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Read each line of input file
	lines := [][]byte{}
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		l := []byte{}
		for i := 0; i < len(t); i++ {
			l = append(l, t[i])
		}
		lines = append(lines, l)
	}

	fmt.Println("Starting:")
	printBoard(lines)

	// Iterate until no more changes
	iter := 0
	for {

		// Make a copy
		lines1 := [][]byte{}
		for _, l := range lines {
			l1 := make([]byte, len(l), len(l))
			copy(l1, l)                 // does this work?
			lines1 = append(lines1, l1) // makes a copy?
		}

		// The following rules are applied to every seat simultaneously:
		// 1. If a seat is empty (L) and there are no occupied seats adjacent
		//    to it, the seat becomes occupied.
		// 2. If a seat is occupied (#) and four or more seats adjacent to it
		//    are also occupied, the seat becomes empty.
		// Otherwise, the seat's state does not change.
		changed := false
		for r := 0; r < len(lines); r++ {
			for c := 0; c < len(lines[r]); c++ {
				nOccup := adjacentOccupied(lines, r, c)
				if lines[r][c] == 'L' && nOccup == 0 {
					lines1[r][c] = '#'
					changed = true
				}
				if lines[r][c] == '#' && nOccup >= 4 {
					lines1[r][c] = 'L'
					changed = true
				}
			}
		}

		lines = lines1
		iter += 1
		fmt.Printf("\nIteration %d:\n", iter)
		printBoard(lines)

		if !changed { //}|| iter > 50 {
			break
		}
	}

}

// Count the  number of adjacent seats around a given seat that are occupied
func adjacentOccupied(lines [][]byte, r int, c int) int {
	result := 0
	for ri := r - 1; ri <= r+1; ri++ {
		for ci := c - 1; ci <= c+1; ci++ {
			if ri == r && ci == c {
				continue
			}
			if ri >= 0 && ri < len(lines) && ci >= 0 && ci < len(lines[ri]) && lines[ri][ci] == '#' {
				result += 1
			}
		}
	}
	return result
}

// Count the  number of seats on a board that are occupied
func occupied(lines [][]byte) int {
	result := 0
	for ri := 0; ri < len(lines); ri++ {
		for ci := 0; ci < len(lines[ri]); ci++ {
			if lines[ri][ci] == '#' {
				result += 1
			}
		}
	}
	return result
}

func printBoard(lines [][]byte) {
	for _, l := range lines {
		fmt.Println(string(l))
	}
	fmt.Printf("%d seats occupied\n", occupied(lines))
}
