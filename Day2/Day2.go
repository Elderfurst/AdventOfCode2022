package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	matches := ReadInput()

	partOne(matches)
	partTwo(matches)
}

type Match struct {
	OpponentHand string
	MyHand       string
}

func ReadInput() []Match {
	// Open our input file
	file, _ := os.Open("Day2/Day2.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	matches := make([]Match, 0)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		split := strings.Split(text, " ")

		match := Match{
			OpponentHand: split[0],
			MyHand:       split[1],
		}

		matches = append(matches, match)
	}

	return matches
}

func partOne(matches []Match) {
	myScore := 0

	for _, match := range matches {
		switch match.MyHand {
		case "X":
			myScore = myScore + 1
		case "Y":
			myScore = myScore + 2
		case "Z":
			myScore = myScore + 3
		}

		if (match.MyHand == "X" && match.OpponentHand == "A") ||
			(match.MyHand == "Y" && match.OpponentHand == "B") ||
			(match.MyHand == "Z" && match.OpponentHand == "C") {
			myScore = myScore + 3
		} else if (match.MyHand == "X" && match.OpponentHand == "C") ||
			(match.MyHand == "Y" && match.OpponentHand == "A") ||
			(match.MyHand == "Z" && match.OpponentHand == "B") {
			myScore = myScore + 6
		}
	}

	fmt.Println(myScore)
}

func partTwo(matches []Match) {
	myScore := 0

	for _, match := range matches {
		var actualHand string

		switch match.MyHand {
		case "X":
			if match.OpponentHand == "A" {
				actualHand = "Z"
			} else if match.OpponentHand == "B" {
				actualHand = "X"
			} else if match.OpponentHand == "C" {
				actualHand = "Y"
			}
		case "Y":
			myScore = myScore + 3
			if match.OpponentHand == "A" {
				actualHand = "X"
			} else if match.OpponentHand == "B" {
				actualHand = "Y"
			} else if match.OpponentHand == "C" {
				actualHand = "Z"
			}
		case "Z":
			myScore = myScore + 6
			if match.OpponentHand == "A" {
				actualHand = "Y"
			} else if match.OpponentHand == "B" {
				actualHand = "Z"
			} else if match.OpponentHand == "C" {
				actualHand = "X"
			}
		}

		switch actualHand {
		case "X":
			myScore = myScore + 1
		case "Y":
			myScore = myScore + 2
		case "Z":
			myScore = myScore + 3
		}
	}

	fmt.Println(myScore)
}
