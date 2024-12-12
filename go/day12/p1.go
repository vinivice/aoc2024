package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    //file, _ := os.Open("day12_input_small1")
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

    areas := make(map[rune]int, 0)
    perimeters := make(map[rune]int, 0)
    nRows := len(regionsMap)
    nCols := len(regionsMap[0])
    for row := 0; row < nRows; row++ {
        for col := 0; col < nCols; col++ {
            region := regionsMap[row][col]

            areas[region] += 1

            if row - 1 < 0 || regionsMap[row - 1][col] != region {
                perimeters[region] += 1
            }
            if row + 1 >= nRows  || regionsMap[row + 1][col] != region {
                perimeters[region] += 1
            }
            if col - 1 < 0 || regionsMap[row][col - 1] != region {
                perimeters[region] += 1
            }
            if col + 1 >= nCols  || regionsMap[row][col + 1] != region {
                perimeters[region] += 1
            }
        }
    }

    total := 0
    for k := range(areas) {
        total += areas[k] * perimeters[k]
    }

    fmt.Println(total)
}
