package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var A, B, C int

func getComboOperandValue(operand int) int {
    switch operand {
    case 0,1,2,3: 
        return operand
    case 4:
        return A
    case 5:
        return B
    case 6:
        return C
    default:
        panic(1)
    }
}

func main() {
    //file, _ := os.Open("day17_input_small")
    file, _ := os.Open("day17_input")

    scanner := bufio.NewScanner(file)

    scanner.Scan()
    A, _ = strconv.Atoi( strings.Split(scanner.Text(), ": ")[1])

    scanner.Scan()
    B, _ = strconv.Atoi( strings.Split(scanner.Text(), ": ")[1])

    scanner.Scan()
    C, _ = strconv.Atoi( strings.Split(scanner.Text(), ": ")[1])

    scanner.Scan()

    scanner.Scan()
    programSrc := strings.Split(scanner.Text(), ": ")[1]

    program := make([]int, 0)
    for _, v := range(strings.Split(programSrc, ",")) {
        c, _ := strconv.Atoi(v)
        program = append(program, c)
    }

    instructionPointer := 0 
    var out bytes.Buffer
    for {
        if instructionPointer >= len(program) {
            break
        }

        opcode := program[instructionPointer]
        operand := program[instructionPointer + 1]

        switch opcode {
        case 0:
            num := A
            denExp := getComboOperandValue(operand)
            A = num >> denExp
        case 1:
            B ^= operand
        case 2:
            B = getComboOperandValue(operand) % 8
        case 3:
            if A != 0 {
                instructionPointer = operand
                continue
            }
        case 4:
            B ^= C
        case 5:
            res := getComboOperandValue(operand) % 8
            out.WriteRune(rune(res + 48))
            out.WriteRune(',')
        case 6:
            num := A
            denExp := getComboOperandValue(operand)
            B = num >> denExp
        case 7:
            num := A
            denExp := getComboOperandValue(operand)
            C = num >> denExp
        }

        instructionPointer += 2
    }

    out.Truncate(out.Len() - 1)
    fmt.Println(out.String())
}
