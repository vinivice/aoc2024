package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkMatchingPatters(design string, patterns []string) bool {
    lenDesign := len(design)
    if lenDesign <= 0 {
        return true
    }

    for _, p := range(patterns) {
        lenPattern := len(p)

        if lenDesign >= lenPattern && design[:lenPattern] == p {
            if checkMatchingPatters(design[lenPattern:], patterns) {
                return true
            }
        }
    }

    return false
}

func main () {
    //file, _ := os.Open("day19_input_small")
    file, _ := os.Open("day19_input")

    scanner := bufio.NewScanner(file)

    scanner.Scan()

    input := scanner.Text()
    patterns := strings.Split(input, ", ")
    
    scanner.Scan()

    total := 0
    for scanner.Scan() {
        input = scanner.Text()
        if checkMatchingPatters(input, patterns) {
            total++
        }
    }

    fmt.Println(total)
}

