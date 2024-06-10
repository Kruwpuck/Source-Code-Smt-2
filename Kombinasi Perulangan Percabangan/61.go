package main

import "fmt"

func main() {
    var nKendaraan, pendapatan, jumlahA, jumlahB, jumlahC int

    fmt.Scan(&nKendaraan)


    for i := 0; i < nKendaraan; i++ {
        var jenis string
        fmt.Scan(&jenis)
        switch jenis {
        case "A":
            pendapatan += 10000
            jumlahA++
        case "B":
            pendapatan += 23000
            jumlahB++
        case "C":
            pendapatan += 45000
            jumlahC++
        }
    }

    fmt.Print(pendapatan, " ")
    fmt.Print(jumlahA, " ")
    fmt.Print(jumlahB, " ")
    fmt.Print(jumlahC, " ")
}