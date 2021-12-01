package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strconv"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    last_last_last := 0
	last_last := 0
	last := 0
	increased := 0
	sumOld := 0
	sumNew := 0
    for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		sumOld = last_last_last + last_last + last
		sumNew = last_last + last + i
		if sumNew > sumOld  {
			increased += 1
		}
		last_last_last = last_last
		last_last = last
		last = i
    }
	increased -= 3
	fmt.Printf("%d\n", increased)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}