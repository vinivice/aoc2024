package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
)

var inputs map[string][]string
var wrongs map[string]bool

func checkBasic(input string) bool {
    op := inputs[input]
    if (op[0][0] == 'x' && op[2][0] == 'y') || (op[0][0] == 'y' && op[2][0] == 'x') {
        return true
    }
    wrongs[input] = true
    return false
}

func checkCarry(input string) bool {
    op := inputs[input]
    left := inputs[op[0]]
    right := inputs[op[2]]

    if op[1] != "OR" {
        wrongs[input] = true
        return false
    }
    
    switch {
    case left[1] == "AND" && right[1] == "AND":
        return true
    case left[1] == "AND":
        wrongs[op[2]] = true
        return false
    case right[1] == "AND":
        wrongs[op[0]] = true
        return false
    }

    return false
}

func checkZ(input string) bool {
    op, _ := inputs[input]

    if op[1] != "XOR" {
        wrongs[input] = true
        return false
    }

    child1 := op[0]
    child2 := op[2]

    switch {
    case inputs[child1][1] == "XOR" && inputs[child2][1] == "OR":
        return checkBasic(child1) && checkCarry(child2)
    case inputs[child1][1] == "OR" && inputs[child2][1] == "XOR":
        return checkBasic(child2) && checkCarry(child1)
    case inputs[child1][1] == "XOR" && inputs[child2][1] == "XOR":
        checkBasic(child1) 
        checkBasic(child2)
        return false
    case inputs[child1][1] == "OR" && inputs[child2][1] == "OR":
        checkCarry(child1)
        checkCarry(child2)
        return false
    case inputs[child1][1] == "XOR":
        checkBasic(child1)
        wrongs[child2] = true
        return false
    case inputs[child2][1] == "XOR":
        wrongs[child1] = true
        checkBasic(child2)
        return false
    case inputs[child1][1] == "OR":
        checkCarry(child1)
        wrongs[child2] = true
        return false
    case inputs[child2][1] == "OR":
        wrongs[child1] = true
        checkCarry(child2)
        return false
    default:
        return false
    }
}

func main() {
    //file, _ := os.Open("day24_input_small2")
    file, _ := os.Open("day24_input")

    inputs = make(map[string][]string, 0)
    wrongs = make(map[string]bool, 0)

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
   
    for scanner.Scan() {
        line := scanner.Text()

        lineParts := strings.Split(line, " ")
        inputs[lineParts[4]] = lineParts[:3]
    }


    var i byte
    //z00, z01 and z45 manually checked and correct
    for i = 2; i < 45; i++ {
        zRegister := string([]byte{'z', 48 + i / 10, 48 + i % 10})
        checkZ(zRegister)
    }

    sortedOutput := make([]string, 0)
    for k := range(wrongs) {
        sortedOutput = append(sortedOutput, k)
    }

    slices.Sort(sortedOutput)

    var out bytes.Buffer

    for _, v := range(sortedOutput) {
        out.WriteString(v)
        out.WriteString(",")
    }

    out.Truncate(out.Len() - 1)

    fmt.Println(out.String())
}
