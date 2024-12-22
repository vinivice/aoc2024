package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Node struct {
    val byte
    input []byte
    visited []byte
}

type Key struct {
    neighbors string
}

func getBestNextInputs(inputString string, pad map[byte]Key) [][]byte {
    input := make([]byte, 0)
    input = append(input, 'A')
    for _, v := range(inputString) {
        input = append(input, byte(v))
    }

    bestInputs := make([][]byte, 0)
    for i := 0; i < len(input) - 1; i++ {
        startVal := input[i]
        start := Node{startVal, make([]byte, 0), make([]byte, 0)}
        target := input[i + 1]

        bestPartialInput := make([][]byte, 0)
        bestPartialInputLen := 1000000000

        checkList := make([]Node, 0)
        checkList = append(checkList, start)
        for len(checkList) > 0 {
            currentNode := checkList[0]
            checkList = checkList[1:]

            if len(currentNode.input) > bestPartialInputLen {
                break
            }

            if currentNode.val == target {
                bestPartialInputLen = len(currentNode.input)
                input := make([]byte, 0)
                input = append(input, currentNode.input...)
                input = append(input, 'A')
                bestPartialInput = append(bestPartialInput, input)
                continue
            }

            key, ok := pad[currentNode.val]
            if !ok {
                panic(1)
            }

            for i := 0; i < len(key.neighbors); i += 2 {
                dir := key.neighbors[i]
                val := key.neighbors[i + 1]
                valIdx := slices.Index(currentNode.visited, val)
                if valIdx < 0 {
                    input := make([]byte, 0)
                    input = append(input, currentNode.input...) 
                    input = append(input, dir)
                    visited := make([]byte, 0)
                    visited = append(visited, currentNode.visited...)
                    visited = append(visited, val)
                    node := Node{val, input, visited}
                    checkList = append(checkList, node)
                }
            }
        }
        tempBestInputs := make([][]byte, 0)
        if len(bestInputs) == 0 {
            tempBestInputs = append(tempBestInputs, bestPartialInput...)
        } else {
            input := make([]byte, 0)
            for _, bp := range(bestInputs) {
                for _, bpp := range(bestPartialInput) {
                    input = make([]byte, 0)
                    input = append(input, bp...)
                    input = append(input, bpp...)
                    tempBestInputs = append(tempBestInputs, input)
                }
            }
        }

        bestInputs = tempBestInputs
    }

    return bestInputs
}

func main() {
    //file, _ := os.Open("day21_input_small")
    file, _ := os.Open("day21_input")

    scanner := bufio.NewScanner(file)

    numPad := make(map[byte]Key, 0)

    numPad['7'] = Key{">8v4"}
    numPad['8'] = Key{"<7v5>9"}
    numPad['9'] = Key{"<8v6"}
    numPad['4'] = Key{"^7v1>5"}
    numPad['5'] = Key{"^8<4>6v2"}
    numPad['6'] = Key{"^9<5v3"}
    numPad['1'] = Key{"^4>2"}
    numPad['2'] = Key{"<1^5>3v0"}
    numPad['3'] = Key{"<2^6vA"}
    numPad['0'] = Key{"^2>A"}
    numPad['A'] = Key{"<0^3"}

    /*
    +---+---+---+
    | 7 | 8 | 9 |
    +---+---+---+
    | 4 | 5 | 6 |
    +---+---+---+
    | 1 | 2 | 3 |
    +---+---+---+
        | 0 | A |
        +---+---+
    */

    keyPad := make(map[byte]Key, 0)

    keyPad['^'] = Key{">Avv"}
    keyPad['A'] = Key{"<^v>"}
    keyPad['<'] = Key{">v"}
    keyPad['v'] = Key{"<<>>^^"}
    keyPad['>'] = Key{"<v^A"}

    /*
        +---+---+
        | ^ | A |
    +---+---+---+
    | < | v | > |
    +---+---+---+
    */

    total := 0
    for scanner.Scan() {
        line := scanner.Text()

        bestInputs1 := getBestNextInputs(line, numPad)

        lenSequence := 10000000000
        for _, bp1 := range(bestInputs1) {
            bestInputs2 := getBestNextInputs(string(bp1), keyPad)
            for _, bp2 := range(bestInputs2) {
                bestInputs3 := getBestNextInputs(string(bp2), keyPad)
                _ = bestInputs3
                for _, bp3 := range(bestInputs3) {
                    if lenSequence > len(bp3) {
                        lenSequence = len(bp3)
                    }
                }
            }
        }

        val, _ := strconv.Atoi(line[:len(line) - 1])
        total += val * lenSequence
    }

    fmt.Println(total)
}
