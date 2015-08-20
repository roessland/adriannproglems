package main

import (
    "fmt"
)

func main() {
    // Read number
    fmt.Printf("Type a number:")
    var num int
    fmt.Scan(&num)


    // Sum number
    var sum int = 0
    for i := 1; i <= num; i = i+1 {
        sum += i
    }

    // Print result
    fmt.Printf("The sum of numbers from 1 to %d is %d\n", num, sum)

}
