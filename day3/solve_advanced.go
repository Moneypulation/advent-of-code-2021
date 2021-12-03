package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func oneRound(inputArr []string, idx int, roundType int) []string {
	var ctr0 = 0
	var ctr1 = 0
	zeroArr := make([]string, 0)
	oneArr := make([]string, 0)
	for _, bitString := range inputArr {
		if bitString[idx] == []byte("0")[0] {
			ctr0 += 1
			zeroArr = append(zeroArr, bitString)
		} else {
			ctr1 += 1
			oneArr = append(oneArr, bitString)
		}
	}
	if ctr0 > ctr1 {
		if roundType == 1 {
			return zeroArr
		} else {
			return oneArr
		}

	} else if ctr1 > ctr0 {
		if roundType == 1 {
			return oneArr
		} else {
			return zeroArr
		}
	} else {
		if roundType == 0 {
			return zeroArr
		} else {
			return oneArr
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var strArr = make([]string, 0)
	var ctr = 0
	var oxy = 0
	var co2 = 0
	for scanner.Scan() {
		//strArr[ctr] = scanner.Text()
		strArr = append(strArr, scanner.Text())
		ctr += 1
	}

	var remaining = strArr
	// oxygen
	for i := 0; i < 12; i++ {
		remaining = oneRound(remaining, i, 1)
		fmt.Printf("len: %d\n", len(remaining))
		if len(remaining) == 1 {
			fmt.Print("yep")
			break
		}
	}

	// co2
	var remaining2 = strArr
	for i := 0; i < 12; i++ {
		remaining2 = oneRound(remaining2, i, 0)
		fmt.Printf("len2: %d\n", len(remaining2))
		if len(remaining2) == 1 {
			fmt.Print("yep yep")
			break
		}
	}

	var shifter = 1 << 11
	for _, element := range remaining[0] {
		if (element - '0') == 1 {
			oxy += shifter
		}
		shifter = shifter >> 1
	}

	shifter = 1 << 11
	for _, element := range remaining2[0] {
		if (element - '0') == 1 {
			co2 += shifter
		}
		shifter = shifter >> 1
	}

	fmt.Printf("Oxy: %d\nCO2: %d\nResult: %d\n", oxy, co2, oxy*co2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
