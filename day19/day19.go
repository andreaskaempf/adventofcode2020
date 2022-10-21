// Advent of Code 2020, Day 19
//
// AK, 18/10/2022

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// Structure for a rule, consists of either a single character,
// or one or two lists of numbers of other rules that must be matched
type Rule struct {
	num  int   // rule number
	char byte  // matching character
	L    []int // sub-rule numbers, left part
	R    []int // sub-rule numbers, right part
}

// Global lists of rules and messages
var rules map[int]Rule
var messages []string

func main() {

	// Read rules and messages from input file
	//readData("sample.txt")  // 2 matches (part 1)
	readData("sample2.txt") // 2/12 matches (part 1/2)
	//readData("input.txt") // 156/363

	// Expand rules and create regular expression
	rx := expandRules()

	// Process messages
	// Sample1 should be: t, f, t, f, f
	matches := 0
	for _, msg := range messages {
		ok := processMessage(msg, rx)
		//fmt.Println(msg, "=>", ok)
		if ok {
			matches++
		}
	}
	fmt.Println("Part 1:", matches, "messages matched")

	// For part 2, change two rules
	// 8: 42 | 42 8
	// 11: 42 31 | 42 11 31
	fmt.Println("Rules 8 and 11:", rules[8], rules[11])
	rules[8] = Rule{8, 0, []int{42}, []int{42, 8}}
	rules[11] = Rule{11, 0, []int{42, 31}, []int{42, 11, 31}}
	fmt.Println("  changed to", rules[8], rules[11])
	rx = expandRules()
	if rx != nil {
		matches = 0
		for _, msg := range messages {
			ok := processMessage(msg, rx)
			//fmt.Println(msg, "=>", ok)
			if ok {
				matches++
			}
		}
		fmt.Println("Part 2:", matches, "messages matched") // 156 for Part 1
	}
}

// Convert nested rules to a large regexp, by iteratively replacing until
// rule 0 is fully expanded. For example (from problem description):
//
//	0: 4 1 5
//	1: 2 3 | 3 2
//	2: 4 4 | 5 5
//	3: 4 5 | 5 4
//	4: "a"
//	5: "b"
//
// Gets expanded in this sequence:
// 4: "a"       -> a
// 5: "b"       -> b
// 2: 4 4 | 5 5 -> (aa|bb)
// 3 4 5 | 5 4  -> (ab|ba)
// 1: 2 3 | 3 2 -> (aa|bb)(ab|ba)|(ab|ba)(aa|bb)
// 0: 4 1 5     -> a(((aa|bb)(ab|ba))|((ab|ba)(aa|bb)))b
//
// So characters are replaced verbatim, but conditions are surrounded by
// parentheses after substitution.
func expandRules() *regexp.Regexp {

	// Map of sub-rules that have been expanded
	expanded := map[int]string{}

	// Expand characters first
	for k, r := range rules {
		if r.char > 0 {
			expanded[k] = string(r.char)
		}
	}

	// Go through remaining rules, expanding any where all sub-rules have
	// already been expanded
	changed := true
	for changed { // keep iterating until no more changes made

		changed = false
		for k, r := range rules {

			// Skip rules that have already been expanded
			if hasKey(expanded, k) {
				continue
			}

			// Check if all the sub-rules for this rule have been expanded,
			// skip it if not
			subrulesExpanded := true
			for _, n := range r.L {
				if !hasKey(expanded, n) {
					subrulesExpanded = false
					break
				}
			}
			for _, n := range r.R {
				if !hasKey(expanded, n) {
					subrulesExpanded = false
					break
				}
			}
			if !subrulesExpanded {
				continue
			}

			// Expand this rule
			rx := "("
			for _, i := range r.L {
				rx += expanded[i]
			}
			if len(r.R) > 0 {
				rx += "|"
				for _, i := range r.R {
					rx += expanded[i]
				}
			}
			rx += ")"
			expanded[k] = rx
			changed = true
			fmt.Println("Expanded", r, "=>", rx)
		}
	}

	// Compile rule zero into regular expression
	patt := expanded[0]
	if len(patt) == 0 {
		fmt.Println("Regex failed to expand")
		for n, r := range rules {
			x, ok := expanded[n]
			if !ok {
				x = "(MISSING)"
			}
			fmt.Println(n, ":", r, "=>", x)
		}
		return nil
	}

	//fmt.Println("Top regex:", patt)
	return regexp.MustCompile(patt)
}

// Check if dictionary has key
func hasKey(dict map[int]string, n int) bool {
	_, ok := dict[n]
	return ok
}

// Process a message, check if matches
func processMessage(msg string, rx *regexp.Regexp) bool {
	match := rx.FindString(msg)
	return len(match) == len(msg)
}

// Read data file into a list of rules and inputs
func readData(filename string) {

	// Read file into list of strings
	data, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(data), "\n")

	rules = map[int]Rule{}
	readingRules := true
	for _, l := range lines {

		// At first blank line, stop reading rules and start reading messages
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			readingRules = false
			continue
		} else if readingRules {
			r := parseRule(l)
			rules[r.num] = r
		} else {
			messages = append(messages, l)
		}
	}
}

// Parse a rule
// E.g., 99: "a"
// or 88: 24 103 | 36 6
func parseRule(s string) Rule {
	words := strings.Split(s, " ")
	n := atoi(words[0][:len(words[0])-1])
	r := Rule{num: n}
	parsingR := false
	for i := 1; i < len(words); i++ {
		w := words[i]
		if w[0] == '"' { // quote means rule is a character
			r.char = w[1]
		} else if words[i] == "|" { // bar means start of right sub-rules
			parsingR = true
		} else if parsingR { // right sub-rules, list of numbers
			r.R = append(r.R, atoi(w))
		} else { // left sub-rules, list of numbers
			r.L = append(r.L, atoi(w))
		}
	}
	return r
}

// Parse number
func atoi(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Could not parse, assuming zero:", s)
		n = 0
	}
	return int(n)
}
