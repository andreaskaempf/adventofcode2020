# Advent of Code 2020

My solutions for the Advent of Code 2021 (done mostly in 2022)

* **Day 1** (Python): Read through a list of numbers, and find two
  (then three) that add up to 2020 (*easy*)
* **Day 2** (Python): Read each line, and validate password against 
  character-count rules (*easy*)
* **Day 3** (Python): Count up the number of '#' encountered when traversing
  an array with a given dx/dy slope for each step. The rows wrap if x 
  coordinate is wider than the data (*easy*).
* **Day 4** (Python): Simple validation of semi-structured text "passports",
  first looking for missing fields, then checking field contents against some
  simple rules (*easy*).
* **Day 5** (Python): Apply a sequence of binary forward/back operators to find
  row and seat in an airplane, and compute a code for these. For Part 1, find
  the maximum seat ID, for Part 2 the missing one.
* **Day 6** (Python): Read a file containing a sequence of letters on each row,
  each letter being a customer's "yes" response to a survey question. Customers
  are in groups, separated by newlines.  Part 1: For each group, count the
  number of questions to which anyone answered "yes". What is the sum of those
  counts?  Part 2: For each group, count the number of questions to which
  everyone answered "yes". What is the sum of those counts? (*easy*)
* **Day 7** (Python): Recursively navigate a tree upwards (Part 1) and
  downwards (Part 2), to count how many bags contain a given type of bag, and
  how many bags a given type of bag ultimately contains. (*easy*)
* **Day 8** (Python): Simulate execute of some simple bytecode, and detect when
  an instruction is executed the second time. For Part 2, find which nop or jmp
  needs to be changed to the opposite instruction, in order to fix the infinite
  loop. (*easy*)

To compile and run a Go program
* Change into the directory with the program
* go build day02.go
* ./day02

To run a Julia program
* Change into the directory with the program
* julia day02.jl

And of course, to run a Python program
* Change into the directory with the program
* python day06.py

AK, Dec 2021-2022
