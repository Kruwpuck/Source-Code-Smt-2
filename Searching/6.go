const N int = 1000
func yangTertukar(T [N]int, total int) int {
    count := 0
    for i := 0; i < total-1; i++ {
        if T[i] > T[i+1] {
            count++
        }
    }
    return count
}