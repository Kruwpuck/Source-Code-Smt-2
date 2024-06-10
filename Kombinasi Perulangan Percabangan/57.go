package main

import "fmt"

func main() {
    var num int

    // Inisialisasi terbesar dan terkecil dengan nilai dari input pertama yang valid
    fmt.Scan(&num)
    if num == 0 {
        fmt.Println("Tidak ada bilangan yang dimasukkan.")
        return
    }

    terbesar := num
    terkecil := num

    for {
        fmt.Scan(&num)
        if num == 0 {
            break
        }
        if num > terbesar {
            terbesar = num
        }
        if num < terkecil {
            terkecil = num
        }
    }

    fmt.Print(terbesar, " ")
    fmt.Print(terkecil, " ")
}