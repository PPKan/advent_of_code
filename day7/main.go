package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Calibration struct {
	answer  int
	numbers []int
}

func parseLine(line string) (Calibration, error) {

	parsedLine := strings.Split(line, ":")

	var numbers []int
	for number := range strings.FieldsSeq(parsedLine[1]) {

		n, err := strconv.Atoi(number)
		if err != nil {
			return Calibration{0, []int{0, 0}}, fmt.Errorf("error converting string to integer %w", err)
		}
		numbers = append(numbers, n)
	}

	answer, err := strconv.Atoi(parsedLine[0])
	if err != nil {
		return Calibration{0, []int{0, 0}}, fmt.Errorf("error converting string to integer %w", err)
	}

	return Calibration{answer, numbers}, nil
}

func (ca *Calibration) checkResult() int {

	for _,num := range ca.numbers {
		fmt.Println(num)
	}

	return 0
}

func main() {
	file, err := os.Open("mini-input")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("failed to close file %v", cerr)
		}
	}()

	scanner := bufio.NewScanner(file)

	part1result := 0
	for scanner.Scan() {
		line := scanner.Text()
		ca, err := parseLine(line)
		if err != nil {
			log.Fatalf("failed to parse line")
		}

		part1result += ca.checkResult()
	}

	fmt.Println(part1result)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during scanning: %v", err)
	}
}
