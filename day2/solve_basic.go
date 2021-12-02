package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	depth := 0
	horizontal := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		number := strings.Split(text, " ")[1]
		i, err := strconv.Atoi(number)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		if strings.HasPrefix(text, "up") {
			depth -= i
		} else if strings.HasPrefix(text, "forward") {
			horizontal += i
		} else if strings.HasPrefix(text, "down") {
			depth += i
		} else {
			fmt.Println("err")
			os.Exit(2)
		}
	}

	fmt.Printf("Result: %d\n", depth*horizontal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
