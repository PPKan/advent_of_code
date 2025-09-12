package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// define const
const charX = "X"
const charM = "M"
const charA = "A"
const charS = "S"
const xmasLen = 3
const masLen = 1

func masCount_v2(matrix [][]string) int {
	count := 0
	matrixLen := len(matrix)

	for i := range matrixLen {
		for j := range matrixLen {

			if matrix[i][j] != charA {
				continue
			}

			if i < matrixLen-masLen && i >= masLen && j < matrixLen-masLen && j >= 0 {

				diag1 := (matrix[i-1][j-1] == charM && matrix[i+1][j+1] == charS) ||
					(matrix[i-1][j-1] == charM && matrix[i+1][j+1] == charS)

			}

		}
	}

	return count
}

func xmasCount_v2(matrix [][]string) int {

	count := 0
	matrixLen := len(matrix)

	distances := [][]int{
		{1, 0},   // left
		{-1, 0},  // right
		{0, 1},   // down
		{0, -1},  // up
		{1, -1},  // leftup
		{-1, -1}, // rightup
		{1, 1},   // leftdown
		{-1, 1},  // rightdown
	}

	for i := range matrixLen {
		for j := range matrixLen {
			for _, distance := range distances {

				x := distance[0]
				y := distance[1]

				endI := i + xmasLen*y
				endJ := j + xmasLen*x

				if endI >= 0 && endI < matrixLen && endJ >= 0 && endJ < matrixLen {
					if matrix[i][j] == charX &&
						matrix[i+1*y][j+1*x] == charM &&
						matrix[i+2*y][j+2*x] == charA &&
						matrix[i+3*y][j+3*x] == charS {
						count += 1
					}

				}
			}

		}
	}
	return count
}

func main() {

	// file, err := os.ReadFile("./mini-input")
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("An error occured while reading file: %v", err)
	}

	data := string(file)

	sepData := strings.Fields(data)

	verLen := len(sepData)
	hexLen := len(strings.Split(sepData[0], ""))

	matrix := make([][]string, verLen, hexLen)

	for i, line := range sepData {
		splitLine := strings.Split(line, "")
		matrix[i] = splitLine
	}

	resultPartOne := xmasCount_v2(matrix)
	resultPartTwo := masCount_v2(matrix)

	fmt.Println(resultPartOne)
	fmt.Println(resultPartTwo)

	// any hint to hardcode build time?
	fmt.Println("\nbuilt time:", time.Now())
}
