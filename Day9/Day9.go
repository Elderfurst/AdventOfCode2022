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
	commands := ReadInput()

	partOne(commands)
	partTwo(commands)
}

type Command struct {
	Direction string
	Distance  int
}

func ReadInput() []Command {
	// Open our input file
	file, _ := os.Open("Day9/Day9.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	commands := make([]Command, 0)

	// scanner.Scan() advances to the next line by default
	for scanner.Scan() {
		// Get each individual line
		text := scanner.Text()

		split := strings.Split(text, " ")

		parsedDistance, _ := strconv.Atoi(split[1])

		command := Command{
			Direction: split[0],
			Distance:  parsedDistance,
		}

		commands = append(commands, command)
	}

	return commands
}

type Node struct {
	X int
	Y int
}

func partOne(commands []Command) {
	visited := make(map[string]bool, 0)

	head := Node{
		0, 0,
	}

	tail := Node{
		0, 0,
	}

	visited["0:0"] = true

	for _, command := range commands {
		for i := 0; i < command.Distance; i++ {
			switch command.Direction {
			case "U":
				head.Y++
			case "D":
				head.Y--
			case "L":
				head.X--
			case "R":
				head.X++
			}

			if head.X == tail.X && head.Y == tail.Y {
				continue
			}

			if head.X == tail.X {
				if head.Y-tail.Y > 1 {
					tail.Y = head.Y - 1
				} else if tail.Y-head.Y > 1 {
					tail.Y = head.Y + 1
				}
			} else if head.Y == tail.Y {
				if head.X-tail.X > 1 {
					tail.X = head.X - 1
				} else if tail.X-head.X > 1 {
					tail.X = head.X + 1
				}
			} else {
				distance := getDistance(&head, &tail)

				if distance >= 2 {
					switch command.Direction {
					case "U":
						tail.X = head.X
						tail.Y = head.Y - 1
					case "D":
						tail.X = head.X
						tail.Y = head.Y + 1
					case "L":
						tail.Y = head.Y
						tail.X = head.X + 1
					case "R":
						tail.Y = head.Y
						tail.X = head.X - 1
					}
				}
			}

			visited[tail.getKey()] = true
		}
	}

	fmt.Println(len(visited))
}

func partTwo(commands []Command) {
	visited := make(map[string]bool, 0)
	nodes := []*Node{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	head := nodes[0]

	visited["0:0"] = true

	for _, command := range commands {
		for i := 0; i < command.Distance; i++ {
			switch command.Direction {
			case "U":
				head.Y++
			case "D":
				head.Y--
			case "L":
				head.X--
			case "R":
				head.X++
			}

			for j := 1; j < len(nodes); j++ {
				node := nodes[j-1]
				nextNode := nodes[j]

				if node.X == nextNode.X && node.Y == nextNode.Y {
					continue
				}

				if node.X == nextNode.X {
					if node.Y-nextNode.Y > 1 {
						nextNode.Y = node.Y - 1
					} else if nextNode.Y-node.Y > 1 {
						nextNode.Y = node.Y + 1
					}
				} else if node.Y == nextNode.Y {
					if node.X-nextNode.X > 1 {
						nextNode.X = node.X - 1
					} else if nextNode.X-node.X > 1 {
						nextNode.X = node.X + 1
					}
				} else {
					distance := getDistance(node, nextNode)

					if distance >= 2 {
						if node.X > nextNode.X && node.Y > nextNode.Y {
							if node.X-nextNode.X == 2 && node.Y-nextNode.Y == 2 {
								nextNode.X = node.X - 1
								nextNode.Y = node.Y - 1
							} else if node.X-nextNode.X == 1 {
								nextNode.X = node.X
								nextNode.Y = node.Y - 1
							} else {
								nextNode.Y = node.Y
								nextNode.X = node.X - 1
							}
						} else if node.X > nextNode.X && node.Y < nextNode.Y {
							if node.X-nextNode.X == 2 && nextNode.Y-node.Y == 2 {
								nextNode.X = node.X - 1
								nextNode.Y = node.Y + 1
							} else if node.X-nextNode.X == 1 {
								nextNode.X = node.X
								nextNode.Y = node.Y + 1
							} else {
								nextNode.Y = node.Y
								nextNode.X = node.X - 1
							}
						} else if node.X < nextNode.X && node.Y > nextNode.Y {
							if nextNode.X-node.X == 2 && node.Y-nextNode.Y == 2 {
								nextNode.X = node.X + 1
								nextNode.Y = node.Y - 1
							} else if nextNode.X-node.X == 1 {
								nextNode.X = node.X
								nextNode.Y = node.Y - 1
							} else {
								nextNode.Y = node.Y
								nextNode.X = node.X + 1
							}
						} else if node.X < nextNode.X && node.Y < nextNode.Y {
							if nextNode.X-node.X == 2 && nextNode.Y-node.Y == 2 {
								nextNode.X = node.X + 1
								nextNode.Y = node.Y + 1
							} else if nextNode.X-node.X == 1 {
								nextNode.X = node.X
								nextNode.Y = node.Y + 1
							} else {
								nextNode.Y = node.Y
								nextNode.X = node.X + 1
							}
						}
					}
				}

				if j == 9 {
					visited[nextNode.getKey()] = true
				}
			}

		}
	}

	fmt.Println(len(visited))
}

func getDistance(head, tail *Node) float64 {
	absX := abs(head.X, tail.X)
	absY := abs(head.Y, tail.Y)

	xSquared := absX * absX
	ySquared := absY * absY

	total := xSquared + ySquared

	root := math.Sqrt(float64(total))

	return root
}

func abs(a, b int) int {
	result := a - b

	if result < 0 {
		return result * -1
	}

	return result
}

func (node Node) getKey() string {
	return fmt.Sprintf("%d:%d", node.X, node.Y)
}
