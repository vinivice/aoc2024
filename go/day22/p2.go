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
    //file, _ := os.Open("day22_input_small2")
    file, _ := os.Open("day22_input")

    scanner := bufio.NewScanner(file)

    totalEarnings := make(map[[4]int]int)

    for scanner.Scan() {
        line := scanner.Text()

        secret, _ := strconv.Atoi(line)
        previousSecret := secret
        diffs := [nITERATIONS]int{}
        prices := [nITERATIONS]int{}
        for i := 0; i < nITERATIONS; i++ {
            nextSecret := calculateSecret(secret)
            previousSecret = secret
            secret = nextSecret
            diffs[i] = (secret % 10) - (previousSecret % 10)
            prices[i] = secret % 10
        }

        buyerEarnings := make(map[[4]int]int)
        for i := 3; i < nITERATIONS; i++ {
            key := [4]int{diffs[i - 3], diffs[i - 2], diffs[i - 1], diffs[i - 0]}

            _, ok := buyerEarnings[key]
            if ok {
                continue
            }

            buyerEarnings[key] = prices[i]
        }

        for k, v := range(buyerEarnings) {
            totalEarnings[k] += v
        }
    }

    total := 0
    for _, v := range(totalEarnings) {
        if v > total {
            total = v
        }
    }

    fmt.Println(total)
}
