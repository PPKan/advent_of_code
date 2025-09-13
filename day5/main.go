package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func strSliToIntSli(strSli []string) ([]int, error) {

	intSli := []int
	for _,str := range strSli {
		num, err := strconv.Atoi(str)
		if err != nil {
			return []int, fmt.Errorf("Error converting string to interger: %v", err)
	}
		intSli = append(intSli, )

	}
}

func parseRule(line string) []int {

	isPart1 := true

	if line == "" {
		isPart1 = false
	}

	ruleMap := make(map[int][]int)

	rule := strings.Split(line, "|")

	if isPart1 {
		fmt.Println(ruleMap)
	}

	return []int{1, 2}

}

func addRule(ruleNums []int) {
}

/*
1. Use scanner
2. Put rules into map
3. iterate through each line, check if value in map
4. if line all right, add middle count into result
*/
func main() {

	// file, err := os.Open("./input")
	file, err := os.Open("./mini-input")
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parseRule(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
