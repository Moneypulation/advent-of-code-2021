package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var boards = make([][5][5]int, 0)
var boardHits = make([][5][5]bool, 0)

func updateBoards(currDrawing int) {
	for i, element := range boards {
		for j, row := range element {
			for k, column := range row {
				//fmt.Printf("[%d,%d,%d]: %d - %d\n", i, j, k, column, len(boards))
				if column == currDrawing {
					boardHits[i][j][k] = true
				}
			}
		}
	}
}

func getSumOfBoard(boardIndex int) int {
	var sum = 0
	for i, row := range boardHits[boardIndex] {
		for j, column := range row {
			if column == false {
				sum += boards[boardIndex][i][j]
			}
		}
	}
	return sum
}

func checkBoards() int {
	for i, element := range boardHits {
		for j, row := range element {
			var lineRes = element[j][0] && element[j][1] && element[j][2] && element[j][3] && element[j][4]
			if lineRes {
				return getSumOfBoard(i)
			}
			for k, _ := range row {
				var lineRes = element[0][k] && element[1][k] && element[2][k] && element[3][k] && element[4][k]
				if lineRes {
					return getSumOfBoard(i)
				}
			}
		}
	}
	return -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ctr = 0
	var firstRun = true
	var drawings string
	var emptyBoard [5][5]int
	var currentBoard [5][5]int
	var skipNext = false
	var emptyHits = [5][5]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}

	// Parse bingo boards
	for scanner.Scan() {
		var currLine = scanner.Text()
		if skipNext {
			skipNext = false
			continue
		}
		if firstRun {
			drawings = currLine
			firstRun = false
			skipNext = true
			continue
		} else if ctr <= 4 {
			// [row][column]
			// 03039004390
			// fmt.Print(currLine + "\n")
			fmt.Sscanf(currLine, "%d %d %d %d %d", &currentBoard[ctr][0], &currentBoard[ctr][1], &currentBoard[ctr][2], &currentBoard[ctr][3], &currentBoard[ctr][4])
			//fmt.Printf("\nGOT: %d\n", currentBoard[ctr][1])

		} else {
			boards = append(boards, currentBoard)
			boardHits = append(boardHits, emptyHits)
			//fmt.Printf("Added last number: %d\n", currentBoard[4][4])
			currentBoard = emptyBoard
			ctr = 0
			continue
		}
		ctr += 1
	}

	// Draw numbers
	var drawingNumbers = strings.Split(drawings, ",")
	for _, element := range drawingNumbers {
		var elementStr, _ = strconv.Atoi(element)
		updateBoards(elementStr)
		if checkBoards() != -1 {
			fmt.Printf("Found! %d\n", checkBoards()*elementStr)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
