const N int = 7

func inMax(T [N]int, awal int, akhir int) int {
    // Menginisialisasi nilai max dengan nilai pada indeks awal
    max := T[awal]
    // Mengulang dari indeks awal hingga akhir
    for i := awal; i <= akhir; i++ {
        if T[i] > max {
            max = T[i]
        }
    }
    return max
}

