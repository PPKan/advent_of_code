package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func removeElementByIndex(oriSlice []int, index int) []int {
	slice := make([]int, len(oriSlice))
	copy(slice, oriSlice)

	if index < 0 || index >= len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func isSafe(numbers []int) bool {

	// edge case
	if len(numbers) < 2 {
		return true
	}

	var trend, distance int
	const maxLevelDiff = 3

	// use the trend to validate the rest
	for i := 0; i < len(numbers)-1; i++ {

		currentNum := numbers[i]
		nextNum := numbers[i+1]

		distance = nextNum - currentNum

		// first case: distance must not be same and make sure trend must not be zero
		if distance == 0 {
			return false
		}

		if i == 0 {
			trend = distance
		}

		// second case: distance must not larger than 3
		if absInt(distance) > maxLevelDiff {
			return false
		}

		// third case: status of the trend must not change
		if trend*distance < 0 {
			return false
		}

	}

	return true
}

func isSafeWithoutOne(oriNumbers []int) bool {

	numbers := make([]int, len(oriNumbers))
	copy(numbers, oriNumbers)

	// fmt.Println("\nThe original array is:", numbers)

	if isSafe(numbers) {
		// fmt.Println("This array is safe")
		return true
	} else {

		for i := range numbers {

			// fmt.Println("The array is unsafe, removing index:", i, "which is", numbers[i])
			result := isSafe(removeElementByIndex(numbers, i))

			if result {
				// fmt.Println("This array is safe by removing index", i, "which is", numbers[i])
				return true
			}

		}

		// fmt.Println("After all, the array is still unsafe.")
		return false
	}
}

func strSliToIntSli(strSli []string) ([]int, error) {

	intSli := make([]int, len(strSli))

	for i, str := range strSli {
		int, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intSli[i] = int
	}

	return intSli, nil
}

func main() {

	file, e := os.Open("./input")
	if e != nil {
		fmt.Println("File not valid")
		panic(e)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeCount := 0
	safeCountWo := 0

	for scanner.Scan() {

		line := scanner.Text()

		numbersString := strings.Fields(line)

		numbers, err := strSliToIntSli(numbersString)
		if err != nil {
			fmt.Println("Error found in text:", numbersString, "Error message:", err)
			continue // skip the scan if error
		}

		if isSafe(numbers) {
			safeCount += 1
		}

		if isSafeWithoutOne(numbers) {
			safeCountWo += 1
		}

	}

	fmt.Println(safeCount)
	fmt.Println(safeCountWo)
}
