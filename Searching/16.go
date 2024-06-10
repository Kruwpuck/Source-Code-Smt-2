const N int = 1000

func ganjilTerbesar(T [N]int, jumlah int) int {
    // Inisialisasi variabel untuk menyimpan bilangan ganjil terbesar
    maxGanjil := -1

    // Iterasi melalui array T
    for i := 0; i < jumlah; i++ {
        // Periksa apakah nilai di indeks ke-i adalah bilangan ganjil dan lebih besar dari maxGanjil
        if T[i]%2 != 0 && T[i] > maxGanjil {
            maxGanjil = T[i]
        }
    }

    // Kembalikan bilangan ganjil terbesar yang ditemukan atau -1 jika tidak ada bilangan ganjil
    return maxGanjil
}
