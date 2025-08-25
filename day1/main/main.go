package main

import (
    "log"
    "bufio"
    "fmt"
    "sort"
    "os"
    "strings"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func absInt(n int) (int) {
    if n < 0 {
        return -n
    } 
    return n
}

func countDist(left []int, right []int) (int) {


    leftCopy := make([]int, len(left))
    copy(leftCopy, left)

    rightCopy := make([]int, len(right))
    copy(rightCopy, right)

    // 8. sorting is not required in sim! or is it?
    // solved: not required! so make it into dist function
    sort.Ints(leftCopy)
    sort.Ints(rightCopy)

    // 7. change name to make it meaningful
    // solved: change count to distance
    // count := 0
    distance := 0

    for i,_ := range leftCopy {

        // 4. Is there any better way to write?
        // vvv old code vvv
        // if right[i] > leftCopy[i] {
        //     distance += right[i] - leftCopy[i]
        // } else if right[i] < leftCopy[i] {
        //     distance += leftCopy[i] - right[i]
        // } else {
        //
        // }

        // solved: subtract and turn the value to absolute
        distance += absInt(leftCopy[i] - rightCopy[i])

    }

    return distance

}

func countSim(left []int, right []int) (int) {

    m := make(map[int]int)

    for _,number := range right {

        // 6. Find a way to simplify this
        // solved: since m[number] will return 0, it's fine to add 1 to it directly
        // _, ok := m[number]
        // if !ok {
        //     m[number] = 0
        // }
        m[number] += 1
    }

    // 7. change name to make it meaningful
    // solved: changed to score
    score := 0

    for _,number := range left {

        // m[number] will be 0 if there're no key
        score += number * m[number]
    }

    return score
}


func main() {

    // 5. Memory will be hard to handle if the file is too big (advanced)
    //    hint: try bufio scanner to execute line by line 
    // data, err := os.ReadFile("./input")

    file, err := os.Open("./input")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    left := []int{}
    right := []int{}

    lineNum := 1

    for scanner.Scan() {
        line := scanner.Text()
        // line = string(line)

        // 3. What if there're no new line character in the file?
        // solved: use string.TrimSuffix to trim \n
        // old code: lines = lines[:len(lines)-1]
        line = strings.TrimSpace(line)


        // 1. Error if space is not definied
        // solved: using strings.Fields to automatically check space or indent
        // old code: splitText := strings.Split(text, "   ")
        splitText := strings.Fields(line)

        // 2. Error is here for a reason, try to handle it
        // solved: use continue to pass the current invalid value
        leftNum, err := strconv.Atoi(splitText[0])
        if err != nil { fmt.Println("Invalid value on line", lineNum, ":", splitText[0]); continue }

        rightNum, err := strconv.Atoi(splitText[1])
        if err != nil { fmt.Println("Invalid value on line", lineNum, ":", splitText[1]); continue }

        left = append(left, leftNum)
        right = append(right, rightNum)
    }


    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


    // 9. left and right is "參考型別"
    // solved: I don't quite get it but I think using a pointer will be a good idea
    sim := countSim(left, right)

    // however, I don't think sorting the mem will be a good idea, or is it?
    dist := countDist(left, right)

    //    fmt.Println(left, right)
    fmt.Println("Answer of question1 is:", dist)
    fmt.Println("Answer of question2 is:", sim)

}
