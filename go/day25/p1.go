package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    //file, _ := os.Open("day25_input_small")
    file, _ := os.Open("day25_input")

    scanner := bufio.NewScanner(file)

    keys := make([][5]int, 0)
    locks := make([][5]int, 0)
    for scanner.Scan() {
        isLock := scanner.Text()[0] == '#'
        pins := [5]int{-1, -1, -1, -1, -1}
        for  {
            line := scanner.Text()

            if line == "" {
                break
            }

            for i, v := range(line) {
                if v == '#' {
                    pins[i]++
                }
            }

            scanner.Scan()
        }

        if isLock {
            locks = append(locks, pins)
        } else {
            keys = append(keys, pins)
        }
    }

    total := 0 
    for _, lock := range(locks) {
        KeyLoop:
        for _, key := range(keys) {
            for i := 0; i < 5; i++ {
                if lock[i] + key[i] > 5 {
                    continue KeyLoop
                }
            }
            total++
        }
    }

    fmt.Println(total)
}
