package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    fmt.Println("What is your name?")

    // Read name
    reader := bufio.NewReader(os.Stdin)
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    if strings.ToLower(name) == "alice" || strings.ToLower(name) == "bob" {
        fmt.Printf("Hi, %v!\n", name)
    } else {
        fmt.Printf("Hi there!\n")
    }
}
