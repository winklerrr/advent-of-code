package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// read input from file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	// part 1
	floor := 0
	for _, char := range input {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
	}
	fmt.Println(floor)

	// part 2
	floor = 0
	for i, char := range input {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
		if floor == -1 {
			fmt.Println(i + 1)
			break
		}
	}
}
	