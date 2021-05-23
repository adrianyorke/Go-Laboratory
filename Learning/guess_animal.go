package main

import "fmt"

const favAnimal string = "dog"

func main() {
    var guess string
    // Create an input loop
    for {
        // Ask the user to guess my favorite animal
        fmt.Println("Guess my favorite animal:")
        // Try to read a line of input from the user. Print out the error 0
        if _, err := fmt.Scanln(&guess); err != nil {
            fmt.Printf("%s\n", err)
            return
        }
        // Did they guess the correct animal?
        if favAnimal == guess {
            // They guessed it!
            fmt.Printf("%q is my favorite animal!\n", favAnimal)
            return
        }
        // Wrong! Have them guess again.
        fmt.Printf("Sorry, %q is not my favorite animal. Guess again.\n", guess)
    }
}