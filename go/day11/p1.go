package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func countStones(number string, blinksLeft int, indent int) int {
    if blinksLeft == 0 {
        return 0
    }

    n, _ := strconv.Atoi(number)

    switch {
    case n == 0:
        return countStones("1", blinksLeft - 1, indent + 1)
    case len(number) % 2 == 0:
        tl, _ := strconv.Atoi(number[:len(number)/2])
        l := strconv.Itoa(tl)
        tr, _ := strconv.Atoi(number[len(number)/2:])
        r := strconv.Itoa(tr)
        nl := countStones(l, blinksLeft - 1, indent + 1)
        nr := countStones(r, blinksLeft - 1, indent + 1)
        return 1 + nl + nr
    default:
        return countStones(strconv.Itoa(2024*n), blinksLeft - 1, indent + 1)
    }
}

func main() {
    //data, _ := os.ReadFile("day11_input_small")
    data, _ := os.ReadFile("day11_input")

    inputs := strings.Split(string(data[:len(data) - 1]), " ")

    total := 0
    for _, v := range(inputs) {
        total += 1 + countStones(v, 25, 0)
    }

    fmt.Println(total)
}
