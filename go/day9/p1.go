package main

import (
	"fmt"
	"os"
)

func main() {
    //data, _ := os.ReadFile("day9_input_small")
    data, _ := os.ReadFile("day9_input")

    data = data[:len(data) - 1]
    for i := range(data) {
        data[i] -= 48
    }

    memorySize := 0
    for _, v := range(data) {
        memorySize += int(v)
    }

    memory := make([]int, memorySize)

    file := true
    id := 0
    address := 0
    for i := 0; i < len(data); i++ {
        if !file {
            address += int(data[i])
            file = true
            continue
        }

        for j := 0; j < int(data[i]); j++ {
            memory[address] = id
            address += 1
        }
        id++
        file = false
    }

    leftmostEmptyBlockAddress := int(data[0])
    rightmostFileBlockAddress := len(memory) - 1

    for leftmostEmptyBlockAddress < rightmostFileBlockAddress {
        for memory[leftmostEmptyBlockAddress] != 0 {
            leftmostEmptyBlockAddress++
        }

        for memory[rightmostFileBlockAddress] == 0 {
            rightmostFileBlockAddress--
        }

        memory[leftmostEmptyBlockAddress] = memory[rightmostFileBlockAddress]
        memory[rightmostFileBlockAddress] = 0
        leftmostEmptyBlockAddress++
        rightmostFileBlockAddress--
    }

    total := 0
    for i, v := range(memory) {
        total += i*v
    }

    fmt.Println(total)
}
