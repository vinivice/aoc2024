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

func getWordCount(data []byte, width int, height int, row int, col int) int {
    mas := [3]byte{'M', 'A', 'S'}
    word := [3]byte{0, 0, 0}
    total := 0

    word[0] = getLetter(data, width, height, row, col + 1)
    word[1] = getLetter(data, width, height, row, col + 2)
    word[2] = getLetter(data, width, height, row, col + 3)
    if word == mas {
        total += 1
    }
    word = [3]byte{0, 0, 0}

    word[0] = getLetter(data, width, height, row, col - 1)
    word[1] = getLetter(data, width, height, row, col - 2)
    word[2] = getLetter(data, width, height, row, col - 3)
    if word == mas {
        total += 1
    }
    word = [3]byte{0, 0, 0}

    word[0] = getLetter(data, width, height, row + 1, col)
    word[1] = getLetter(data, width, height, row + 2, col)
    word[2] = getLetter(data, width, height, row + 3, col)
    if word == mas {
        total += 1
    }
    word = [3]byte{0, 0, 0}

    word[0] = getLetter(data, width, height, row - 1, col)
    word[1] = getLetter(data, width, height, row - 2, col)
    word[2] = getLetter(data, width, height, row - 3, col)
    if word == mas {
        total += 1
    }
    word = [3]byte{0, 0, 0}

    word[0] = getLetter(data, width, height, row + 1, col + 1)
    word[1] = getLetter(data, width, height, row + 2, col + 2)
    word[2] = getLetter(data, width, height, row + 3, col + 3)
    if word == mas {
        total += 1
    }
    word = [3]byte{0, 0, 0}

    word[0] = getLetter(data, width, height, row + 1, col - 1)
    word[1] = getLetter(data, width, height, row + 2, col - 2)
    word[2] = getLetter(data, width, height, row + 3, col - 3)
    if word == mas {
        total += 1
    }
    word = [3]byte{0, 0, 0}

    word[0] = getLetter(data, width, height, row - 1, col - 1)
    word[1] = getLetter(data, width, height, row - 2, col - 2)
    word[2] = getLetter(data, width, height, row - 3, col - 3)
    if word == mas {
        total += 1
    }
    word = [3]byte{0, 0, 0}

    word[0] = getLetter(data, width, height, row - 1, col + 1)
    word[1] = getLetter(data, width, height, row - 2, col + 2)
    word[2] = getLetter(data, width, height, row - 3, col + 3)
    if word == mas {
        total += 1
    }

    return total
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
            if curr == 'X' {
                total += getWordCount(data, width, height, row, col)
            }
        }
    }
    
    fmt.Println(total)
}

