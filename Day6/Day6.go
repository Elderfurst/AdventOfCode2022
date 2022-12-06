package main

import (
	"fmt"
	"os"
)

func main() {
	data := ReadInput()

	partOne(data)
	partTwo(data)
}

func ReadInput() string {
	content, _ := os.ReadFile("Day6/Day6.txt")

	return string(content)
}

func partOne(data string) {
	fmt.Println(findMarker(data, 4))
}

func partTwo(data string) {
	fmt.Println(findMarker(data, 14))
}

func findMarker(data string, length int) int {
	for i := range data {
		window := data[i : i+length]

		if !duplicatesExist(window) {
			return i + length
		}
	}

	return 0
}

func duplicatesExist(input string) bool {
	dict := make(map[string]bool, 0)

	for _, character := range input {
		dict[string(character)] = true
	}

	return len(dict) != len(input)
}
