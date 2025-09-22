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

type Wall struct {
	left  int
	right int
	up    int
	down  int
}

// map[1, 2]{obstacle}
func parseMap(blockMap string) (map[Coordinate]string, Wall, Coordinate, error) {

	sliceMap := strings.Fields(blockMap)
	parsedMap := make(map[Coordinate]string)
	wall := Wall{
		left:  0,
		right: len(sliceMap[0]),
		up:    0,
		down:  len(sliceMap),
	}

	var startingCoor Coordinate

	for i, verLine := range sliceMap {
		for j, byteNode := range verLine {

			node := string(byteNode)
			currentCoor := Coordinate{X: i, Y: j}

			switch node {
			case ".":
				parsedMap[currentCoor] = "dot"
			case "#":
				parsedMap[currentCoor] = "block"
			case "^":
				startingCoor = currentCoor
				parsedMap[currentCoor] = "up"
			// case "v":
			// 	parsedMap[currentCoor] = "down"
			// case ">":
			// 	parsedMap[currentCoor] = "right"
			// case "<":
			// 	parsedMap[currentCoor] = "left"
			default:
				return nil, Wall{}, Coordinate{}, fmt.Errorf("unrecognized node")
			}

		}
	}
	// fmt.Println(parsedMap, startingCoor)
	return parsedMap, wall, startingCoor, nil
}

func findExit(dirMap map[Coordinate]string, wall Wall, startingCoor Coordinate) int {

	currentCoor := startingCoor

	direction := Coordinate{X: -1, Y: 0}

	count := 1


	walkedMap := make(map[Coordinate]struct{})
	walkedMap[startingCoor] = struct{}{}

	for {

		nextCoor := Coordinate{currentCoor.X + direction.X, currentCoor.Y + direction.Y}

		if dirMap[nextCoor] == "block" {
			// fmt.Println("meeting a block at next coordinate")
			// fmt.Println("Direction", direction, "Next coordinate", nextCoor)
			// if direction is up, turn right
			if direction.X == -1 && direction.Y == 0 {
				direction = Coordinate{X: 0, Y: 1}
			} else if direction.X == 0 && direction.Y == 1 {
				// if direction is right, turn down
				direction = Coordinate{X: 1, Y: 0}
			} else if direction.X == 1 && direction.Y == 0 {
				// if direction is down, turn left
				direction = Coordinate{X: 0, Y: -1}
			} else if direction.X == 0 && direction.Y == -1 {
				// if direction is left, turn up
				direction = Coordinate{X: -1, Y: 0}
			}
			continue
		}

		// walk
		currentCoor = nextCoor

		_, ok := walkedMap[currentCoor]
		if !ok {
			walkedMap[currentCoor] = struct{}{}
			count += 1
		}

		if currentCoor.Y < wall.left+1 || currentCoor.Y >= wall.right-1 ||
			currentCoor.X < wall.up+1 || currentCoor.X >= wall.down-1 {
			// fmt.Println("Break with", currentCoor, wall)
			break
		}
	}


	return count
}

func main() {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("Error reading file %v", err)
	}

	data := string(file)

	dirMap, wall, startingCoor, err := parseMap(data)
	if err != nil {
		log.Fatalf("Error parsing map data")
	}

	result := findExit(dirMap, wall, startingCoor)

	fmt.Println(result)

}
