const N = 5

func searchGanjil(T [N]int) bool {
    for i := 0; i < N; i++ {
        if T[i]%2 != 0 {
            return true
        }
    }
    return false
}
