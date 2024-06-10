const N int = 1000

func genapTerbesar(T [N]int, jumlah int) int {
    // Inisialisasi variabel untuk menyimpan bilangan genap terbesar
    maxGenap := -1

    // Iterasi melalui array T
    for i := 0; i < jumlah; i++ {
        // Periksa apakah nilai di indeks ke-i adalah bilangan genap dan lebih besar dari maxGenap
        if T[i]%2 == 0 && T[i] > maxGenap {
            maxGenap = T[i]
        }
    }
    // Kembalikan bilangan genap terbesar yang ditemukan atau -1 jika tidak ada bilangan genap
    return maxGenap
}
