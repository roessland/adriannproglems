package main

import (
    "fmt"
)

func main() {
    for i := 1; i <= 12; i += 1 {
        for j := 1; j <= 12; j += 1 {
            fmt.Printf("%4d", i*j)
        }
        fmt.Printf("\n")
    }
}
