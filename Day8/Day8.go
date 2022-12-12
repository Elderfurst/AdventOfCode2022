package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	treeMap := ReadInput()

	partOne(treeMap)
	partTwo(treeMap)
}

func ReadInput() [][]int {
	// Open our input file
	file, _ := os.Open("Day8/Day8.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	treeMap := make([][]int, 0)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		treeRow := make([]int, 0)

		for _, letter := range text {
			parsed, _ := strconv.Atoi(string(letter))

			treeRow = append(treeRow, parsed)
		}

		treeMap = append(treeMap, treeRow)
	}

	return treeMap
}

func partOne(treeMap [][]int) {
	numTrees := len(treeMap) + len(treeMap) + len(treeMap[0]) + len(treeMap[0]) - 4

	for i := 1; i < len(treeMap)-1; i++ {
		for j := 1; j < len(treeMap[i])-1; j++ {
			treeValue := treeMap[i][j]

			seeNorth := true
			for n := 0; n < i; n++ {
				north := treeMap[n][j]

				if north >= treeValue {
					seeNorth = false
					break
				}
			}

			seeEast := true
			for e := j + 1; e < len(treeMap[i]); e++ {
				east := treeMap[i][e]

				if east >= treeValue {
					seeEast = false
					break
				}
			}

			seeSouth := true
			for s := i + 1; s < len(treeMap); s++ {
				south := treeMap[s][j]

				if south >= treeValue {
					seeSouth = false
					break
				}
			}

			seeWest := true
			for w := 0; w < j; w++ {
				west := treeMap[i][w]

				if west >= treeValue {
					seeWest = false
					break
				}
			}

			if seeNorth || seeEast || seeSouth || seeWest {
				numTrees++
			}
		}
	}

	fmt.Println(numTrees)
}

func partTwo(treeMap [][]int) {
	bestScore := 0

	for i := 1; i < len(treeMap)-1; i++ {
		for j := 1; j < len(treeMap[i])-1; j++ {
			scenicScore := 0

			treeValue := treeMap[i][j]

			northScore := 0
			for n := i - 1; n >= 0; n-- {
				north := treeMap[n][j]

				northScore++

				if north >= treeValue {
					break
				}
			}

			eastScore := 0
			for e := j + 1; e < len(treeMap[i]); e++ {
				east := treeMap[i][e]

				eastScore++

				if east >= treeValue {
					break
				}
			}

			southScore := 0
			for s := i + 1; s < len(treeMap); s++ {
				south := treeMap[s][j]

				southScore++

				if south >= treeValue {
					break
				}
			}

			westScore := 0
			for w := j - 1; w >= 0; w-- {
				west := treeMap[i][w]

				westScore++

				if west >= treeValue {
					break
				}
			}

			scenicScore = northScore * eastScore * southScore * westScore

			if scenicScore > bestScore {
				bestScore = scenicScore
			}
		}
	}

	fmt.Println(bestScore)
}
