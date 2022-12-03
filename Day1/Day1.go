package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	// Open our input file
	file, _ := os.Open("Day1/Day1.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxElf := 0
	currentElf := 0

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		if text == "" {
			if currentElf > maxElf {
				maxElf = currentElf
			}

			currentElf = 0

			continue
		}

		// Convert our string line into an int
		calories, _ := strconv.Atoi(text)

		currentElf = currentElf + calories
	}

	fmt.Println(maxElf)
}

func partTwo() {
	// Open our input file
	file, _ := os.Open("Day1/Day1.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	elves := make([]int, 0)

	var currentElf = 0

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		if text == "" {
			elves = append(elves, currentElf)

			currentElf = 0

			continue
		}

		// Convert our string line into an int
		calories, _ := strconv.Atoi(text)

		currentElf = currentElf + calories
	}

	sort.Ints(elves)

	elvesLength := len(elves)

	fmt.Println(elves[elvesLength-1] + elves[elvesLength-2] + elves[elvesLength-3])
}
