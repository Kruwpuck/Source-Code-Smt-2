func kelipatanTiga(jumlah *int){
    /*I.S. jumlah diinisialisasi dengan nilai 0
    F.S. jumlah telah berisi jumlah dari bilangan-bilangan ganjil kelipatan 3
    */
    var x int 
    fmt.Scan(&x)
    for x >= 0{
        if x % 3 == 0 && x % 2 != 0{
            *jumlah += x
        }
        fmt.Scan(&x)
    }
}