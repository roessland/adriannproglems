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
    text, _ := reader.ReadString('\n')
    text = strings.TrimSpace(text)

    fmt.Printf("Hi, %v!\n", text)
}
