type arrOfInt [N]int

const N int = 1000

func isPrima(num int) bool {
    if num <= 1 {
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}

func prima(T arrOfInt, total int, length *int, P *arrOfInt) {
    *length = -1
    for i := 0; i < total; i++ {
        if isPrima(T[i]) {
            *length++
            (*P)[*length] = T[i]
        }
    }
}
