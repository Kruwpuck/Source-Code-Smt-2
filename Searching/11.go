const N = 5
func searchGenap(T [N]int) bool {
    for _, num := range T {
        if num%2 == 0 {
            return true
        }
    }
    return false
}
