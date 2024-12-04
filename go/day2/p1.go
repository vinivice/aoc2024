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

        diffLen := len(entries) - 1
        diff := make([]int, diffLen)

        for i := 1; i < len(entries); i++ {
            first, _ := strconv.Atoi(entries[i])
            second, _ := strconv.Atoi(entries[i - 1])
            diff[i - 1] = first - second
        }
        slices.Sort(diff)

        if (diff[0] >= 1 && diff[diffLen - 1] <= 3) || (diff[0] >= -3 && diff[diffLen - 1] <= -1) {
            total += 1
        }
    }

    fmt.Println(total)


}
