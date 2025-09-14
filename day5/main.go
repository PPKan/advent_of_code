package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func strSliToIntSli(strSli []string) ([]int, error) {

	var intSli []int
	for _, str := range strSli {
		num, err := strconv.Atoi(str)
		if err != nil {
			return []int{0}, fmt.Errorf("error converting string to interger: %v", err)
		}
		intSli = append(intSli, num)
	}

	return intSli, nil
}

func parseRuleParts(data string) (map[int][]int, [][]int, error) {

	sections := strings.Split(data, "\n\n")

	part1 := make(map[int][]int)

	for lines := range strings.FieldsSeq(sections[0]) {
		strSep := strings.Split(lines, "|")
		numSep, err := strSliToIntSli(strSep)
		if err != nil {
			return nil, nil, fmt.Errorf("\nerror parsing part 1 lines: %v", err)
		}

		first := numSep[0]
		second := numSep[1]

		val, ok := part1[first]
		if ok && len(val) != 0 {
			part1[first] = append(val, second)
		} else {
			part1[first] = []int{second}
		}
	}

	var part2 [][]int

	for lines := range strings.FieldsSeq(sections[1]) {
		strSep := strings.Split(lines, ",")
		numSep, err := strSliToIntSli(strSep)
		if err != nil {
			return nil, nil, fmt.Errorf("\nerror parsing part 2 lines: %v", err)
		}

		part2 = append(part2, numSep)
	}

	return part1, part2, nil
}

func checkUpdate(update []int, rules map[int][]int) (int, error) {

	if len(update)%2 == 0 {
		return 0, fmt.Errorf("rule length should be odd")
	}

	pass := true

	for i := range update {
		for j := range update {

			if j <= i {
				continue
			}

			if !slices.Contains(rules[update[i]], update[j]) {
				// fmt.Println("The pass did not pass since", update[j],
				// 	"does not inside", update[i], rules[update[i]])

				return 0, nil
			}
		}

		if !pass {
			break
		}
	}

	if pass {
		// fmt.Println("The test has passed.", update[(len(update)-1)/2], "has been added.")
		return update[(len(update)-1)/2], nil
	}

	return 0, nil

}

func checkUpdates(updates [][]int, rules map[int][]int) (int, error) {

	result := 0

	for _, update := range updates {
		num, err := checkUpdate(update, rules)
		if err != nil {
			return 0, fmt.Errorf("error checking updates: %v", err)
		}
		result += num
	}

	return result, nil

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

func renewUpdates(updates [][]int, rules map[int][]int) (int, error) {

	result := 0

	for _, update := range updates {

		var seqSlice [][]int

		testNum, err := checkUpdate(update, rules)
		if err != nil {
			return 0, fmt.Errorf("error checking renew updates: %v", err)
		}
		if testNum != 0 {
			// result += testNum
			continue
		}

		for i := range update {

			count := 0

			for j := range update {

				if i == j {
					continue
				}

				if slices.Contains(rules[update[j]], update[i]) {
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
			return 0, fmt.Errorf("malfunction sequence detected %v", err)
		}

		num, err := checkUpdate(sortedSlice, rules)
		if err != nil {
			return 0, fmt.Errorf("error checking updates: %v", err)
		}

		// fmt.Println(sortedSlice, result)

		result += num
	}

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

	result1, err := checkUpdates(updates, rules)
	if err != nil {
		log.Fatalf("\nerror checking updates %v", err)
	}

	result2, err := renewUpdates(updates, rules)
	if err != nil {
		log.Fatalf("\nerror checking updates %v", err)
	}
	// fmt.Println(rules)
	// fmt.Println(updates)

	fmt.Println("result of part1:", result1)
	fmt.Println("result of part2:", result2)

}
