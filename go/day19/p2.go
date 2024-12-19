package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countMatchingPatters(design string, patterns []string, memory map[string]int) int {
    numberArrangements, ok := memory[design]
    if ok {
        return numberArrangements
    }

    lenDesign := len(design)
    if lenDesign <= 0 {
        return 1
    }

    subtotal := 0
    for _, p := range(patterns) {
        lenPattern := len(p)

        if lenDesign >= lenPattern && design[:lenPattern] == p {
            numberArrangements = countMatchingPatters(design[lenPattern:], patterns, memory)
            memory[design[lenPattern:]] = numberArrangements
            subtotal += numberArrangements
        }
    }

    return subtotal
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
    memory := make(map[string]int)
    for scanner.Scan() {
        input = scanner.Text()
        total += countMatchingPatters(input, patterns, memory)
    }

    fmt.Println(total)
}

