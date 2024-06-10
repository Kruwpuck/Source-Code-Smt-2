package main

import "fmt"

func main() {
    const jumlahPertandingan = 38
    var golMasukkan, golKemasukan, poin, selisihGol int

    for i := 0; i < jumlahPertandingan; i++ {
        var golM, golK int
        fmt.Scan(&golM, &golK)
        
        golMasukkan += golM
        golKemasukan += golK

        if golM > golK {
            poin += 3
        } else if golM == golK {
            poin += 1
        }
    }

    selisihGol = golMasukkan - golKemasukan

    fmt.Print(poin, " ")
    fmt.Print(golMasukkan, " ")
    fmt.Print(golKemasukan, " ")
    fmt.Print(selisihGol, " ")
}