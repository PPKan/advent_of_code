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

func NewGridResource(
	parsedGrid map[Coordinate]string,
	wall Wall,
	startingCoor Coordinate,
	isWalked bool,
	walkedGrid map[Coordinate]struct{}) GridResource {

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
	return NewGridResource(parsedMap, wall, startingCoor, false, nil), nil
}

func (gr *GridResource) findExit() (int, bool) {

	// dirMap map[Coordinate]string, wall Wall, startingCoor Coordinate)

	currentCoor := gr.StartingCoor

	direction := Coordinate{X: -1, Y: 0}

	count := 1

	gr.IsWalked = true
	gr.WalkedGrid = make(map[Coordinate]struct{})
	gr.WalkedGrid[gr.StartingCoor] = struct{}{}

	statusMap := make(map[Coordinate][]Coordinate)

	directions := []Coordinate{
		// up
		{X: -1, Y: 0},
		// right
		{X: 0, Y: 1},
		// down
		{X: 1, Y: 0},
		// left
		{X: 0, Y: -1},
	}

	for {
		nextCoor := Coordinate{currentCoor.X + direction.X, currentCoor.Y + direction.Y}

		dirs, ok1 := statusMap[nextCoor]
		if ok1 && slices.Contains(dirs, direction) {
			// fmt.Println(nextCoor, "hit twice. Returning.")
			return -1, true
		} else {
			statusMap[nextCoor] = append(dirs, direction)
		}

		if gr.ParsedGrid[nextCoor] == "block" {

			for i := range directions {
				if direction == directions[i] {
					// fmt.Println("From", direction, "to", directions[(i+1)%4])
					direction = directions[(i+1)%4]
					break
				}
			}

			continue
		}

		// walk
		currentCoor = nextCoor

		_, ok2 := gr.WalkedGrid[currentCoor]
		if !ok2 {
			gr.WalkedGrid[currentCoor] = struct{}{}
			count += 1
		}

		if currentCoor.Y < gr.Wall.left || currentCoor.Y >= gr.Wall.right ||
			currentCoor.X < gr.Wall.up || currentCoor.X >= gr.Wall.down {
			// fmt.Println("Break with", currentCoor, wall)
			break
		}
	}

	return count, false
}

func (gr *GridResource) addObstacle(obsPosition Coordinate) {
	gr.ParsedGrid[obsPosition] = "block"
}

func (gr *GridResource) returnObstacle(obsPosition Coordinate) {
	gr.ParsedGrid[obsPosition] = "dot"
}

func (gr *GridResource) forceLoop() int {
	// func forceLoop(dirMap map[Coordinate]string, wall Wall, startingCoor Coordinate, walkedMap map[Coordinate]struct{}) int {

	// remove starting coordinate from the walked map
	delete(gr.WalkedGrid, gr.StartingCoor)

	count := 0
	for coor := range gr.WalkedGrid {

		// gr.ParsedGrid = addObstacle(gr.ParsedGrid, coor)
		gr.addObstacle(coor)

		simGr := NewGridResource(gr.ParsedGrid, gr.Wall, gr.StartingCoor, false, nil)
		_, isLoop := simGr.findExit()

		if isLoop {
			count += 1
		}

		gr.returnObstacle(coor)
	}

	return count
}

func main() {
	fmt.Println("===Start of the script===")
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

	part1result, _ := parsedGridResource.findExit()
	part2result := parsedGridResource.forceLoop()

	fmt.Println("")
	fmt.Println("Part 1 result:", part1result)
	fmt.Println("Part 2 result:", part2result)

	fmt.Println("\n===End of the script===")
}
