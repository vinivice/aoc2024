package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    //file, _ := os.Open("day12_input_small3")
    file, _ := os.Open("day12_input")

    scanner := bufio.NewScanner(file)

    gardenMap := make([][]rune, 0)
    regionsMap := make([][]rune, 0)
    for scanner.Scan() {
        line := scanner.Text()

        tempLine := make([]rune, 0)
        for _, c := range(line) {
            tempLine = append(tempLine, c)
        }

        gardenMap = append(gardenMap, tempLine)
        tempRegionLine := make([]rune, len(tempLine))
        regionsMap = append(regionsMap, tempRegionLine)
    }

    nextRegion := rune(1)
    for probeRow, line := range(regionsMap) {
        for probeCol := range(line) {
            if regionsMap[probeRow][probeCol] != 0 {
                continue
            }

            probedPlantType := gardenMap[probeRow][probeCol]

            gardenPlotsToCheck := make(map[[2]int]bool, 0)
            gardenPlotsToCheck[[2]int{probeCol, probeRow}] = true
            for len(gardenPlotsToCheck) > 0 {
                var currGardenPlot [2]int
                for k := range gardenPlotsToCheck {
                   currGardenPlot = k 
                   delete(gardenPlotsToCheck, k)
                   break
                }
                
                row := currGardenPlot[1]
                col := currGardenPlot[0]

                regionsMap[row][col] = nextRegion

                if row - 1 >= 0 && regionsMap[row - 1][col] == 0 && gardenMap[row - 1][col] == probedPlantType {
                    gardenPlotsToCheck[[2]int{col, row - 1}] = true
                }
                if row + 1 < len(regionsMap) && regionsMap[row + 1][col] == 0 && gardenMap[row + 1][col] == probedPlantType {
                    gardenPlotsToCheck[[2]int{col, row + 1}] = true
                }
                if col - 1 >= 0 && regionsMap[row][col - 1] == 0 && gardenMap[row][col - 1] == probedPlantType {
                    gardenPlotsToCheck[[2]int{col - 1, row}] = true
                }
                if col + 1 < len(regionsMap[0]) && regionsMap[row][col + 1] == 0 && gardenMap[row][col + 1] == probedPlantType {
                    gardenPlotsToCheck[[2]int{col + 1, row}] = true
                }
            }
            nextRegion++
        }
    }

    nRows := len(regionsMap)
    nCols := len(regionsMap[0])
    areas := make(map[rune]int, 0)
    corners := make(map[rune]int, 0)
    for row := 0; row < nRows; row++ {
        for col := 0; col < nCols; col++ {
            region := regionsMap[row][col]

            areas[region] += 1

            rowAbove := row - 1
            rowBelow := row + 1
            colLeft := col - 1
            colRight := col + 1

            //number of corners == number of sides
            //convex top left corner and concave bottom right
            n1 := rune(0)
            n2 := rune(0)
            n3 := rune(0)
            if colLeft >= 0 {
                n1 = regionsMap[row][colLeft]
            }
            if colLeft >= 0 && rowAbove >= 0 {
                n2 = regionsMap[rowAbove][colLeft]
            }
            if rowAbove >= 0 {
                n3 = regionsMap[rowAbove][col]
            }
            //convex
            if n1 != region && n3 != region {
                corners[region] += 1
            }
            //concave
            if n1 == region && n2 != region && n3 == region {
                corners[region] += 1
            }

            //convex top right corner and concave bottom left
            n1 = rune(0)
            n2 = rune(0)
            n3 = rune(0)
            if colRight < nCols {
                n1 = regionsMap[row][colRight]
            }
            if colRight < nCols && rowAbove >= 0 {
                n2 = regionsMap[rowAbove][colRight]
            }
            if rowAbove >= 0 {
                n3 = regionsMap[rowAbove][col]
            }
            //convex
            if n1 != region && n3 != region {
                corners[region] += 1
            }
            //concave
            if n1 == region && n2 != region && n3 == region {
                corners[region] += 1
            }

            //convex bottom left corner and concave top right
            n1 = rune(0)
            n2 = rune(0)
            n3 = rune(0)
            if colLeft >= 0 {
                n1 = regionsMap[row][colLeft]
            }
            if colLeft >= 0 && rowBelow < nRows {
                n2 = regionsMap[rowBelow][colLeft]
            }
            if rowBelow < nRows {
                n3 = regionsMap[rowBelow][col]
            }
            //convex
            if n1 != region && n3 != region {
                corners[region] += 1
            }
            //concave
            if n1 == region && n2 != region && n3 == region {
                corners[region] += 1
            }

            //convex bottom right corner and concave top left
            n1 = rune(0)
            n2 = rune(0)
            n3 = rune(0)
            if colRight < nCols {
                n1 = regionsMap[row][colRight]
            }
            if colRight < nCols && rowBelow < nRows {
                n2 = regionsMap[rowBelow][colRight]
            }
            if rowBelow < nRows {
                n3 = regionsMap[rowBelow][col]
            }
            //convex
            if n1 != region && n3 != region {
                corners[region] += 1
            }
            //concave
            if n1 == region && n2 != region && n3 == region {
                corners[region] += 1
            }
        }
    }

    total := 0
    for k := range(areas) {
        total += areas[k] * corners[k]
    }

    fmt.Println(total)
}
