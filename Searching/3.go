const N int = 1000

func hereThere(T [N]int, total int, inikah int) int {
    for i := 0; i < total; i++ {
        if T[i] >= inikah {
            return T[i]
        }
    }
    return -1
}