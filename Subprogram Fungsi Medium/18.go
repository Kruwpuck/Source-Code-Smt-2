func cekBothGanjil(a int, b int) bool {
    return a%2 != 0 && b%2 != 0
}

func cekBothGenap(a int, b int) bool {
    return a%2 == 0 && b%2 == 0
}

func hasilKali(a int, b int) int {
    return a * b
}

func hasilJumlah(a int, b int) int {
    return a + b
}
func countNumber(a int, b int) int {
    if cekBothGenap(a, b) {
        return hasilKali(a, b)
    } else if cekBothGanjil(a, b) {
        return hasilJumlah(a, b)
    }
    return 0
}