package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
    col int
    row int
    length int
    cheatTimeRemaining int
}


func main() {
    //file, _ := os.Open("day20_input_small")
    //const expectedTimeSaved = 2
    file, _ := os.Open("day20_input")
    const EXPECTEDTIMESAVED = 100

    const CHEATSAVAILABLE = 20

    scanner := bufio.NewScanner(file)

    var start, end [2]int
    raceMap := make([][]rune, 0)
    for row := 0; scanner.Scan(); row++ {
        line := scanner.Text()

        lineParts := make([]rune, 0)
        for col, c := range(line) {
            lineParts = append(lineParts, c)

            if c == 'S' {
                start = [2]int{col, row}
            }

            if c == 'E' {
                end = [2]int{col, row}
            }
        }
        raceMap = append(raceMap, lineParts)
    }

    normalPathLenght := 0

    checklist := make([]Node, 0)
    checklist = append(checklist, Node{start[0], start[1], 0, 0})
    distancesNormalPath := make(map[[2]int]int, 0)
    for len(checklist) > 0 {
        currentNode := checklist[0]
        checklist = checklist[1:]

        if currentNode.col == end[0] && currentNode.row == end[1] {
            normalPathLenght = currentNode.length
            break
        }

        if currentNode.col < 0 || currentNode.col >= len(raceMap[0]) || currentNode.row < 0 || currentNode.row >= len(raceMap) {
            continue
        }


        if raceMap[currentNode.row][currentNode.col] == '#' {
            continue
        }

        _, ok := distancesNormalPath[[2]int{currentNode.col, currentNode.row}]
        if ok {
            continue
        }

        distancesNormalPath[[2]int{currentNode.col, currentNode.row}] = currentNode.length

        checklist = append(checklist, Node{currentNode.col + 1, currentNode.row, currentNode.length + 1, 0})
        checklist = append(checklist, Node{currentNode.col - 1, currentNode.row, currentNode.length + 1, 0})
        checklist = append(checklist, Node{currentNode.col, currentNode.row + 1, currentNode.length + 1, 0})
        checklist = append(checklist, Node{currentNode.col, currentNode.row - 1, currentNode.length + 1, 0})
    }

    raceMap[start[1]][start[0]] = '.'
    raceMap[end[1]][end[0]] = '.'

    total := 0
    for coordinate, lengthFromStart := range(distancesNormalPath) {
        maxPathLength := normalPathLenght - EXPECTEDTIMESAVED
        if lengthFromStart > maxPathLength {
            continue
        }

        checklist := make([]Node, 0)
        checklist = append(checklist, Node{coordinate[0] + 1, coordinate[1], 1, CHEATSAVAILABLE - 1})
        checklist = append(checklist, Node{coordinate[0] - 1, coordinate[1], 1, CHEATSAVAILABLE - 1})
        checklist = append(checklist, Node{coordinate[0], coordinate[1] + 1, 1, CHEATSAVAILABLE - 1})
        checklist = append(checklist, Node{coordinate[0], coordinate[1] - 1, 1, CHEATSAVAILABLE - 1})
        visitedNodes := make(map[[2]int]bool, 0)
        visitedNodes[coordinate] = true
        for len(checklist) > 0 {
            currentNode := checklist[0]
            checklist = checklist[1:]

            if currentNode.col < 0 || currentNode.col >= len(raceMap[0]) || currentNode.row < 0 || currentNode.row >= len(raceMap) {
                continue
            }

            if currentNode.cheatTimeRemaining < 0 {
                continue
            }

            currentNode.cheatTimeRemaining--

            _, ok := visitedNodes[[2]int{currentNode.col, currentNode.row}]
            if ok {
                continue
            }

            if raceMap[currentNode.row][currentNode.col] == '.' {
                distanceStartCheat := lengthFromStart
                distanceEndCheat := distancesNormalPath[[2]int{currentNode.col, currentNode.row}]
                if currentNode.col == end[0] && currentNode.row == end[1] {
                    distanceEndCheat = normalPathLenght
                }
                totalDistance := distanceStartCheat + currentNode.length + (normalPathLenght - distanceEndCheat)
                if totalDistance <= maxPathLength {
                    total++
                }
            }

            visitedNodes[[2]int{currentNode.col, currentNode.row}] = true

            checklist = append(checklist, Node{currentNode.col + 1, currentNode.row, currentNode.length + 1, currentNode.cheatTimeRemaining})
            checklist = append(checklist, Node{currentNode.col - 1, currentNode.row, currentNode.length + 1, currentNode.cheatTimeRemaining})
            checklist = append(checklist, Node{currentNode.col, currentNode.row + 1, currentNode.length + 1, currentNode.cheatTimeRemaining})
            checklist = append(checklist, Node{currentNode.col, currentNode.row - 1, currentNode.length + 1, currentNode.cheatTimeRemaining})
        }
    }

    fmt.Println(total)
}
