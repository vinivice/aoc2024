package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    //f, err := os.Open("day1_input_small")
    f, err := os.Open("day1_input")

    if err != nil {
        panic(1)
    }

    scanner := bufio.NewScanner(f)

    left := make([]int, 0)
    right := make([]int, 0)
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " ")
        l, _ := strconv.Atoi(line[0])
        r, _ := strconv.Atoi(line[3])
        left = append(left, l)
        right = append(right, r)
    }

    frequencyMap := make(map[int]int, 0)

    for _, v := range(right) {
        frequencyMap[v] += 1
    }

    total := 0
    for _, v := range(left) {
        total += v * frequencyMap[v]
    }

    fmt.Println(total)
}
