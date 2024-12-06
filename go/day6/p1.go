package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Vec2 struct {
    x, y int
}

func (v *Vec2) rotateRight90() {
    x := -v.y
    y := v.x

    v.x = x
    v.y = y
}

func (v Vec2) sumCoordinates(v2 Vec2) (int, int) {
    return v.x + v2.x, v.y + v2.y
}


func main() {
    //file, _ := os.Open("day6_input_small")
    file, _ := os.Open("day6_input")
    scanner := bufio.NewScanner(file)

    floorMap := [][]string{}

    for scanner.Scan() {
        line := scanner.Text()
        floorMap = append(floorMap, strings.Split(line, ""))
    }

    dimensions := Vec2{x: len(floorMap[0]), y: len(floorMap)}

    currPosition := Vec2{0, 0}
    for i, v := range(floorMap) {
        for idx, c := range(v) {
            if c == "^" {
                currPosition = Vec2{x: idx, y: i}
                break
            }
        }
    }
    currDir := Vec2{0, -1}

    floorMap[currPosition.y][currPosition.x] = "x"

    for {
        floorMap[currPosition.y][currPosition.x] = "x"
        x, y := currPosition.sumCoordinates(currDir)

        if x < 0 || x >= dimensions.x || y < 0 || y >= dimensions.y {
            break
        }

        if floorMap[y][x] == "#" {
            currDir.rotateRight90()
            continue
        }
        currPosition.x = x
        currPosition.y = y
        
    }

    total := 0
    for _, v := range(floorMap) {
        for _, c := range(v) {
            if c == "x" {
                total += 1
            }
        }
    }

    fmt.Println(total)
}
