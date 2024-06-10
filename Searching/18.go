const N int = 1000

func lebihKecil(T [N]int, jumlah, target int) int {

    total := 0
    for i := 0; i < jumlah; i++ {
        if T[i] < target {
            total++
        }
    }
    return total
}
