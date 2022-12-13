package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	commands := ReadInput()

	partOne(commands)
	partTwo(commands)
}

type Command struct {
	Action string
	Count  int
}

func ReadInput() []Command {
	// Open our input file
	file, _ := os.Open("Day10/Day10.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	commands := make([]Command, 0)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		split := strings.Split(text, " ")

		if len(split) == 1 {
			command := Command{
				Action: split[0],
				Count:  0,
			}

			commands = append(commands, command)
		} else {
			parsedCount, _ := strconv.Atoi(split[1])

			command := Command{
				Action: split[0],
				Count:  parsedCount,
			}

			commands = append(commands, command)
		}
	}

	return commands
}

func partOne(commands []Command) {
	cycles := make(map[int]int, 0)
	currentCycle := 0
	currentValue := 1

	cycles[currentCycle] = currentValue

	for _, command := range commands {
		switch command.Action {
		case "noop":
			currentCycle++
			cycles[currentCycle] = currentValue
		case "addx":
			for i := 0; i < 2; i++ {
				currentCycle++
				cycles[currentCycle] = currentValue
			}

			currentValue += command.Count
		}
	}

	totalStrength := 0

	for cycle, value := range cycles {
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			totalStrength += cycle * value
		}
	}

	fmt.Println(totalStrength)
}

func partTwo(commands []Command) {
	cycles := make(map[int]int, 0)
	currentCycle := 0
	currentValue := 1

	cycles[currentCycle] = currentValue

	for _, command := range commands {
		switch command.Action {
		case "noop":
			printCycle(currentCycle, currentValue)

			currentCycle++
			cycles[currentCycle] = currentValue

		case "addx":
			for i := 0; i < 2; i++ {
				printCycle(currentCycle, currentValue)

				currentCycle++
				cycles[currentCycle] = currentValue
			}

			currentValue += command.Count
		}
	}
}

func printCycle(currentCycle, currentValue int) {
	value := currentCycle % 40

	if value >= currentValue-1 && value <= currentValue+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if currentCycle > 0 && value == 39 {
		fmt.Println()
	}
}
