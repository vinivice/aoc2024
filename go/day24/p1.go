package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    //file, _ := os.Open("day24_input_small2")
    file, _ := os.Open("day24_input")

    scanner := bufio.NewScanner(file)

    registers := make(map[string]bool)
    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            break
        }

        lineParts := strings.Split(line, ": ")

        registers[lineParts[0]] = lineParts[1] == "1"
    }
   
    inputs := make([][]string, 0)
    for scanner.Scan() {
        line := scanner.Text()

        lineParts := strings.Split(line, " ")
        inputs = append(inputs, lineParts)
    }

    for len(inputs) > 0 {
        input := inputs[0]
        inputs = inputs[1:]

        _, ok1 := registers[input[0]]
        _, ok2 := registers[input[2]]

        if !ok1 || !ok2 {
            inputs = append(inputs, input)
            continue
        }

        switch input[1] {
        case "AND":
            registers[input[4]] = registers[input[0]] && registers[input[2]]
        case "OR":
            registers[input[4]] = registers[input[0]] || registers[input[2]]
        case "XOR":
            registers[input[4]] = registers[input[0]] != registers[input[2]]
        }

    }

    output := 0
    for k, v := range(registers) {
        if k[0] != 'z' {
            continue
        }

        if v {
            id, _ := strconv.Atoi(k[1:])
            output += 1 << id
        }
    }
    fmt.Println(output)
}
