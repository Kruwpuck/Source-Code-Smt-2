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

func primaTerbesar(T [N]int, total int) int {
    /*
    algoritma untuk menampilkan nilai bilangan proma paling besar 
    note :  mengembalikan nilai integer bilangan prima 
    */
    prima := 0 // Inisialisasi prima dengan nilai 0
    for i := 0; i < total; i++ {
        if isPrima(T[i]) && T[i] > prima {
            prima = T[i]
        }
    }
    return prima
}
