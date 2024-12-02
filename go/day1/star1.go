package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

    slices.Sort(left)
    slices.Sort(right)

    total := 0
    for i := range(left) {
        distance := left[i] - right[i]
        if distance < 0 {
            distance = -distance
        }

        total += distance
    }

    fmt.Println(total)

}
