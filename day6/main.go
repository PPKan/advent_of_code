package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

// make values into struct (APIize the code)
// Arguments that will not be passing around are:
// 1. Original data
//
// Arguments that will be passing around are:
// 1. Parsed Map
// 2. Wall Map
// 3. Walked Map
type GridResource struct {
	ParsedGrid   map[Coordinate]string
	Wall         Wall
	StartingCoor Coordinate
	IsWalked     bool
	WalkedGrid   map[Coordinate]struct{}
}

func NewGridResource(parsedGrid map[Coordinate]string, wall Wall,
	startingCoor Coordinate, isWalked bool, walkedGrid map[Coordinate]struct{}) GridResource {

	var gridResource GridResource

	gridResource.ParsedGrid = parsedGrid
	gridResource.Wall = wall
	gridResource.StartingCoor = startingCoor
	gridResource.IsWalked = isWalked

	if !isWalked {
		gridResource.WalkedGrid = nil
		return gridResource
	}

	gridResource.WalkedGrid = walkedGrid
	return gridResource
}

func parseMap(blockMap string) (GridResource, error) {
	// (map[Coordinate]string, Wall, Coordinate, error)

	sliceMap := strings.Fields(blockMap)
	parsedMap := make(map[Coordinate]string)
	wall := Wall{
		left:  1,
		right: len(sliceMap[0]) - 1,
		up:    1,
		down:  len(sliceMap) - 1,
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
			default:
				return NewGridResource(nil, Wall{}, Coordinate{}, false, nil), fmt.Errorf("unrecognized node")
				// return nil, Wall{}, Coordinate{}, fmt.Errorf("unrecognized node")
			}

		}
	}
	// fmt.Println(parsedMap, startingCoor)
	// return parsedMap, wall, startingCoor, nil
	return NewGridResource(parsedMap, wall, startingCoor, false, nil), nil
}

func findExit(gridResource GridResource) (int, map[Coordinate]struct{}, bool) {

	// dirMap map[Coordinate]string, wall Wall, startingCoor Coordinate)

	dirMap := gridResource.ParsedGrid
	wall := gridResource.Wall
	startingCoor := gridResource.StartingCoor

	currentCoor := startingCoor

	direction := Coordinate{X: -1, Y: 0}

	count := 1

	walkedMap := make(map[Coordinate]struct{})
	walkedMap[startingCoor] = struct{}{}

	blockMap := make(map[Coordinate][]Coordinate)

	for {

		nextCoor := Coordinate{currentCoor.X + direction.X, currentCoor.Y + direction.Y}

		if dirMap[nextCoor] == "block" {
			// fmt.Println(nextCoor, "is block")
			dirs, ok := blockMap[nextCoor]
			if ok && slices.Contains(dirs, direction) {
				// fmt.Println(nextCoor, "hit twice. Returning.")
				return -1, nil, true
			} else {
				blockMap[nextCoor] = append(dirs, direction)
			}
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

		if currentCoor.Y < wall.left || currentCoor.Y >= wall.right ||
			currentCoor.X < wall.up || currentCoor.X >= wall.down {
			// fmt.Println("Break with", currentCoor, wall)
			break
		}
	}

	return count, walkedMap, false
}

func addObstacle(dirMap map[Coordinate]string, obsPosition Coordinate) map[Coordinate]string {
	dirMap[obsPosition] = "block"
	return dirMap
}

func returnObstacle(dirMap map[Coordinate]string, obsPosition Coordinate) map[Coordinate]string {
	dirMap[obsPosition] = "dot"
	return dirMap
}

func forceLoop(dirMap map[Coordinate]string, wall Wall, startingCoor Coordinate, walkedMap map[Coordinate]struct{}) int {

	// remove starting coordinate from the walked map
	delete(walkedMap, startingCoor)

	count := 0
	for coor := range walkedMap {

		dirMap = addObstacle(dirMap, coor)

		_, _, isLoop := findExit(NewGridResource(dirMap, wall, startingCoor, false, nil))
		if isLoop {
			count += 1
		}

		dirMap = returnObstacle(dirMap, coor)

	}

	return count
}

func main() {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("Error reading file %v", err)
	}

	data := string(file)

	parsedGridResource, err := parseMap(data)
	// dirMap, wall, startingCoor, err := parseMap(data)
	if err != nil {
		log.Fatalf("Error parsing map data")
	}

	part1result, walkedMap, _ := findExit(parsedGridResource)

	dirMap := parsedGridResource.ParsedGrid
	wall := parsedGridResource.Wall
	startingCoor := parsedGridResource.StartingCoor

	part2result := forceLoop(dirMap, wall, startingCoor, walkedMap)

	fmt.Println("")
	fmt.Println("Part 1 result:", part1result)
	fmt.Println("Part 2 result:", part2result)

}
