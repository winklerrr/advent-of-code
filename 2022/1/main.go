package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// read input from file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// read int per line and add to sum at current slice position, add to new sum after empty line
	scanner := bufio.NewScanner(file)

	var caloriesSum int
	var elfCalories []int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elfCalories = append(elfCalories, caloriesSum)
			caloriesSum = 0
		} else {
			calories, _ := strconv.Atoi(line)
			caloriesSum += calories
		}
	}
	elfCalories = append(elfCalories, caloriesSum)

	// part 1: find the elf with the most calories
	var maxCalories int
	for _, calories := range elfCalories {
		if calories > maxCalories {
			maxCalories = calories
		}
	}
	fmt.Println(maxCalories)

	// part 2: find the top 3 elves with the most calories and add their calories together
	var top3Calories []int
	for i := 0; i < 3; i++ {
		var maxCalories int
		var maxCaloriesIndex int
		for index, calories := range elfCalories {
			if calories > maxCalories {
				maxCalories = calories
				maxCaloriesIndex = index
			}
		}
		top3Calories = append(top3Calories, maxCalories)
		elfCalories[maxCaloriesIndex] = 0
	}
	var top3CaloriesSum int
	for _, calories := range top3Calories {
		top3CaloriesSum += calories
	}
	fmt.Println(top3CaloriesSum)
}


