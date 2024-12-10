package main

import (
	"bufio"
	"fmt"
	"os"
)

func trailWalk(topologicalMap [][]byte, start [2]int) map[[2]int]bool {
    ends := make(map[[2]int]bool, 0)

    height := len(topologicalMap)
    width := len(topologicalMap[0])
    
    currRow := start[1]
    currCol := start[0]
    currId := topologicalMap[currRow][currCol]

    if currId == 9 {
        ends[start] = true
        return ends
    }


    if currRow - 1 >= 0 && topologicalMap[currRow - 1][currCol] == currId + 1 {
        tempEnds := trailWalk(topologicalMap, [2]int{currCol, currRow - 1})

        for k := range(tempEnds) {
            ends[k] = true
        }
    }

    if currRow + 1 < height && topologicalMap[currRow + 1][currCol] == currId + 1 {
        tempEnds := trailWalk(topologicalMap, [2]int{currCol, currRow + 1})

        for k := range(tempEnds) {
            ends[k] = true
        }
    }
    
    if currCol - 1 >= 0 && topologicalMap[currRow][currCol - 1] == currId + 1 {
        tempEnds := trailWalk(topologicalMap, [2]int{currCol - 1, currRow})

        for k := range(tempEnds) {
            ends[k] = true
        }
    }

    if currCol + 1 < width && topologicalMap[currRow][currCol + 1] == currId + 1 {
        tempEnds := trailWalk(topologicalMap, [2]int{currCol + 1, currRow})

        for k := range(tempEnds) {
            ends[k] = true
        }
    }

    return ends
}

func getScores(topologicalMap [][]byte, trailheads [][2]int) int {
    total := 0
    for _, t := range(trailheads) {
        total += len(trailWalk(topologicalMap, t))
    }

    return total

}

func main() {
    //file, _ := os.Open("day10_input_small")
    file, _ := os.Open("day10_input")

    scanner := bufio.NewScanner(file)

    trailheads := make([][2]int, 0)
    topologicalMap := make([][]byte, 0)
    currRow := 0
    for scanner.Scan() {
        line := scanner.Text()
        lastIndex := len(topologicalMap)
        topologicalMap = append(topologicalMap, make([]byte, 0))

        for currCol, c := range(line) {
            topologicalMap[lastIndex] = append(topologicalMap[lastIndex], byte(c - 48))

            if c - 48 == 0 {
                trailheads = append(trailheads, [2]int{currCol, currRow} )
            }
        }

        currRow++
    }

    fmt.Println(getScores(topologicalMap, trailheads))
}
