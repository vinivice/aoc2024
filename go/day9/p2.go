package main

import (
	"fmt"
	"os"
	"slices"
)

type Metadata struct {
    file bool
    id int
    size int
}

func main() {
    //data, _ := os.ReadFile("day9_input_small")
    data, _ := os.ReadFile("day9_input")

    data = data[:len(data) - 1]
    for i := range(data) {
        data[i] -= 48
    }

    memoryMetadata := make([]Metadata, 0)
    
    file := true
    id := 0
    for i := 0; i < len(data); i++ {
        m := Metadata{file, id, int(data[i])}
        memoryMetadata = append(memoryMetadata, m)
        if file {
            id++
        }
        file = !file
    }

    for filePointer := len(memoryMetadata) - 1; filePointer > 0; filePointer-- {
        if !memoryMetadata[filePointer].file {
            continue
        }

        for emptySpacePointer := 0; emptySpacePointer < filePointer; emptySpacePointer++ {
            if memoryMetadata[emptySpacePointer].file || memoryMetadata[emptySpacePointer].size < memoryMetadata[filePointer].size {
                continue
            }

            currEmptySpace := memoryMetadata[emptySpacePointer]
            currFile := memoryMetadata[filePointer]
            memoryMetadata[filePointer].file = false

            currEmptySpace.size -= currFile.size

            left := memoryMetadata[:emptySpacePointer]
            right := memoryMetadata[emptySpacePointer + 1:]

            memoryMetadata = slices.Concat(left, []Metadata{currFile, currEmptySpace}, right)

            break
        }
    }

    address := 0
    total := 0
    for _, v := range(memoryMetadata) {
        if v.file {
            n := v.size
            total += (n*address + n*(n-1)/2) * v.id
        }
        address += v.size
    }

    fmt.Println(total)
}
