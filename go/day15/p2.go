package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkVerticalMove(floorMap [][]rune, currentPosition [2]int, move [2]int) bool {
    nextPosition := [2]int{currentPosition[0] + move[0], currentPosition[1] + move[1]}

    switch floorMap[nextPosition[1]][nextPosition[0]] {
    case '#':
        return false
    case '[':
        boxLeft := nextPosition
        boxRight := nextPosition
        boxRight[0] += 1
        return checkVerticalMove(floorMap, boxLeft, move) && checkVerticalMove(floorMap, boxRight, move) 
    case ']':
        boxRight := nextPosition
        boxLeft := nextPosition
        boxLeft[0] -= 1
        return checkVerticalMove(floorMap, boxLeft, move) && checkVerticalMove(floorMap, boxRight, move) 
    case '.' :
        return true
    }

    return false 
}

func moveVertically(floorMap [][]rune, currentPosition [2]int, move [2]int) {
    nextPosition := [2]int{currentPosition[0] + move[0], currentPosition[1] + move[1]}

    switch floorMap[nextPosition[1]][nextPosition[0]] {
    case '[' :
        boxLeft := nextPosition
        boxRight := nextPosition
        boxRight[0] += 1
        moveVertically(floorMap, boxLeft, move)
        moveVertically(floorMap, boxRight, move) 

    case ']':
        boxRight := nextPosition
        boxLeft := nextPosition
        boxLeft[0] -= 1
        moveVertically(floorMap, boxLeft, move)
        moveVertically(floorMap, boxRight, move) 
    }

    temp := floorMap[nextPosition[1]][nextPosition[0]]
    floorMap[nextPosition[1]][nextPosition[0]] = floorMap[currentPosition[1]][currentPosition[0]]
    floorMap[currentPosition[1]][currentPosition[0]] = temp
}

func verticalMove(floorMap [][]rune, robotPosition [2]int, move [2]int) [2]int {
    if checkVerticalMove(floorMap, robotPosition, move) {
        moveVertically(floorMap, robotPosition, move)
        robotPosition[0] += move[0]
        robotPosition[1] += move[1]
    }

    return robotPosition
}
func horizontalMove(floorMap [][]rune, robotPosition [2]int, move [2]int) [2]int {
    currentPosition := robotPosition
    OuterLoop:
    for {
        nextPosition := [2]int{currentPosition[0] + move[0], currentPosition[1] + move[1]}

        switch floorMap[nextPosition[1]][nextPosition[0]] {
        case '#':
            break OuterLoop
        case '[', ']':
            currentPosition[0] += 2 * move[0]
            currentPosition[1] += 2 * move[1]
            continue OuterLoop
        case '.':
            for p := nextPosition; p != robotPosition;  {
                prev := p
                prev[0] -= move[0]
                prev[1] -= move[1]
                floorMap[p[1]][p[0]] = floorMap[prev[1]][prev[0]]
                p = prev
            }

            floorMap[robotPosition[1]][robotPosition[0]] = '.' 
            robotPosition[0] += move[0]
            robotPosition[1] += move[1]
            break OuterLoop
        }
    }
    return robotPosition
}


func main() {
    //file, _ := os.Open("day15_input_small1")
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
            switch c {
            case '#':
                currLine = append(currLine, '#')
                currLine = append(currLine, '#')
            case 'O':
                currLine = append(currLine, '[')
                currLine = append(currLine, ']')
            case '.':
                currLine = append(currLine, '.')
                currLine = append(currLine, '.')
            case '@':
                robotPosition = [2]int{2 * col, row}
                currLine = append(currLine, '@')
                currLine = append(currLine, '.')
            }
        }
        floorMap = append(floorMap, currLine)
    }

    for scanner.Scan() {
        line := scanner.Text()
        for _, moveDir := range(line) {
            switch moveDir {
            case '^':
                move := [2]int{0, -1}
                robotPosition = verticalMove(floorMap, robotPosition, move)
            case 'v':
                move := [2]int{0, 1}
                robotPosition = verticalMove(floorMap, robotPosition, move)
            case '<':
                move := [2]int{-1, 0}
                robotPosition = horizontalMove(floorMap, robotPosition, move)
            case '>':
                move := [2]int{1, 0}
                robotPosition = horizontalMove(floorMap, robotPosition, move)
            }
        }
    }
   
    total := 0
    for row, line := range(floorMap) {
        for col, c := range(line) {
            if c == '[' {
                total += 100 * row + col
            }
        }
    }

    fmt.Println(total)
}

