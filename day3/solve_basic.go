package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var bitArr [12]int
	var ctr int
	for scanner.Scan() {
		ctr += 1
		for i, runee := range scanner.Text() {
			var bit = runee - '0'
			if bit == 1 {
				bitArr[i] += 1
			}
		}
	}
	var result = 0
	var otherResult = 0
	var shifter = 1 << 11

	for _, element := range bitArr {
		if element > (ctr / 2) {
			result += shifter
		} else {
			otherResult += shifter
		}
		shifter = shifter >> 1
	}
	fmt.Printf("%d\n", result*otherResult)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
