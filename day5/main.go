package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 1. for any error checkings, use %w insted of %v, explain in the solution
// solution: %w operates like %v, but I provides alternate functionability like As Is and Unwrap
//
//	which gives some more ways to parse the error
//
// 2. change the naming (hint: use its core func to name)
// solution: change strSli... to parseInts and renewUpdates to sortUpdates
// 5. Idiomatic Go: capacity and memory allocation (space) issue
// solved: use make to specify len and use i to assign var
func parseInts(strSli []string) ([]int, error) {

	intSli := make([]int, len(strSli))
	for i, str := range strSli {
		num, err := strconv.Atoi(str)
		if err != nil {
			return []int{0}, fmt.Errorf("error converting string to interger: %w", err)
		}
		intSli[i] = num
	}

	return intSli, nil
}

func parseRuleParts(data string) (map[int]map[int]struct{}, [][]int, error) {

	sections := strings.Split(data, "\n\n")

	part1 := make(map[int]map[int]struct{})

	for lines := range strings.FieldsSeq(sections[0]) {
		strSep := strings.Split(lines, "|")
		numSep, err := parseInts(strSep)
		if err != nil {
			return nil, nil, fmt.Errorf("\nerror parsing part 1 lines: %w", err)
		}

		first := numSep[0]
		second := numSep[1]

		_, ok := part1[first]
		if ok {
			part1[first][second] = struct{}{}
		} else {
			part1[first] = make(map[int]struct{})
			part1[first][second] = struct{}{}
		}
	}

	var part2 [][]int

	for lines := range strings.FieldsSeq(sections[1]) {
		strSep := strings.Split(lines, ",")
		numSep, err := parseInts(strSep)
		if err != nil {
			return nil, nil, fmt.Errorf("\nerror parsing part 2 lines: %w", err)
		}

		part2 = append(part2, numSep)
	}

	return part1, part2, nil
}

// 3. Efficiency issue (hint: change the map[int][]int to map[int]map[int]bool or map[int]map[int]struct)
// solved: modified data structure using map to do checking
func checkUpdate(update []int, rules map[int]map[int]struct{}) (int, []int, error) {

	if len(update)%2 == 0 {
		return 0, nil, fmt.Errorf("rule length should be odd")
	}

	for i := range update {
		for j := range update {

			if j <= i {
				continue
			}

			_, ok := rules[update[i]][update[j]]
			if !ok {
				// returning malfunctioned update
				return 0, update, nil
			}
		}
	}

	return update[(len(update)-1)/2], nil, nil
}

func HasDuplicates(slice []int) bool {
	seen := make(map[int]bool)
	for _, item := range slice {
		if seen[item] {
			return true // Duplicate found
		}
		seen[item] = true
	}
	return false // No duplicates found
}

// 4. duplication of the code
// solved: modified main to make use of checkUpdate and its return value
func sortUpdates(update []int, rules map[int]map[int]struct{}) (int, error) {

	result := 0

	var seqSlice [][]int

	for i := range update {

		count := 0

		for j := range update {

			if i == j {
				continue
			}

			_, ok := rules[update[j]][update[i]]
			if ok {
				count += 1
			}
		}

		seqSlice = append(seqSlice, []int{update[i], count})
	}

	sort.Slice(seqSlice, func(i, j int) bool {
		return seqSlice[i][1] < seqSlice[j][1]
	})

	var sortedSlice []int
	for _, num := range seqSlice {
		sortedSlice = append(sortedSlice, num[0])
	}

	// slice competency check: count must not duplicate
	isDup := HasDuplicates(sortedSlice)
	if isDup {
		return 0, fmt.Errorf("malfunction sequence detected")
	}

	num, malUpdate, err := checkUpdate(sortedSlice, rules)
	if err != nil {
		return 0, fmt.Errorf("error checking updates %w", err)
	}
	if malUpdate != nil {
		return 0, fmt.Errorf("malfunctioned update found")
	}
	result += num

	return result, nil
}

func main() {

	// file, err := os.Open("./input")
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("\nError reading file %v", err)
	}

	data := string(file)
	rules, updates, err := parseRuleParts(data)
	if err != nil {
		log.Fatalf("\nerror parsing rule parts %v", err)
	}

	result1 := 0
	result2 := 0

	for _, update := range updates {

		num, malUpdate, err := checkUpdate(update, rules)
		if err != nil {
			log.Fatalf("Error checking updates %v", err)
		}
		if num != 0 {
			result1 += num
		}
		if malUpdate != nil {
			malNum, err := sortUpdates(malUpdate, rules)
			if err != nil {
				log.Fatalf("Error sorting updates %v", err)
			}
			result2 += malNum
		}

	}

	fmt.Println("result of part1:", result1)
	fmt.Println("result of part2:", result2)

}
