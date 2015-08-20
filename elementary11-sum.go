package main

import (
    "fmt"
    "math/big"
)

// Compute 4*(1 - 1/3 + 1/5 - 1/7 + 1/9 - ...)
func main() {
    sum := big.NewRat(0, 1)
    var sumTimes4 big.Rat

    for i, sign := int64(1), int64(1); i < 10000000; i, sign = i+2, -sign {
        part := big.NewRat(sign, i)
        sum.Add(sum, part)

        if (i+1) % 1000 == 0 {
            sumTimes4.Mul(big.NewRat(4,1), sum)
            fmt.Printf("%v, %v\n", i+1, sumTimes4.FloatString(10))
        }
    }

}
