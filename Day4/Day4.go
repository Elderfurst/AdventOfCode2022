package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	elfSets := ReadInput()

	partOne(elfSets)
	partTwo(elfSets)
}

type ElfSet struct {
	ElfOne Elf
	ElfTwo Elf
}

type Elf struct {
	Start int
	End   int
}

func ReadInput() []ElfSet {
	// Open our input file
	file, _ := os.Open("Day4/Day4.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	elfSets := make([]ElfSet, 0)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		elfSplit := strings.Split(text, ",")

		elfOne := strings.Split(elfSplit[0], "-")
		elfTwo := strings.Split(elfSplit[1], "-")

		elfOneStart, _ := strconv.Atoi(elfOne[0])
		elfOneEnd, _ := strconv.Atoi(elfOne[1])
		elfTwoStart, _ := strconv.Atoi(elfTwo[0])
		elfTwoEnd, _ := strconv.Atoi(elfTwo[1])

		set := ElfSet{
			ElfOne: Elf{
				Start: elfOneStart,
				End:   elfOneEnd,
			},
			ElfTwo: Elf{
				Start: elfTwoStart,
				End:   elfTwoEnd,
			},
		}

		elfSets = append(elfSets, set)
	}

	return elfSets
}

func partOne(elfSets []ElfSet) {
	counter := 0

	for _, elfSet := range elfSets {
		contains := containsSet(elfSet)

		if contains {
			counter = counter + 1
		}
	}

	fmt.Println(counter)
}

func partTwo(elfSets []ElfSet) {
	counter := 0

	for _, elfSet := range elfSets {
		overlap := setOverlap(elfSet)

		if overlap {
			counter = counter + 1
		}
	}

	fmt.Println(counter)
}

func abs(elf Elf) int {
	result := elf.Start - elf.End

	if result < 0 {
		return result * -1
	}

	return result
}

func containsSet(elfSet ElfSet) bool {
	elfOneAbs := abs(elfSet.ElfOne)
	elfTwoAbs := abs(elfSet.ElfTwo)

	var bigger Elf
	var smaller Elf

	if elfOneAbs > elfTwoAbs {
		bigger = elfSet.ElfOne
		smaller = elfSet.ElfTwo
	} else {
		bigger = elfSet.ElfTwo
		smaller = elfSet.ElfOne
	}

	if bigger.Start <= smaller.Start &&
		bigger.End >= smaller.End {
		return true
	}

	return false
}

func setOverlap(elfSet ElfSet) bool {
	var first Elf
	var second Elf

	if elfSet.ElfOne.Start <= elfSet.ElfTwo.Start {
		first = elfSet.ElfOne
		second = elfSet.ElfTwo
	} else {
		first = elfSet.ElfTwo
		second = elfSet.ElfOne
	}

	if second.Start <= first.End {
		return true
	}

	return false
}
