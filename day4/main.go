package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// define time variable
var buildTime string

// define const
const charX = "X"
const charM = "M"
const charA = "A"
const charS = "S"
const xmasLen = 3
const masLen = 1

func masCount_v2(matrix [][]string) int {
	count := 0
	verLen := len(matrix)
	horLen := len(matrix[0])

	for i := 1; i < verLen-masLen; i++ {
		for j := 1; j < horLen-masLen; j++ {

			if matrix[i][j] != charA {
				continue
			}


				leftUp := matrix[i-1][j-1]
				rightDown := matrix[i+1][j+1]
				rightUp := matrix[i-1][j+1]
				leftDown := matrix[i+1][j-1]

				diag1 := (leftUp == charM && rightDown == charS) ||
					(leftUp == charS && rightDown == charM)

				diag2 := (rightUp == charM && leftDown == charS) ||
					(rightUp == charS && leftDown == charM)

				if diag1 && diag2 {
					count++
				}


		}
	}

	return count
}

func xmasCount_v2(matrix [][]string) int {

	count := 0
	verLen := len(matrix)
	horLen := len(matrix[0])

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

	for i := range matrix {
		for j := range matrix[i] {
			for _, distance := range distances {

				x := distance[0]
				y := distance[1]

				endI := i + xmasLen*y
				endJ := j + xmasLen*x

				if endI >= 0 && endI < verLen && endJ >= 0 && endJ < horLen {
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
	if buildTime == "" {
		buildTime = "Build Time not set"
	}
	fmt.Println("\nBuild time:", buildTime)
}
