package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
}

func main() {
    //file, _ := os.Open("day16_input_small2")
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
                startPosition = Node{col, row, 0, EAST}
            }
        }
        raceMap = append(raceMap, lineComponents)
    }

    cmp := func(a, b Node) int {
        return a.score - b.score
    }

    checkList := make([]Node, 0)
    checkList = append(checkList, startPosition)
    for len(checkList) > 0 {
        currentNode := checkList[0]
        checkList = checkList[1:]

        if raceMap[currentNode.row][currentNode.col] == 'E' {
            fmt.Println(currentNode.score)
            break
        }

        if raceMap[currentNode.row][currentNode.col] == '#' {
            continue
        }

        dirCol, dirRow := currentNode.direction.getVectorComponents()
        rDirCol, rDirRow := -dirRow, dirCol
        lDirCol, lDirRow := dirRow, -dirCol

        direct := Node{currentNode.col + dirCol, currentNode.row + dirRow, currentNode.score + 1, currentNode.direction}
        right := Node{currentNode.col + rDirCol, currentNode.row + rDirRow, currentNode.score + 1001, (currentNode.direction + 1) % 4}
        left := Node{currentNode.col + lDirCol, currentNode.row + lDirRow, currentNode.score + 1001, (currentNode.direction + 3) % 4}

        addDirect := true
        addRight := true
        addLeft := true

        for idx, node := range(checkList) {
            if node.col == direct.col && node.row == direct.row && node.direction == direct.direction {
                addDirect = false
                if direct.score < node.score {
                    checkList[idx].score = direct.score
                    checkList[idx].direction = direct.direction
                }
            }
            if node.col == right.col && node.row == right.row && node.direction == right.direction {
                addRight = false
                if right.score < node.score {
                    checkList[idx].score = right.score
                    checkList[idx].direction = right.direction
                }
            }
            if node.col == left.col && node.row == left.row && node.direction == left.direction {
                addLeft = false
                if left.score < node.score {
                    checkList[idx].score = left.score
                    checkList[idx].direction = left.direction
                }
            }
        }
        
        if addDirect {
            checkList = append(checkList, direct)
        }
        if addRight {
            checkList = append(checkList, right)
        }
        if addLeft {
            checkList = append(checkList, left)
        }

        slices.SortFunc(checkList, cmp)
    }
}
