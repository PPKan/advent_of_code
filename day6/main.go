package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

// map[1, 2]{obstacle}
func parseMap(blockMap string) (map[Coordinate]string, error) {

	sliceMap := strings.Fields(blockMap)

	parsedMap := make(map[Coordinate]string)

	for i, verLine := range sliceMap {
		for j, byteNode := range verLine {
			node := string(byteNode)

		
			switch node {
			case ".":
				parsedMap[Coordinate{X: i, Y: j}] = "dot"
			case "#":
				parsedMap[Coordinate{X: i, Y: j}] = "block"
			case "^":
				parsedMap[Coordinate{X: i, Y: j}] = "up"
			case "v":
				parsedMap[Coordinate{X: i, Y: j}] = "down"
			case ">":
				parsedMap[Coordinate{X: i, Y: j}] = "right"
			case "<":
				parsedMap[Coordinate{X: i, Y: j}] = "left"
			default:
				return nil, fmt.Errorf("unrecognized node")
			}

		}
	}
	fmt.Println(parsedMap)
	return parsedMap, nil
}

func findExit() {
}

func main() {
	file, err := os.ReadFile("./mini-input")
	if err != nil {
		log.Fatalf("Error reading file %v", err)
	}

	data := string(file)

	parseMap(data)

}
