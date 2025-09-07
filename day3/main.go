package main

import (
	// "bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

// global regex variables
var mulRegex = regexp.MustCompile(`mul\(([1-9][0-9]{0,2}),([1-9][0-9]{0,2})\)`)
var doRegex = regexp.MustCompile(`do\(\)|don\'t\(\)|mul\(([1-9][0-9]{0,2}),([1-9][0-9]{0,2})\)`)

func stringMul(firstStr string, secStr string) (int, error) {

	firstNum, err := strconv.Atoi(firstStr)
	if err != nil {
		return 0, fmt.Errorf("Error string conversion %v", err)
	}

	secNum, err := strconv.Atoi(secStr)
	if err != nil {
		return 0, fmt.Errorf("Error string conversion %v", err)
	}

	return firstNum * secNum, nil
}

func mulAllMatches(data string) (int, error) {

	result := 0

	allSubMatches := mulRegex.FindAllStringSubmatch(data, -1)

	for _, matches := range allSubMatches {

		firstNum := matches[1]
		secNum := matches[2]

		num, err := stringMul(firstNum, secNum)
		if err != nil {
			return 0, fmt.Errorf("Error string multiplication %v", err)
		}
		result += num
	}

	return result, nil
}

func mulAllDoMatches(data string) (int, error) {

	allSubMatches := doRegex.FindAllStringSubmatch(data, -1)

	const doCommand string = "do()"
	const dontCommand string = "don't()"

	result := 0
	switchEnabled := true

	for _, matches := range allSubMatches {

		command := matches[0]
		firstNum := matches[1]
		secNum := matches[2]

		switch command {
		case doCommand:
			switchEnabled = true
		case dontCommand:
			switchEnabled = false
		// if above two cases matches, default won't be triggered
		default:
			if switchEnabled {
				num, err := stringMul(firstNum, secNum)
				if err != nil {
					return 0, fmt.Errorf("Error string multiplication %v", err)
				}
				result += num
			}

		}
	}

	return result, nil
}

func main() {

	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("An error occured while reading file: %v", err)
	}

	data := string(file)

	partOneResult, err := mulAllMatches(data)
	if err != nil {
		log.Fatalf("An error occoured in Part 1: %v", err)
	}
	fmt.Println(partOneResult) // part 1: 169021493 (correct)

	partTwoResult, err := mulAllDoMatches(data)
	if err != nil {
		log.Fatalf("An error occoured in Part 2: %v", err)
	}
	fmt.Println(partTwoResult) // part 2: 111762583 (correct)

	fmt.Println("\nbuilt time:", time.Now())
}
