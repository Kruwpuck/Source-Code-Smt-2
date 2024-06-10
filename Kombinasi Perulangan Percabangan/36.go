package main

import "fmt"

func main() {
    var dadu1, dadu2, countTidakSama int
    for {
   
        fmt.Scan(&dadu1, &dadu2)
        
        if (dadu1%2 == 0 && dadu2%2 != 0) || (dadu1%2 != 0 && dadu2%2 == 0) {
            countTidakSama++
        } else if (dadu1%2 == 0 && dadu2%2 == 0) || (dadu1%2 != 0 && dadu2%2 != 0) {
            break // Berhenti jika muncul angka dadu sama-sama genap atau sama-sama ganjil
        }
    }

    fmt.Print(countTidakSama)
}