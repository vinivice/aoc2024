package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Node struct {
    position [2]int
    score int
}

func main() {
    //file, _ := os.Open("day18_input_small")
    //maxDimension := 6
    //nFallenBytes := 12
    file, _ := os.Open("day18_input")
    maxDimension := 70
    nFallenBytes := 1024

    scanner := bufio.NewScanner(file)

    corruptedCoordinates := make(map[[2]int]bool, 0)
    for i := 0; scanner.Scan() && i < nFallenBytes; i++ {
        line := scanner.Text()
        lineParts := strings.Split(line, ",")
        x, _ := strconv.Atoi(lineParts[0])
        y, _ := strconv.Atoi(lineParts[1])

        corruptedCoordinates[[2]int{x, y}] = true
    }

    visitedCoordinates := make(map[[2]int]int, 0)

    checkList := make([]Node, 0)
    checkList = append(checkList, Node{[2]int{0, 0}, 0})

    for len(checkList) > 0 {
        currentNode := checkList[0]
        checkList = checkList[1:]

        if currentNode.position[0] < 0 || currentNode.position[0] > maxDimension || currentNode.position[1] < 0 || currentNode.position[1] > maxDimension {
            continue
        }

        _, ok := visitedCoordinates[currentNode.position]
        if ok {
            continue
        }

        _, ok = corruptedCoordinates[currentNode.position]
        if ok {
            continue
        }

        visitedCoordinates[currentNode.position] = currentNode.score
        if currentNode.position[0] == maxDimension && currentNode.position[1] == maxDimension {
            break
        }

        x := currentNode.position[0]
        y := currentNode.position[1]

        checkList = append(checkList, Node{[2]int{x - 1, y}, currentNode.score + 1})
        checkList = append(checkList, Node{[2]int{x + 1, y}, currentNode.score + 1})
        checkList = append(checkList, Node{[2]int{x, y - 1}, currentNode.score + 1})
        checkList = append(checkList, Node{[2]int{x, y + 1}, currentNode.score + 1})

        slices.SortFunc(checkList, func(a, b Node) int {
            return a.score - b.score
        })
    }

    fmt.Println(visitedCoordinates[[2]int{maxDimension, maxDimension}])
}
