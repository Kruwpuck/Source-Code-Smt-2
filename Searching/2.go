const N int = 1000

func insertOne(T *[N]int, akhir *int, baru int) {
    // Inisialisasi variabel
    found := -1
    i := 0

    // Cari posisi untuk menyisipkan nilai baru
    for i <= *akhir && found == -1 {
        if T[i] > baru {
            found = i
        }
        i = i + 1
    }

    // Jika tidak ditemukan posisi yang lebih besar, maka masukkan di akhir
    if found == -1 {
        found = *akhir + 1
    }

    // Perbarui indeks akhir
    *akhir = *akhir + 1

    // Geser elemen-elemen setelah posisi ditemukan
    for i = *akhir; i > found; i-- {
        T[i] = T[i-1]
    }

    // Masukkan elemen baru di posisi yang ditemukan
    T[found] = baru
}