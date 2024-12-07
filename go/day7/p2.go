package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkOperation(target, lh int, inputs []string, operator string) bool {
    if len(inputs) == 0 {
        return target == lh 
    }

    rh, _ := strconv.Atoi(inputs[0])
    exponent := len(inputs[0])
    newLh := 0
    switch operator {
    case "+":
        newLh = lh + rh
    case "*":
        newLh = lh * rh
    case "||":
        newLh = int(math.Pow10(exponent)) * lh + rh
    }

    newInputs := inputs[1:]

    if checkOperation(target, newLh, newInputs, "+") {
        return true
    }
    if checkOperation(target, newLh, newInputs, "*") {
        return true
    }
    
    return checkOperation(target, newLh, newInputs, "||")
}

func checkMatch(target int, inputs []string) int {
    if checkOperation(target, 0, inputs, "+") || checkOperation(target, 0, inputs, "*") || checkOperation(target, 0, inputs, "||") {
        return target
    }

    return 0
    
}

func main() {
    //file, _ := os.Open("day7_input_small")
    file, _ := os.Open("day7_input")

    scanner := bufio.NewScanner(file)


    total := 0
    for scanner.Scan() {
        line := scanner.Text()

        values := strings.Split(line, ":")

        target, _ := strconv.Atoi(values[0])
        inputs := strings.Split(values[1], " ")

        total += checkMatch(target, inputs)
    }

    fmt.Println(total)
}
