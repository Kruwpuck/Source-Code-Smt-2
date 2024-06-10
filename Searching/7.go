const N int = 1000
func indexSearch(T [N]int, x int, n int) int {
    for i := 0; i < n; i++ {
        if T[i] == x {
            return i
        }
    }
    return -1
}