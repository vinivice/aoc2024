package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    file, _ := os.Open("day14_input")

    scanner := bufio.NewScanner(file)

    dimensions := [2]int{101, 103}

    positions := make([][2]int, 0)
    velocities := make([][2]int, 0)

    for scanner.Scan() {
        line := scanner.Text()
        lineParts := strings.Split(line, " ")

        tempP := strings.Split(strings.Split(lineParts[0], "=")[1], ",")
        tempV := strings.Split(strings.Split(lineParts[1], "=")[1], ",")

        var P, V  [2]int

        P[0], _ = strconv.Atoi(tempP[0])
        P[1], _ = strconv.Atoi(tempP[1])

        positions = append(positions, P)

        V[0], _ = strconv.Atoi(tempV[0])
        V[1], _ = strconv.Atoi(tempV[1])

        velocities = append(velocities, V)
    }

    floorMap := make([][103]rune, 101)
    for i := range(floorMap) {
        for j := range(floorMap[i]) {
            floorMap[i][j] = '.'
        }
    }

    treshold := 50
    for elapsedTime := 0; elapsedTime < dimensions[0] * dimensions[1]; elapsedTime++ {
        for index := range(positions) {
            floorMap[positions[index][0]][positions[index][1]] = '.'
            positions[index][0] = ((positions[index][0] + velocities[index][0]) % dimensions[0] + dimensions[0]) % dimensions[0]
            positions[index][1] = ((positions[index][1] + velocities[index][1]) % dimensions[1] + dimensions[1]) % dimensions[1]
            floorMap[positions[index][0]][positions[index][1]] = '*'
        }

        printDiagram := false
        regionsMap := make([][103]bool, 101)
        for _, p := range(positions) {
            if regionsMap[p[0]][p[1]] {
                continue
            }

            checkList := make(map[[2]int]bool, 0)

            checkList[[2]int{p[0], p[1]}] = true 
            clusterSize := 0
            for len(checkList) > 0 {
                var position [2]int
                for k := range checkList {
                    position = k 
                    delete(checkList, k)
                    break
                }

                x := position[0]
                y := position[1]

                regionsMap[x][y] = true
                clusterSize++

                if x - 1 >= 0 && !regionsMap[x - 1][y] && floorMap[x - 1][y] == '*' {
                    checkList[[2]int{x - 1, y}] = true
                }
                if x + 1 < len(regionsMap) && !regionsMap[x + 1][y] && floorMap[x + 1][y] == '*' {
                    checkList[[2]int{x + 1, y}] = true
                }
                if y - 1 >= 0 && !regionsMap[x][y - 1] && floorMap[x][y - 1] == '*' {
                    checkList[[2]int{x, y - 1}] = true
                }
                if y + 1 < len(regionsMap[0]) && !regionsMap[x][y + 1] && floorMap[x][y + 1] == '*' {
                    checkList[[2]int{x, y + 1}] = true
                }
            }

            if clusterSize > treshold {
                printDiagram = true
            }
        }

        if printDiagram {
            for _, v := range(floorMap) {
                fmt.Println(string(v[:]))
            }
            fmt.Println(elapsedTime + 1)
        }
    }
}
