package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func countStones(calculatedValues map[[2]int]int, number string, blinksLeft int, indent int) int {
    if blinksLeft == 0 {
        return 0
    }

    n, _ := strconv.Atoi(number)

    key := [2]int{n, blinksLeft}
    val, ok := calculatedValues[key]
    if ok {
        return val
    }

    switch {
    case n == 0:
        calculatedValues[key] = countStones(calculatedValues, "1", blinksLeft - 1, indent + 1)
    case len(number) % 2 == 0:
        tl, _ := strconv.Atoi(number[:len(number)/2])
        l := strconv.Itoa(tl)
        tr, _ := strconv.Atoi(number[len(number)/2:])
        r := strconv.Itoa(tr)
        nl := countStones(calculatedValues, l, blinksLeft - 1, indent + 1)
        nr := countStones(calculatedValues, r, blinksLeft - 1, indent + 1)
        calculatedValues[key] =  1 + nl + nr
    default:
        calculatedValues[key] = countStones(calculatedValues, strconv.Itoa(2024*n), blinksLeft - 1, indent + 1)
    }

    return calculatedValues[key]
}

func main() {
    //data, _ := os.ReadFile("day11_input_small")
    data, _ := os.ReadFile("day11_input")

    inputs := strings.Split(string(data[:len(data) - 1]), " ")
    calcalculatedValues := make(map[[2]int]int, 0)

    total := 0
    for _, v := range(inputs) {
        total += 1 + countStones(calcalculatedValues, v, 75, 0)
    }

    fmt.Println(total)
}
