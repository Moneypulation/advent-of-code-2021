package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var coordsFrom = make([][2]int, 0)
var coordsTo = make([][2]int, 0)
var diagram [1000][1000]int

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isHigher(a, b int) bool {
	return a >= b
}

func getAtLeastTwo() int {
	var sum = 0
	for i, element := range diagram {
		for j, _ := range element {
			if diagram[i][j] >= 2 {
				sum += 1
			}
		}
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			diagram[i][j] = 0
		}
	}

	for scanner.Scan() {
		var currLine = scanner.Text()
		var currFrom = [2]int{0, 0}
		var currTo = [2]int{0, 0}
		fmt.Sscanf(currLine, "%d,%d -> %d,%d", &currFrom[0], &currFrom[1], &currTo[0], &currTo[1])

		if currFrom[0] == currTo[0] {
			var higher = getMax(currFrom[1], currTo[1])
			var lower = getMin(currFrom[1], currTo[1])
			for i := lower; i <= higher; i++ {
				diagram[currFrom[0]][i] += 1
			}
		} else if currFrom[1] == currTo[1] {
			var higher = getMax(currFrom[0], currTo[0])
			var lower = getMin(currFrom[0], currTo[0])
			for i := lower; i <= higher; i++ {
				diagram[i][currFrom[1]] += 1
			}
		} else {
			if isHigher(currFrom[0], currTo[0]) {
				// x--
				if isHigher(currFrom[1], currTo[1]) {
					// y--
					for i, j := currFrom[0], currFrom[1]; i >= currTo[0]; i, j = i-1, j-1 {
						diagram[i][j] += 1
					}
				} else {
					// y++
					for i, j := currFrom[0], currFrom[1]; i >= currTo[0]; i, j = i-1, j+1 {
						diagram[i][j] += 1
					}
				}
			} else {
				// x++
				if isHigher(currFrom[1], currTo[1]) {
					// y --
					for i, j := currFrom[0], currFrom[1]; i <= currTo[0]; i, j = i+1, j-1 {
						diagram[i][j] += 1
					}
				} else {
					// y++
					for i, j := currFrom[0], currFrom[1]; i <= currTo[0]; i, j = i+1, j+1 {
						diagram[i][j] += 1
					}
				}
			}
		}
	}

	fmt.Printf("Result: %d\n", getAtLeastTwo())

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
