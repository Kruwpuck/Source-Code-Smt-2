func cekKesamaanKarakter(c1, c2 byte) bool {
    var huruf bool
    huruf = (c1 >= 65 && c1 <= 90) || (c1 >= 97 && c1 <= 122)
    if huruf {
        // Jika keduanya huruf, periksa kesamaan karakter lowercase dan uppercase
        return c1 == c2 || c1-32 == c2 || c2-32 == c1
    } else {
        // Jika salah satu atau kedua karakter bukan huruf, cukup periksa kesamaan langsung
        return c1 == c2
    }
}
