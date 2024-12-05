package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
    neighbors map[string]*Node
}

func getValidListMiddle(nodes map[string]*Node, printOrder []string) int {
    for i := 0; i < len(printOrder) - 1; i++ {
        curr, _ := nodes[printOrder[i]]
        _, ok := curr.neighbors[printOrder[i + 1]]
        if !ok {
            return 0
        }
    }

    middle, _ :=  strconv.Atoi(printOrder[(len(printOrder) - 1) / 2])
    return middle
}

func main() {
    //file, _ := os.Open("day5_input_small")
    file, _ := os.Open("day5_input")


    scanner := bufio.NewScanner(file)

    nodes := make(map[string]*Node)

    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            break
        }
        pageNumbers := strings.Split(line, "|")
        left := pageNumbers[0]
        right := pageNumbers[1]

        leftNode, ok := nodes[left]
        if !ok {
            leftNode = &Node{make(map[string]*Node)}
            nodes[left] = leftNode
        }

        rightNode, ok := nodes[right]
        if !ok {
            rightNode = &Node{make(map[string]*Node)}
            nodes[right] = rightNode
        }

        leftNode.neighbors[right] = rightNode

    }

    total := 0
    for scanner.Scan() {
        line := scanner.Text()

        printOrder := strings.Split(line, ",")
        total += getValidListMiddle(nodes, printOrder)
    }

    fmt.Println(total)
}
