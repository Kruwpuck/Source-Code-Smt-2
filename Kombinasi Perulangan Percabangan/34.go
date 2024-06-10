package main

import "fmt"

func main() {
    var dadu1, dadu2 int
    var countEvenPairs int

    for {
        fmt.Scan(&dadu1, &dadu2)
        
        // Berhenti jika kedua angka ganjil
        if dadu1%2 != 0 && dadu2%2 != 0 {
            break
        }
        
        // Tambahkan ke hitungan jika kedua angka genap
        if dadu1%2 == 0 && dadu2%2 == 0 {
            countEvenPairs++
        }
    }

    fmt.Println(countEvenPairs)
}