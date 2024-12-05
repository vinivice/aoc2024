package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
    "strconv"
)

type Node struct {
    neighbors map[string]*Node
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
        clone := slices.Clone(printOrder)

        compare := func(i, j int) bool {
            left := nodes[printOrder[i]]
            _, ok := left.neighbors[printOrder[j]]
            return ok
        }

        sort.Slice(printOrder, compare)
        ok := slices.Compare(clone, printOrder) == 0
        if !ok {
            middle, _ :=  strconv.Atoi(printOrder[(len(printOrder) - 1) / 2])
            total += middle
        }
    }

    fmt.Println(total)
}
