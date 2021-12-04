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
var boardsWon = make([]bool, 0)

func updateBoards(currDrawing int) {
	for i, element := range boards {
		for j, row := range element {
			for k, column := range row {
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

func checkBoards() []int {
	var retArr = make([]int, 0)
	for i, element := range boardHits {
		for j, row := range element {
			var lineRes = element[j][0] && element[j][1] && element[j][2] && element[j][3] && element[j][4]
			if lineRes {
				retArr = append(retArr, i)
			}
			for k, _ := range row {
				var lineRes = element[0][k] && element[1][k] && element[2][k] && element[3][k] && element[4][k]
				if lineRes {
					retArr = append(retArr, i)
				}
			}
		}
	}
	return retArr
}

func checkBoardsWin() [2]int {
	var ctrLost = 0
	var indexLost = 0
	for index, element := range boardsWon {
		if element == false {
			indexLost = index
			ctrLost += 1
		}
	}
	if ctrLost == 1 {
		return [2]int{1, indexLost}
	} else {
		return [2]int{0, 0}
	}
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
	var nextWins = false
	var lastWinningIndex = 0

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
			fmt.Sscanf(currLine, "%d %d %d %d %d", &currentBoard[ctr][0], &currentBoard[ctr][1], &currentBoard[ctr][2], &currentBoard[ctr][3], &currentBoard[ctr][4])

		} else {
			boards = append(boards, currentBoard)
			boardHits = append(boardHits, emptyHits)
			boardsWon = append(boardsWon, false)
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

		for _, element := range checkBoards() {
			boardsWon[element] = true
		}

		var ret = checkBoardsWin()

		if nextWins && ret[0] == 0 {
			var sum = getSumOfBoard(lastWinningIndex)
			fmt.Printf("RESULT: %d*%d=%d\n", sum, elementStr, sum*elementStr)
			break
		}

		if ret[0] == 1 {
			nextWins = true
			lastWinningIndex = ret[1]
			fmt.Printf("LAST WINNING: %d\n", lastWinningIndex)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
