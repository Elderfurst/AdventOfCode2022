package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stacks, procedures := ReadInput()

	// Since each part actually modifies 'stacks' it's easier to just comment out which part you don't want to run
	// instead of dealing with deep cloning the map to pass to two separate methods /shrug

	// partOne(stacks, procedures)
	partTwo(stacks, procedures)
}

type Procedure struct {
	MoveCount     int
	StartingStack int
	EndingStack   int
}

func ReadInput() (map[int][]string, []Procedure) {
	// Open our input file
	file, _ := os.Open("Day5/Day5.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	stacks := make(map[int][]string)
	procedures := make([]Procedure, 0)

	parseProcedures := false

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		if text == "" {
			parseProcedures = true
			continue
		}

		if parseProcedures {
			split := strings.Split(text, " ")

			moveCount, _ := strconv.Atoi(split[1])
			start, _ := strconv.Atoi(split[3])
			end, _ := strconv.Atoi(split[5])

			procedure := Procedure{
				MoveCount:     moveCount,
				StartingStack: start,
				EndingStack:   end,
			}

			procedures = append(procedures, procedure)
		} else {
			if strings.Contains(text, "1") {
				continue
			}

			for i, value := range text {
				parsed := string(value)

				if parsed != " " {
					stacks[i+1] = append(stacks[i+1], parsed)
				}
			}
		}
	}

	return stacks, procedures
}

func partOne(stacks map[int][]string, procedures []Procedure) {
	for _, procedure := range procedures {
		start := procedure.StartingStack
		end := procedure.EndingStack

		for i := 0; i < procedure.MoveCount; i++ {
			movingCrate := stacks[start][0]

			// prepend our crate since the 'top' of the stack is actually the beginning of the slice
			stacks[end] = append([]string{movingCrate}, stacks[end]...)

			// remove the crate after it's been moved
			stacks[start] = stacks[start][1:]
		}
	}

	for i := 1; i <= len(stacks); i++ {
		fmt.Print(stacks[i][0])
	}
}

func partTwo(stacks map[int][]string, procedures []Procedure) {
	for _, procedure := range procedures {
		start := procedure.StartingStack
		end := procedure.EndingStack

		for i := procedure.MoveCount; i > 0; i-- {
			position := i - 1

			movingCrate := stacks[start][position]

			// prepend our crate since the 'top' of the stack is actually the beginning of the slice
			stacks[end] = append([]string{movingCrate}, stacks[end]...)

			// remove the crate after it's been moved
			stacks[start] = append(stacks[start][:position], stacks[start][position+1:]...)
		}
	}

	for i := 1; i <= len(stacks); i++ {
		fmt.Print(stacks[i][0])
	}
}
