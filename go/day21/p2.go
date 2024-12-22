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

type Input struct {
    inputsList [][]byte
    next *Input
}

type RecordKey struct {
    input string
    maxDepth int
}

var bestInputRecord map[RecordKey]int

func (inp *Input) getMinLen(maxDepth int, pad map[byte]Key) int {
    minLen := 1000000000000000000 
    if maxDepth > 0 {
        for _, input := range(inp.inputsList) {
            key := RecordKey{string(input), maxDepth}

            l, ok := bestInputRecord[key]
            if !ok {
                node := getBestNextInputs(string(input), pad)
                l = node.getMinLen(maxDepth - 1, pad)
                bestInputRecord[key] = l
            }

            if l < minLen {
                minLen = l
            }
        }
    } else {
        for _, input := range(inp.inputsList) {
            l := len(input)
            if l < minLen {
                minLen = l
            }
        }
    }

    nextLen := 0
    if inp.next != nil {
        nextLen = inp.next.getMinLen(maxDepth, pad)
    }

    return minLen + nextLen
}
    

func getBestNextInputs(inputString string, pad map[byte]Key) Input {
    input := make([]byte, 0)
    input = append(input, 'A')
    for _, v := range(inputString) {
        input = append(input, byte(v))
    }

    root := Input{make([][]byte, 0), nil}
    currentInput := &root
    previousInput := &root
    for i := 0; i < len(input) - 1; i++ {
        startVal := input[i]
        start := Node{startVal, make([]byte, 0), make([]byte, 0)}
        target := input[i + 1]

        bestGeneratedInputLen := 1000000000000000000

        checkList := make([]Node, 0)
        checkList = append(checkList, start)
        for len(checkList) > 0 {
            currentNode := checkList[0]
            checkList = checkList[1:]

            if len(currentNode.input) > bestGeneratedInputLen {
                break
            }

            if currentNode.val == target {
                bestGeneratedInputLen = len(currentNode.input)
                newInput := make([]byte, 0)
                newInput = append(newInput, currentNode.input...)
                newInput = append(newInput, 'A')
                currentInput.inputsList = append(currentInput.inputsList, newInput)
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
                    newInput := make([]byte, 0)
                    newInput = append(newInput, currentNode.input...) 
                    newInput = append(newInput, dir)
                    visited := make([]byte, 0)
                    visited = append(visited, currentNode.visited...)
                    visited = append(visited, val)
                    node := Node{val, newInput, visited}
                    checkList = append(checkList, node)
                }
            }
        }

        nextInput := Input{make([][]byte, 0), nil}
        currentInput.next = &nextInput
        previousInput = currentInput
        currentInput = currentInput.next
    }

    previousInput.next = nil
    return root
}

func main() {
    bestInputRecord = make(map[RecordKey]int, 0)

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

    const nMIDDLELAYERS = 25
    total := 0
    for scanner.Scan() {
        line := scanner.Text()

        inputs := getBestNextInputs(line, numPad)

        lenInput := inputs.getMinLen(nMIDDLELAYERS, keyPad)

        val, _ := strconv.Atoi(line[:len(line) - 1])
        total += val * lenInput
    }

    fmt.Println(total)
}
