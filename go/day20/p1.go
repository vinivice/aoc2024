package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Node struct {
    col int
    row int
    length int
    cheatAvailable bool
}

func checkPath(raceMap [][]rune, start, end [2]int, maxPathLenght int) (int, int) {
    checklist := make([]Node, 0)
    checklist = append(checklist, Node{start[0], start[1], 0, false})
    visitedNodes := make(map[[2]int]bool, 0)
    total := 0
    pathLenght := 0
    for len(checklist) > 0 {
        currentNode := checklist[0]
        checklist = checklist[1:]

        if currentNode.length > maxPathLenght {
            break
        }

        if currentNode.col == end[0] && currentNode.row == end[1] {
            total++
            pathLenght = currentNode.length
            continue
        }

        if currentNode.col < 0 || currentNode.col >= len(raceMap[0]) || currentNode.row < 0 || currentNode.row >= len(raceMap) {
            continue
        }


        if raceMap[currentNode.row][currentNode.col] == '#' {
            continue
        }

        _, ok := visitedNodes[[2]int{currentNode.col, currentNode.row}]
        if ok {
            continue
        }


        visitedNodes[[2]int{currentNode.col, currentNode.row}] = true

        checklist = append(checklist, Node{currentNode.col + 1, currentNode.row, currentNode.length + 1, false})
        checklist = append(checklist, Node{currentNode.col - 1, currentNode.row, currentNode.length + 1, false})
        checklist = append(checklist, Node{currentNode.col, currentNode.row + 1, currentNode.length + 1, false})
        checklist = append(checklist, Node{currentNode.col, currentNode.row - 1, currentNode.length + 1, false})

        slices.SortFunc(checklist, func(a, b Node) int {
            return a.length - b.length
        })
    }

    return total, pathLenght
}

func main() {
    //file, _ := os.Open("day20_input_small")
    //const EXPECTEDTIMESAVED = 2
    file, _ := os.Open("day20_input")
    const EXPECTEDTIMESAVED = 100

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

    _, normalPathLenght := checkPath(raceMap, start, end, 10000000000)

    raceMap[start[1]][start[0]] = '.'
    raceMap[end[1]][end[0]] = '.'
    
    total := 0
    for row := 1; row < len(raceMap) - 1; row++ {
        for col := 1; col < len(raceMap[0]) - 1; col++ {
            if raceMap[row][col] != '#' {
                continue
            }

            if (raceMap[row - 1][col] != '.' || raceMap[row + 1][col] != '.') && (raceMap[row][col - 1] != '.' || raceMap[row][col + 1] != '.') {
                continue
            }
            raceMap[row][col] = '.'
            
            subtotal, _ := checkPath(raceMap, start, end, normalPathLenght - EXPECTEDTIMESAVED)
            total += subtotal

            raceMap[row][col] = '#'
        }
    }
    fmt.Println(total)
}
