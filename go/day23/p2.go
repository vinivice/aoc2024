package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
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

    for nodeVal, node := range(nodes) {
        findCycles(node, nodeVal, 3, nil)
    }

    allMeshs := make([][]*Node, 0)
    for cycle := range(cycles) {
        mesh := make([]*Node, 0)

        mesh = append(mesh, nodes[cycle[0]])
        mesh = append(mesh, nodes[cycle[1]])
        mesh = append(mesh, nodes[cycle[2]])

        allMeshs = append(allMeshs, mesh)
    }

    checkList := make([][]*Node, 0)
    checkList = append(checkList, allMeshs...)

    fullMeshes := make(map[string][]*Node, 0)
    for len(checkList) > 0 {
        currMesh := checkList[0]
        checkList = checkList[1:]

        neighborsIntersection := make([]*Node, 0)
        neighborsIntersection = append(neighborsIntersection, currMesh[0].neighbors...)

        for _, node := range(currMesh) {
            for i := 0; i < len(neighborsIntersection); i++ {
                index := slices.Index(node.neighbors, neighborsIntersection[i])
                if index < 0 {
                    neighborsIntersection = append(neighborsIntersection[:i], neighborsIntersection[i + 1:]... )
                    i--
                }
            }
        }

        if len(neighborsIntersection) > 0 {
            currMesh = append(currMesh, neighborsIntersection[0])
            checkList = append(checkList, currMesh)
            continue
        } 

        nodes := make([]string, 0)
        for _, node := range(currMesh) {
            nodes = append(nodes, node.val)
        }

        slices.SortFunc(nodes, func(a, b string) int {
            if a[0] != b[0] {
                return int(a[0]) - int(b[0])
            }
            return int(a[1]) - int(b[1])
        })

        key := make([]byte, 0)
        for _, n := range(nodes) {
            key = append(key, n[0])
            key = append(key, n[1])
            key = append(key, ',')
        }

        fullMeshes[string(key[:len(key) - 1])] = currMesh
    }

    biggestLanParty := ""
    for meshNodesList := range(fullMeshes) {
        if len(meshNodesList) > len(biggestLanParty) {
            biggestLanParty = meshNodesList
        }
    }

    fmt.Println(biggestLanParty)
}
