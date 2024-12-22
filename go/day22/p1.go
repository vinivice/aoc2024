package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const nITERATIONS = 2000

const MOD = 16777216
func calculateSecret(seed int) int {
    s := seed
    s = ((s << 6) ^ s) & (MOD - 1)
    s = ((s >> 5) ^ s) & (MOD - 1)
    s = ((s << 11) ^ s) & (MOD - 1)

    return s
}

func main() {
    //file, _ := os.Open("day22_input_small")
    file, _ := os.Open("day22_input")

    scanner := bufio.NewScanner(file)

    total := 0
    for scanner.Scan() {
        line := scanner.Text()

        secret, _ := strconv.Atoi(line)
        for i := 0; i < nITERATIONS; i++ {
            secret = calculateSecret(secret)
        }

        total += secret
    }

    fmt.Println(total)
}
