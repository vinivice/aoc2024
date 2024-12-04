package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func basicCheck(entries []int) bool {
    previousDiff := entries[1] - entries[0]
    if previousDiff < -3 || (previousDiff > -1 && previousDiff < 1) || previousDiff > 3 {
        return false
    }

    diff := 0
    for i := 2; i < len(entries); i++ {
        diff = entries[i] - entries[i - 1]
        if diff < -3 || (diff > -1 && diff < 1) || diff > 3 || ((previousDiff < 0) != (diff < 0)) {
            return false
        }
        previousDiff = diff

    }

    return true
}

func check(entries []int) bool {
    if basicCheck(entries) {
        return true
    }

    e := entries[1:]
    if basicCheck(e) {
        return true
    }

    previousDiff := entries[1] - entries[0]
    if previousDiff < -3 || (previousDiff > -1 && previousDiff < 1) || previousDiff > 3 {
        e = slices.Concat(entries[0:1], entries[2:])
        return basicCheck(e)
    }
    
    diff := 0
    for i := 2; i < len(entries); i++ {
        diff = entries[i] - entries[i - 1]
        if diff < -3 || (diff > -1 && diff < 1) || diff > 3 || ((previousDiff < 0) != (diff < 0)) {
            e := slices.Concat(entries[0:i - 1], entries[i:])
            res := basicCheck(e) 
            if res {
                return true
            }

            e = slices.Concat(entries[0:i], entries[i + 1:])
            res = basicCheck(e) 
            return res
        }
        previousDiff = diff
    }

    return true
}

func main() {
    //file, err := os.Open("day2_input_small")
    file, err := os.Open("day2_input")
    if err != nil {
        panic(1)
    }

    scanner := bufio.NewScanner(file)

    total := 0
    for scanner.Scan() {
        line := scanner.Text()
        entries := strings.Split(line, " ")
        entriesInt := make([]int, len(entries))
        for i := 0; i < len(entries); i++ {
            entriesInt[i], _ = strconv.Atoi(entries[i])
        }

        if check(entriesInt) {
            total += 1
        }
    }

    fmt.Println(total)
}
