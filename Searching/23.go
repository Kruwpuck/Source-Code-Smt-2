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

func primaTerkecil(T [N]int, total int) int {
    /*
    algoritma untuk menampilkan nilai bilangan prima paling kecil 
    note :  mengembalikan nilai integer bilangan prima 
    */
    prima := N // Inisialisasi prima dengan nilai N
    for i := 0; i < total; i++ {
        if isPrima(T[i]) && T[i] < prima {
            prima = T[i]
        }
    }
    // Jika tidak ada bilangan prima dalam array, kembalikan nilai N
    if prima == N {
        return N
    }
    return prima
}
