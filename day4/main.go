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

	distances := [][]int{
		{1, -1},  // leftup
		{-1, -1}, // rightup
		{1, 1},   // leftdown
		{-1, 1},  // rightdown
	}

	for i := range matrixLen {
		for j := range matrixLen {
			if matrix[i][j] == charA {
				continue
			}

			for _, distance := range distances {

				if !(i < matrixLen-masLen && j < matrixLen-masLen && i >= masLen && j >= masLen) {
					continue
				}

				x := distance[0]
				y := distance[1]

				current := matrix[i+y][i+x]
				opposite := matrix[i-y][i-x]

				if (current == charS && opposite == charM) || (current == charM && opposite == charS) {
					count += 1
				}

			}
		}
	}

	return count
}
func masCount(matrix [][]string) int {
	count := 0
	matrixLen := len(matrix)

	for i := 0; i < matrixLen; i++ {
		for j := 0; j < matrixLen; j++ {

			// set wall
			if i < matrixLen-1 && j < matrixLen-1 && i >= 1 && j >= 1 {

				center := matrix[i][j]
				leftUp := matrix[i-1][j-1]
				rightUp := matrix[i-1][j+1]
				leftDown := matrix[i+1][j-1]
				rightDown := matrix[i+1][j+1]

				if center != "A" {
					continue
				}

				if leftUp == "M" && rightDown == "S" {
					if rightUp == "M" && leftDown == "S" {
						count += 1
						continue
					}

					if rightUp == "S" && leftDown == "M" {
						count += 1
						continue
					}
				}

				if leftUp == "S" && rightDown == "M" {
					if rightUp == "M" && leftDown == "S" {
						count += 1
						continue
					}

					if rightUp == "S" && leftDown == "M" {
						count += 1
						continue
					}
				}

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

				if x == 1 && j >= matrixLen-xmasLen {
					continue
				}

				if y == 1 && i >= matrixLen-xmasLen {
					continue
				}

				if x == -1 && j < xmasLen {
					continue
				}

				if y == -1 && i < xmasLen {
					continue
				}

				// fmt.Println(matrixLen, i, j, x, y, distance)

				if matrix[i][j] == charX &&
					matrix[i+1*y][j+1*x] == charM &&
					matrix[i+2*y][j+2*x] == charA &&
					matrix[i+3*y][j+3*x] == charS {
					count += 1
				}

			}

		}
	}
	return count
}

func xmasCount(matrix [][]string) int {
	count := 0
	matrixLen := len(matrix)

	for i := 0; i < matrixLen; i++ {
		for j := 0; j < matrixLen; j++ {

			// left count
			if j < matrixLen-3 {
				if matrix[i][j] == "X" &&
					matrix[i][j+1] == "M" &&
					matrix[i][j+2] == "A" &&
					matrix[i][j+3] == "S" {
					count += 1
				}
			}

			// right count
			if j >= 3 {
				if matrix[i][j] == "X" &&
					matrix[i][j-1] == "M" &&
					matrix[i][j-2] == "A" &&
					matrix[i][j-3] == "S" {
					count += 1
				}
			}

			// down count
			if i < matrixLen-3 {
				if matrix[i][j] == "X" &&
					matrix[i+1][j] == "M" &&
					matrix[i+2][j] == "A" &&
					matrix[i+3][j] == "S" {
					count += 1
				}

			}
			// up count
			if i >= 3 {
				if matrix[i][j] == "X" &&
					matrix[i-1][j] == "M" &&
					matrix[i-2][j] == "A" &&
					matrix[i-3][j] == "S" {
					count += 1
				}
			}

			// left down
			if i < matrixLen-3 && j < matrixLen-3 {
				if matrix[i][j] == "X" &&
					matrix[i+1][j+1] == "M" &&
					matrix[i+2][j+2] == "A" &&
					matrix[i+3][j+3] == "S" {
					count += 1
				}

			}

			// right down
			if i < matrixLen-3 && j >= 3 {
				if matrix[i][j] == "X" &&
					matrix[i+1][j-1] == "M" &&
					matrix[i+2][j-2] == "A" &&
					matrix[i+3][j-3] == "S" {
					count += 1
				}
			}

			// left up
			if i >= 3 && j < matrixLen-3 {
				if matrix[i][j] == "X" &&
					matrix[i-1][j+1] == "M" &&
					matrix[i-2][j+2] == "A" &&
					matrix[i-3][j+3] == "S" {
					count += 1
				}
			}

			// right up
			if i >= 3 && j >= 3 {
				if matrix[i][j] == "X" &&
					matrix[i-1][j-1] == "M" &&
					matrix[i-2][j-2] == "A" &&
					matrix[i-3][j-3] == "S" {
					count += 1
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

	fmt.Println("\nbuilt time:", time.Now())

}
