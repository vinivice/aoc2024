package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    //file, _ := os.Open("day13_input_small")
    file, _ := os.Open("day13_input")

    scanner := bufio.NewScanner(file)

    total := 0
    for scanner.Scan() {
        line := scanner.Text()
        lineParts := strings.Split(line, ",")
        
        coordsX := strings.Index(lineParts[0], "+")
        coordsY := strings.Index(lineParts[1], "+")

        x, _ := strconv.Atoi(lineParts[0][coordsX + 1:])
        y, _ := strconv.Atoi(lineParts[1][coordsY + 1:])

        buttonA := [2]int{x, y}


        scanner.Scan()
        line = scanner.Text()
        lineParts = strings.Split(line, ",")
        
        coordsX = strings.Index(lineParts[0], "+")
        coordsY = strings.Index(lineParts[1], "+")

        x, _ = strconv.Atoi(lineParts[0][coordsX + 1:])
        y, _ = strconv.Atoi(lineParts[1][coordsY + 1:])

        buttonB := [2]int{x, y}

        scanner.Scan()
        line = scanner.Text()
        lineParts = strings.Split(line, ",")
        
        coordsX = strings.Index(lineParts[0], "=")
        coordsY = strings.Index(lineParts[1], "=")

        x, _ = strconv.Atoi(lineParts[0][coordsX + 1:])
        y, _ = strconv.Atoi(lineParts[1][coordsY + 1:])

        prize := [2]int{x, y}

        scanner.Scan()

        detS := buttonA[0] * buttonB[1] - buttonB[0] * buttonA[1]
        detA := prize[0] * buttonB[1] - buttonB[0] * prize[1]
        detB := buttonA[0] * prize[1] - prize[0] * buttonA[1]

        if detA % detS != 0 || detB % detS != 0 {
            continue
        }

        A := detA / detS
        B := detB / detS 

        if A < 0 || B < 0 {
            continue
        }

        total += 3 * A + B
    }
    fmt.Println(total)
}
