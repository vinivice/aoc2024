package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "slices"
)

type Node struct {
    val string
    neighbors []*Node
}

var cycles map[[3]string]bool
func findCycles(start *Node, root string, maxDepth int, path []string) {
    if maxDepth == 0 {
        if start.val == root {
            slices.SortFunc(path, func(a, b string) int {
                if a[0] != b[0] {
                    return int(a[0]) - int(b[0])
                }
                return int(a[1]) - int(b[1])
            })

            key := [3]string{path[0], path[1], path[2]}
            cycles[key] = true
        }
        return
    }

    for _, neighbor := range(start.neighbors) {
        if neighbor.val == root || slices.Index(path, neighbor.val) < 0 {
              currPath := make([]string, 0)
              currPath = append(currPath, path...)
              currPath = append(currPath, start.val)

              findCycles(neighbor, root, maxDepth - 1, currPath)
        }
    }
}

func main() {
    cycles = make(map[[3]string]bool, 0)
    //file, _ := os.Open("day23_input_small")
    file, _ := os.Open("day23_input")

    scanner := bufio.NewScanner(file)

    nodes := make(map[string]*Node)
    for scanner.Scan() {
        line := scanner.Text()

        lineParts := strings.Split(line, "-")

        node1, ok := nodes[lineParts[0]]
        if !ok {
            node1 = &Node{lineParts[0], make([]*Node, 0)}
            nodes[lineParts[0]] = node1
        }

        node2, ok := nodes[lineParts[1]]
        if !ok {
            node2 = &Node{lineParts[1], make([]*Node, 0)}
            nodes[lineParts[1]] = node2
        }

        node1.neighbors = append(node1.neighbors, node2)
        node2.neighbors = append(node2.neighbors, node1)
    }

    for nodeValue, node := range(nodes) {
        if nodeValue[0] == 't' {
            findCycles(node, nodeValue, 3, nil)
        }
    }

    fmt.Println(len(cycles))
}
