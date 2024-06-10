const N int = 1000

func sudahTerurut(T [N]int, total int) bool {
    for i := 0; i < total-1; i++ {
        if T[i] > T[i+1] {
            return false
        }
    }
    return true
}