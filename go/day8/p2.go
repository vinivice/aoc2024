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

                atn := a
                antinodes[atn] = true
                for {
                    atn[0] += delta[0]
                    atn[1] += delta[1]


                    if atn[0] < 0 || atn[0] >= nCols || atn[1] < 0 || atn[1] >= nRows {
                        break
                    }
                    antinodes[atn] = true
                }

                atn = a
                for {
                    atn[0] -= delta[0]
                    atn[1] -= delta[1]


                    if atn[0] < 0 || atn[0] >= nCols || atn[1] < 0 || atn[1] >= nRows {
                        break
                    }
                    antinodes[atn] = true
                }

            }
        }
    }

    fmt.Println(len(antinodes))
}

