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

func checkLoops(originalMap [][]string, start [2]int, newBlock [2]int, dir Vec2, dimensions Vec2) int {
    floorMap := make([][]string, len(originalMap))
    for i, v := range(originalMap) {
        floorMap[i] = append([]string{}, v...)
    }

    floorMap[newBlock[1]][newBlock[0]] = "#"
    currPosition := Vec2{x: start[0], y: start[1]}
    currDir := dir

    path := [][4]int{}

    for {

        for _, p := range(path) {
            if currPosition.x == p[0] && currPosition.y == p[1] && 
             currDir.x == p[2] && currDir.y == p[3] {
                 return 1
             }
        }
        path = append(path, [4]int{currPosition.x, currPosition.y, currDir.x, currDir.y})

        x, y := currPosition.sumCoordinates(currDir)

        if x < 0 || x >= dimensions.x || y < 0 || y >= dimensions.y {
            return 0
        }

        if floorMap[y][x] == "#" {
            currDir.rotateRight90()
            continue
        }
        currPosition.x = x
        currPosition.y = y
    }
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

    path := make(map[[2]int]bool, 0)
    start := [2]int{currPosition.x, currPosition.y}

    for {
        floorMap[currPosition.y][currPosition.x] = "x"
        path[[2]int{currPosition.x, currPosition.y}] = true

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

    delete(path, start)
    dir := Vec2{0, -1}
    total := 0
    for k := range(path) {
        total += checkLoops(floorMap, start, k, dir, dimensions)
    }

    fmt.Println(total)
}
