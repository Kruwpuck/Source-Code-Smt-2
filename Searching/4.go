const N int = 1000
func countThis(T [N]int, total int, ini int) int {
    count := 0
    for i := 0; i < total; i++ {
        if T[i] == ini {
            count++
        }
    }
    return count
}