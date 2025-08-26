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


func isSafe(numbers []int) bool {

	isSafe := true
	numLen := len(numbers)

	for i, _ := range numbers {

		distance := numbers[i] - numbers[i+1]

		fmt.Print(distance, " ")

		// the original value doesn't matter
		var isClimb bool

		if distance > 0 {
			isClimb = true
		} else if distance < 0 {
			isClimb = false
		} else {
			isSafe = false
			fmt.Println("Break for neither an inc or dec. dist:", distance, numbers)
			return isSafe
		}

		if i != 0 {
			prevDistance := numbers[i-1] - numbers[i]

			if prevDistance > 0 && isClimb == false {
				isSafe = false
				fmt.Println("Break for changing the climb status dist:", distance, numbers)
				return isSafe
			} else if prevDistance < 0 && isClimb == true {
				isSafe = false
				fmt.Println("Break for changing the climb status dist:", distance, numbers)
				return isSafe
			}

		}

		// unsafe because levers differ by at least 1 and most 3
		if absInt(distance) > 3 {
			isSafe = false
			fmt.Println("Break for dist > 3. Dist:", distance, numbers)
			break
		}

		// end of the loop
		if i == numLen-2 {
			fmt.Println(numbers, isSafe)
			break
		}
	}

	return isSafe
}

func strSliToIntSli(strSli []string) []int {

	intSli := make([]int, len(strSli))

	for i, str := range strSli {
		int, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Value type incorrect:", str)
			fmt.Println(err)
		}

		intSli[i] = int
	}

	return intSli
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
	unsafeCount := 0

	for scanner.Scan() {

		line := scanner.Text()

		numbersString := strings.Fields(line)

		numbers := strSliToIntSli(numbersString)

		if isSafe(numbers) {
			safeCount += 1
		} else {
			unsafeCount += 1
		}

	}

	fmt.Println(safeCount)

}
