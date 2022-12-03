package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rucksacks := ReadInput()

	partOne(rucksacks)
	partTwo(rucksacks)
}

type Rucksack struct {
	CompartmentOne string
	CompartmentTwo string
	FullRucksack   string
}

func ReadInput() []Rucksack {
	// Open our input file
	file, _ := os.Open("Day3/Day3.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	rucksacks := make([]Rucksack, 0)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		midpoint := len(text) / 2

		compartmentOne := text[:midpoint]
		compartmentTwo := text[midpoint:]

		rucksack := Rucksack{
			CompartmentOne: compartmentOne,
			CompartmentTwo: compartmentTwo,
			FullRucksack:   text,
		}

		rucksacks = append(rucksacks, rucksack)
	}

	return rucksacks
}

func partOne(rucksacks []Rucksack) {
	priority := 0
	priorities := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, rucksack := range rucksacks {
		matchingLetter := matchingLetter2(rucksack.CompartmentOne, rucksack.CompartmentTwo)

		currentPriority := strings.Index(priorities, matchingLetter) + 1

		priority = priority + currentPriority
	}

	fmt.Println(priority)
}

func matchingLetter2(string1, string2 string) string {
	for _, item := range string1 {
		for _, otherItem := range string2 {
			if item == otherItem {
				return string(item)
			}
		}
	}

	return ""
}

func partTwo(rucksacks []Rucksack) {
	priority := 0
	priorities := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 0; i < len(rucksacks); i = i + 3 {
		matchingLetter := matchingLetter3(rucksacks[i].FullRucksack, rucksacks[i+1].FullRucksack, rucksacks[i+2].FullRucksack)

		currentPriority := strings.Index(priorities, matchingLetter) + 1

		priority = priority + currentPriority
	}

	fmt.Println(priority)
}

func matchingLetter3(string1, string2, string3 string) string {
	for _, item := range string1 {
		for _, otherItem := range string2 {
			for _, thirdItem := range string3 {
				if item == otherItem && item == thirdItem {
					return string(item)
				}
			}
		}
	}

	return ""
}
