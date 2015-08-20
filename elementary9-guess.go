package main

import (
    "fmt"
    "math/rand"
)

func main() {
    secret := rand.Int() % 500

    oldguess := secret
    var guess int

    fmt.Printf("Enter your guess: ")
    fmt.Scanf("%d", &guess)
    tries := 1
    for guess != secret {
        if guess < secret {
            fmt.Println("Too low!")
        } else if guess > secret {
            fmt.Println("Too high!")
        }
        fmt.Printf("Enter your guess: ")
        fmt.Scanf("%d", &guess)

        if guess != oldguess {
            tries++
        }
        oldguess = guess
    }
    fmt.Printf("That's correct, the number was %v\n", secret)
    fmt.Printf("You used %v tries", tries)
}
