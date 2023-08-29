package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	monkeys := ReadInput()
	partOne(monkeys)

	monkeysTwo := ReadInput()
	partTwo(monkeysTwo)
}

type Monkey struct {
	ID             int
	Items          []int
	Operation      string
	OperationValue int
	TestValue      int
	TestSuccess    int
	TestFailure    int
	InspectCount   int
}

func ReadInput() map[int]*Monkey {
	// Open our input file
	file, _ := os.Open("Day11/Day11.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	monkeys := make(map[int]*Monkey, 0)
	var currentMonkey *Monkey

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		if text == "" {
			continue
		}

		trimmed := strings.Trim(text, " ")

		split := strings.Split(trimmed, " ")

		switch split[0] {
		case "Monkey":
			trimmedID := strings.Trim(split[1], ":")

			monkeyID, _ := strconv.Atoi(trimmedID)

			currentMonkey = &Monkey{
				ID:    monkeyID,
				Items: make([]int, 0),
			}
		case "Starting":
			items := split[2:]

			for _, item := range items {
				trimmedItem := strings.Trim(item, ",")
				parsedItem, _ := strconv.Atoi(trimmedItem)
				currentMonkey.Items = append(currentMonkey.Items, parsedItem)
			}
		case "Operation:":
			operation := split[4]
			value, _ := strconv.Atoi(split[5])

			currentMonkey.Operation = operation
			currentMonkey.OperationValue = value
		case "Test:":
			test := split[3]

			parsedTest, _ := strconv.Atoi(test)

			currentMonkey.TestValue = parsedTest
		case "If":
			recipient := split[5]

			parsedRecipient, _ := strconv.Atoi(recipient)

			switch split[1] {
			case "true:":
				currentMonkey.TestSuccess = parsedRecipient
			case "false:":
				currentMonkey.TestFailure = parsedRecipient

				monkeys[currentMonkey.ID] = currentMonkey
			}
		}
	}

	return monkeys
}

func partOne(monkeys map[int]*Monkey) {
	monkeyKeys := make([]int, 0, len(monkeys))
	monkeyList := make([]*Monkey, 0)

	for monkeyID, monkey := range monkeys {
		monkeyKeys = append(monkeyKeys, monkeyID)
		monkeyList = append(monkeyList, monkey)
	}

	sort.Ints(monkeyKeys)

	for i := 0; i < 20; i++ {
		for _, key := range monkeyKeys {
			monkey := monkeys[key]

			for _, item := range monkey.Items {
				var worry int

				switch monkey.Operation {
				case "*":
					if monkey.OperationValue == 0 {
						worry = item * item
					} else {
						worry = item * monkey.OperationValue
					}
				case "+":
					if monkey.OperationValue == 0 {
						worry = item + item
					} else {
						worry = item + monkey.OperationValue
					}
				}

				downgradedWorry := worry / 3

				testResult := downgradedWorry%monkey.TestValue == 0

				var recipient *Monkey

				if testResult {
					recipient = monkeys[monkey.TestSuccess]
				} else {
					recipient = monkeys[monkey.TestFailure]
				}

				recipient.Items = append(recipient.Items, downgradedWorry)

				monkey.InspectCount++
			}

			monkey.Items = make([]int, 0)
		}
	}

	sort.Slice(monkeyList, func(i, j int) bool {
		return monkeyList[i].InspectCount > monkeyList[j].InspectCount
	})

	first := monkeyList[0].InspectCount
	second := monkeyList[1].InspectCount

	fmt.Println(first * second)
}

func partTwo(monkeys map[int]*Monkey) {
	monkeyKeys := make([]int, 0, len(monkeys))
	monkeyList := make([]*Monkey, 0)

	divisor := 1

	for monkeyID, monkey := range monkeys {
		monkeyKeys = append(monkeyKeys, monkeyID)
		monkeyList = append(monkeyList, monkey)
		divisor = divisor * monkey.TestValue
	}

	sort.Ints(monkeyKeys)

	for i := 0; i < 10000; i++ {
		for _, key := range monkeyKeys {
			monkey := monkeys[key]

			for _, item := range monkey.Items {
				var worry int

				switch monkey.Operation {
				case "*":
					if monkey.OperationValue == 0 {
						worry = item * item
					} else {
						worry = item * monkey.OperationValue
					}
				case "+":
					if monkey.OperationValue == 0 {
						worry = item + item
					} else {
						worry = item + monkey.OperationValue
					}
				}

				downgradedWorry := worry % divisor

				testResult := downgradedWorry%monkey.TestValue == 0

				var recipient *Monkey

				if testResult {
					recipient = monkeys[monkey.TestSuccess]
				} else {
					recipient = monkeys[monkey.TestFailure]
				}

				recipient.Items = append(recipient.Items, downgradedWorry)

				monkey.InspectCount++
			}

			monkey.Items = make([]int, 0)
		}
	}

	sort.Slice(monkeyList, func(i, j int) bool {
		return monkeyList[i].InspectCount > monkeyList[j].InspectCount
	})

	first := monkeyList[0].InspectCount
	second := monkeyList[1].InspectCount

	fmt.Println(first * second)
}
