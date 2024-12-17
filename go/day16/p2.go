package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int
const (
    EAST Direction = iota
    SOUTH
    WEST
    NORTH
)

func (d Direction) getVectorComponents() (int, int) {
    switch d {
    case EAST:
        return 1, 0
    case SOUTH:
        return 0, 1
    case WEST:
        return -1, 0
    case NORTH:
        return 0, -1
    }

    return 0, 0
}

type Node struct {
    col int
    row int
    score int
    direction Direction
    bestPathsCoordinates map[[2]int]bool
}

func main() {
    //file, _ := os.Open("day16_input_small1")
    file, _ := os.Open("day16_input")

    scanner := bufio.NewScanner(file)

    var startPosition Node
    raceMap := make([][]rune, 0)
    for row := 0; scanner.Scan(); row++ {
        line := scanner.Text()

        lineComponents := make([]rune, 0)
        for col, c := range(line) {
            lineComponents = append(lineComponents, c)
            if c == 'S' {
                startPosition = Node{col, row, 0, EAST, make(map[[2]int]bool, 0)}
            }
        }
        raceMap = append(raceMap, lineComponents)
    }

    checkList := make(map[[3]int]Node, 0)
    checkList[[3]int{startPosition.col, startPosition.row, int(startPosition.direction)}] = startPosition
    visited := make(map[[3]int]int, 0)
    for len(checkList) > 0 {
        currentNode := Node{-1, -1, 10000000, -1, nil} 
        currentKey := [3]int{-1, -1, -1}
        for k, v := range(checkList) {
            if v.score < currentNode.score {
                currentNode = v
                currentKey = k
            }
        }
        delete(checkList, currentKey)

        score, ok := visited[[3]int{currentNode.col, currentNode.row, int(currentNode.direction)}]

        if ok && score < currentNode.score {
            continue
        }

        if raceMap[currentNode.row][currentNode.col] == 'E' {
            fmt.Println(len(currentNode.bestPathsCoordinates) + 1)
            break
        }

        if raceMap[currentNode.row][currentNode.col] == '#' {
            continue
        }

        currentNode.bestPathsCoordinates[[2]int{currentNode.col, currentNode.row}] = true
        pathD := make(map[[2]int]bool, 0)
        pathR := make(map[[2]int]bool, 0)
        pathL := make(map[[2]int]bool, 0)

        for k, v := range(currentNode.bestPathsCoordinates) {
            pathD[k] = v
            pathR[k] = v
            pathL[k] = v
        }

        dirCol, dirRow := currentNode.direction.getVectorComponents()
        rDirCol, rDirRow := -dirRow, dirCol
        lDirCol, lDirRow := dirRow, -dirCol

        direct := Node{currentNode.col + dirCol, currentNode.row + dirRow, currentNode.score + 1, currentNode.direction, pathD}
        right := Node{currentNode.col + rDirCol, currentNode.row + rDirRow, currentNode.score + 1001, (currentNode.direction + 1) % 4, pathR}
        left := Node{currentNode.col + lDirCol, currentNode.row + lDirRow, currentNode.score + 1001, (currentNode.direction + 3) % 4, pathL}

        keyD := [3]int{direct.col, direct.row, int(direct.direction)}
        keyR := [3]int{right.col, right.row, int(right.direction)}
        keyL := [3]int{left.col, left.row, int(left.direction)}

        node, ok := checkList[keyD]
        if !ok {
            checkList[keyD] = direct
        } else if direct.score == node.score {
            temp := Node{node.col, node.row, node.score, node.direction, make(map[[2]int]bool, 0)}
            for k := range(direct.bestPathsCoordinates) {
                temp.bestPathsCoordinates[[2]int{k[0], k[1]}] = true
            }
            for k := range(node.bestPathsCoordinates) {
                temp.bestPathsCoordinates[[2]int{k[0], k[1]}] = true
            }
            checkList[keyD] = temp
        } else if direct.score < node.score {
            checkList[keyD] = direct
        }

        node, ok = checkList[keyR]
        if !ok {
            checkList[keyR] = right
        } else if right.score == node.score {
            temp := Node{node.col, node.row, node.score, node.direction, make(map[[2]int]bool, 0)}
            for k := range(right.bestPathsCoordinates) {
                temp.bestPathsCoordinates[[2]int{k[0], k[1]}] = true
            }
            for k := range(node.bestPathsCoordinates) {
                temp.bestPathsCoordinates[[2]int{k[0], k[1]}] = true
            }
            checkList[keyR] = temp
        } else if right.score < node.score {
            checkList[keyR] = right
        }

        node, ok = checkList[keyL]
        if !ok {
            checkList[keyL] = left
        } else if left.score == node.score {
            temp := Node{node.col, node.row, node.score, node.direction, make(map[[2]int]bool, 0)}
            for k := range(left.bestPathsCoordinates) {
                temp.bestPathsCoordinates[[2]int{k[0], k[1]}] = true
            }
            for k := range(node.bestPathsCoordinates) {
                temp.bestPathsCoordinates[[2]int{k[0], k[1]}] = true
            }
            checkList[keyL] = temp
        } else if left.score < node.score {
            checkList[keyL] = left
        }

        visited[[3]int{currentNode.col, currentNode.row, int(currentNode.direction)}] = currentNode.score
    }
}
