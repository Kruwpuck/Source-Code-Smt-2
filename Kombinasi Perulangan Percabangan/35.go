package main

import "fmt"

func main() {
    var dadu1, dadu2, countGanjil int
    for {

        fmt.Scan(&dadu1, &dadu2)
        
        if dadu1%2 == 1 && dadu2%2 == 1 {
            countGanjil++
        } else if dadu1%2 == 0 && dadu2%2 == 0 {
            break // Berhenti jika muncul angka dadu sama-sama genap
        }
    }

    fmt.Print(countGanjil)
}