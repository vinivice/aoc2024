package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    //file, _ := os.Open("day15_input_small2")
    file, _ := os.Open("day15_input")

    scanner := bufio.NewScanner(file)

    floorMap := make([][]rune, 0)
    robotPosition := [2]int{0, 0}
    for row := 0; scanner.Scan(); row++ {
        line := scanner.Text()

        if line == "" {
            break
        }

        currLine := make([]rune, 0)
        for col, c := range(line) {
            currLine = append(currLine, c)
            if c == '@' {
                robotPosition = [2]int{col, row}
            }
        }
        floorMap = append(floorMap, currLine)
    }

    for scanner.Scan() {
        line := scanner.Text()
        for _, moveDir := range(line) {
            move := [2]int{0, 0}

            switch moveDir {
            case '^':
                move = [2]int{0, -1}
            case 'v':
                move = [2]int{0, 1}
            case '<':
                move = [2]int{-1, 0}
            case '>':
                move = [2]int{1, 0}
            }

            currentPosition := robotPosition
            OuterLoop:
            for {
                nextPosition := [2]int{currentPosition[0] + move[0], currentPosition[1] + move[1]}

                switch floorMap[nextPosition[1]][nextPosition[0]] {
                case '#':
                    break OuterLoop
                case 'O':
                    currentPosition = nextPosition
                    continue OuterLoop
                case '.':
                    floorMap[nextPosition[1]][nextPosition[0]] = floorMap[currentPosition[1]][currentPosition[0]]
                    floorMap[robotPosition[1]][robotPosition[0]] = '.' 

                    robotPosition[0] += move[0]
                    robotPosition[1] += move[1]
                    floorMap[robotPosition[1]][robotPosition[0]] = '@' 
                    break OuterLoop
                }
            }
        }
    }

    total := 0
    for row, line := range(floorMap) {
        for col, c := range(line) {
            if c == 'O' {
                total += 100 * row + col
            }
        }
    }

    fmt.Println(total)
}

