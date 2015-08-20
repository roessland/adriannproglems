package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "math/big"
)

func main() {
    // Read number
    fmt.Printf("Type a number:")
    var num int64
    fmt.Scan(&num)

    // Read command
    fmt.Printf("Sum or product:")
    reader := bufio.NewReader(os.Stdin)
    command, _ := reader.ReadString('\n')
    command = strings.TrimSpace(command)

    // Do stuff
    if command == "sum" {
        var sum int64 = 0
        for i := int64(1); i <= num; i = i+1 {
            sum += i
        }
        fmt.Printf("The sum of numbers from 1 to %d is %d\n", num, sum)
    } else if command == "product" {
        var product big.Int
        product.MulRange(int64(1), num)
        fmt.Printf("The product of numbers from 1 to %d is %v\n", num, product.String())
    } else {
        fmt.Errorf("Command %v not recognized", command)
        os.Exit(1)
    }

}
