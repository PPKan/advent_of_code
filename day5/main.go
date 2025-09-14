package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

func checkUpdates(updates [][]int, rules map[int][]int) (int, error) {

	result := 0

	for _, update := range updates {

		if len(update)%2 == 0 {
			return 0, fmt.Errorf("rule length should be odd")
		}

		pass := true

		/*
			1. The rest of the numbers should be contained in its previous values' slice
			e.g. [1~n] should be in the slice of map[0]
			2. so the algo should be, iterate through first numbers,
			and check if the rest be in the firsts' slice
		*/
		fmt.Println("\nCheck begins with ===", update, "===")
		for i := range update {
			for j := range update {

				if j <= i {
					continue
				}

				if !slices.Contains(rules[update[i]], update[j]) {
					fmt.Println("The pass did not pass since", update[j],
						"does not inside", update[i], rules[update[i]])

					pass = false
					break
				}
			}

			if !pass {
				break
			}
		}

		if pass {
			fmt.Println("The test has passed.", update[(len(update)-1)/2], "has been added.")
			result += update[(len(update)-1)/2]
		}

	}

	fmt.Println("End of the check")

	return result, nil

}

/*
1. Use scanner
2. Put rules into map
3. iterate through each line, check if value in map
4. if line all right, add middle count into result
*/
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

	result, err := checkUpdates(updates, rules)
	if err != nil {
		log.Fatalf("\nerror checking updates %v", err)
	}
	// fmt.Println(rules)
	// fmt.Println(updates)

	fmt.Println(result)

}
