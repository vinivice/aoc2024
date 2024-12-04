package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseNumber(input []byte, position int) (int, int) {
    for i := 0; ; i++ {
        if input[position + i] < 48 || input[position + i] > 57 {
            number, _ := strconv.Atoi(string(input[position:position + i]))
            return number, i
        }
    }
}

func parseMul(input []byte, position int) (int, int) {
    if string(input[position:position+4]) != "mul(" {
        return 0, position
    }

    position = position + 4
    X, length := parseNumber(input, position)
    position += length
    
    if string(input[position]) != "," {
        return 0, position
    }
    position++

    Y, length := parseNumber(input, position)
    position += length

    if string(input[position]) != ")" {
        return 0, position
    }
    return X * Y, position
}

func main() {
    //data, _ := os.ReadFile("day3_input_small")
    //data, _ := os.ReadFile("day3_input_small_2")
    data, _ := os.ReadFile("day3_input")

    do := true
    total := 0
    for i := 0; i < len(data); i++ {
        if do && data[i] == 'm' {
            num, position := parseMul(data, i)
            total += num
            i = position
        }
        if data[i] == 'd' {
            if string(data[i:i + 4]) == "do()" {
                do = true
            }
            if string(data[i:i + 7]) == "don't()" {
                do = false
            }
        }
    }

    fmt.Println(total)
}
