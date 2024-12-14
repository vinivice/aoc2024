package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    //file, _ := os.Open("day14_input_small")
    file, _ := os.Open("day14_input")

    scanner := bufio.NewScanner(file)

    timePassed := 100
    //dimensions := [2]int{11, 7}
    dimensions := [2]int{101, 103}

    quadrants := [4]int{0, 0, 0, 0}
    for scanner.Scan() {
        line := scanner.Text()
        lineParts := strings.Split(line, " ")

        tempP := strings.Split(strings.Split(lineParts[0], "=")[1], ",")
        tempV := strings.Split(strings.Split(lineParts[1], "=")[1], ",")

        var P, V, finalP [2]int

        P[0], _ = strconv.Atoi(tempP[0])
        P[1], _ = strconv.Atoi(tempP[1])

        V[0], _ = strconv.Atoi(tempV[0])
        V[1], _ = strconv.Atoi(tempV[1])

        finalP[0] = ((P[0] + timePassed * V[0]) % dimensions[0] + dimensions[0]) % dimensions[0]
        finalP[1] = ((P[1] + timePassed * V[1]) % dimensions[1] + dimensions[1]) % dimensions[1]

        if finalP[0] > dimensions[0] / 2 && finalP[1] < dimensions[1] / 2 {
            quadrants[0] += 1
            continue
        }
        if finalP[0] < dimensions[0] / 2 && finalP[1] < dimensions[1] / 2 {
            quadrants[1] += 1
            continue
        }
        if finalP[0] < dimensions[0] / 2 && finalP[1] > dimensions[1] / 2 {
            quadrants[2] += 1
            continue
        }
        if finalP[0] > dimensions[0] / 2 && finalP[1] > dimensions[1] / 2 {
            quadrants[3] += 1
            continue
        }
    }

    total := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
    fmt.Println(total)
}
