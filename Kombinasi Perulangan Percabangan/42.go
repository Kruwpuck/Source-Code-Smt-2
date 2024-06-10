package main

import "fmt"

func main() {
    var input string
    fmt.Scan(&input)

    var found bool

    // Periksa apakah huruf 'k' atau 'q' ada dalam input
    for _, char := range input {
        if char == 'k' || char == 'q' {
            found = true
            break
        }
    }

    fmt.Println(found)
}