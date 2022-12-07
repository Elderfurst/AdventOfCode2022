package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	directories := ReadInput()

	partOne(directories)
	partTwo(directories)
}

type Directory struct {
	Name     string
	Size     int
	Parent   *Directory
	Children map[string]*Directory
}

func ReadInput() map[string]*Directory {
	// Open our input file
	file, _ := os.Open("Day7/Day7.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	directories := make(map[string]*Directory, 0)

	mainDirectory := &Directory{
		Name:     "/",
		Size:     0,
		Children: make(map[string]*Directory, 0),
	}

	directories["/"] = mainDirectory

	var currentDirectory *Directory

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		split := strings.Split(text, " ")

		if split[0] == "$" {
			if split[1] == "cd" {
				if split[2] == "/" {
					currentDirectory = directories["/"]
				} else if split[2] == ".." {
					currentDirectory = currentDirectory.Parent
				} else {
					currentDirectory = currentDirectory.Children[split[2]]
				}
			}
		} else {
			if split[0] == "dir" {
				_, found := currentDirectory.Children[split[1]]
				if !found {
					newDirectory := &Directory{
						Name:     split[1],
						Size:     0,
						Parent:   currentDirectory,
						Children: make(map[string]*Directory, 0),
					}

					fullName := buildFullDirectoryName(newDirectory)

					directories[fullName] = newDirectory
					currentDirectory.Children[newDirectory.Name] = newDirectory
				}
			} else {
				fileSize, _ := strconv.Atoi(split[0])

				updateSizes(currentDirectory, fileSize)
			}
		}
	}

	return directories
}

func partOne(directories map[string]*Directory) {
	totalSize := 0

	for _, directory := range directories {
		if directory.Size <= 100000 {
			totalSize += directory.Size
		}
	}

	fmt.Println(totalSize)
}

func buildFullDirectoryName(directory *Directory) string {
	name := directory.Name

	if directory.Parent != nil {
		name = buildFullDirectoryName(directory.Parent) + "/" + name
	}

	return name
}

func updateSizes(directory *Directory, size int) {
	directory.Size += size

	if directory.Parent != nil {
		updateSizes(directory.Parent, size)
	}
}

func partTwo(directories map[string]*Directory) {
	totalSpace := 70000000
	neededSpace := 30000000
	usedSpace := directories["/"].Size
	unusedSpace := totalSpace - usedSpace
	spaceToClear := neededSpace - unusedSpace

	spaceDifference := math.MaxInt
	finalSpace := 0

	for _, directory := range directories {
		if directory.Size > spaceToClear {
			tempDifference := directory.Size - spaceToClear
			if tempDifference < spaceDifference {
				spaceDifference = tempDifference
				finalSpace = directory.Size
			}
		}
	}

	fmt.Println(finalSpace)
}
