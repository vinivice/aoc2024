package main

import (
	"fmt"
	"os"
	"slices"
)

func getLetter(data []byte, width int, height int, row int, col int) byte {
    if row < 0 || row >= width || col < 0 || col >= height {
        return 48
    }

    index := row * (width + 1) + col
    return data[index]
}

func checkX(data []byte, width int, height int, row int, col int) int {
    mmss := [4]byte{'M', 'M', 'S', 'S'}
    smms := [4]byte{'S', 'M', 'M', 'S'}
    ssmm := [4]byte{'S', 'S', 'M', 'M'}
    mssm := [4]byte{'M', 'S', 'S', 'M'}
    corners := [4]byte{0, 0, 0, 0}

    corners[0] = getLetter(data, width, height, row - 1, col - 1)
    corners[1] = getLetter(data, width, height, row - 1, col + 1)
    corners[2] = getLetter(data, width, height, row + 1, col + 1)
    corners[3] = getLetter(data, width, height, row + 1, col - 1)

    if corners == mmss || corners == smms || corners == ssmm || corners == mssm {
        return 1
    }

    return 0
}

func main() {

    data, _ := os.ReadFile("day4_input")
    //data, _ := os.ReadFile("day4_input_small")
    width := slices.Index(data, '\n')
    height := len(data) / width - 1

    total := 0
    for row := 0; row < width; row++ {
        for col := 0; col < height; col++ {
            curr := getLetter(data, width, height, row, col)
            if curr == 'A' {
                total += checkX(data, width, height, row, col)
            }
        }
    }
    
    fmt.Println(total)
}

