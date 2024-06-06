//isi prosedur saja
func jumlahGenapKelipatan4(bilangan int) {
    /*I.S. terdefinisi bilangan merupakan bilangan bulat
    F.S. program mengeluarkan jumlah dari bilangan kelipatan 4 yang diinputkan
    */
    var jumlah int
    
    for bilangan >= 0{
        if bilangan % 2 == 0 && bilangan % 4 == 0{
          jumlah += bilangan  
        }
        fmt.Scan(&bilangan)
    }
    fmt.Print(jumlah)
}