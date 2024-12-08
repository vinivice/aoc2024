package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    //file, _ := os.Open("day8_input_small")
    file, _ := os.Open("day8_input")

    scanner := bufio.NewScanner(file)

    antennasLocations := make(map[rune][][2]int)
    nCols := 0

    row := 0
    for scanner.Scan() {
        line := scanner.Text()
        nCols = len(line)

        for col, v := range(line) {
            if v != '.' {
                list, ok := antennasLocations[v]

                if !ok {
                    list = make([][2]int, 0)
                }

                antennasLocations[v] = append(list, [2]int{col, row})
            }
        }
        row++
    }

    nRows := row

    antinodes := make(map[[2]int]bool, 0)

    for _, v := range(antennasLocations) {
        for i, a := range(v) {
            for _, b := range(v[i + 1:]) {
                if a == b {
                    continue
                }

                delta := [2]int{a[0] - b[0], a[1] - b[1]}
                atn1 := [2]int{a[0] + delta[0], a[1] + delta[1]}
                atn2 := [2]int{b[0] - delta[0], b[1] - delta[1]}

                if atn1[0] >= 0 && atn1[0] < nCols && atn1[1] >= 0 && atn1[1] < nRows {
                    antinodes[atn1] = true
                }

                if atn2[0] >= 0 && atn2[0] < nCols && atn2[1] >= 0 && atn2[1] < nRows {
                    antinodes[atn2] = true
                }
            }
        }
    }

    fmt.Println(len(antinodes))
}

