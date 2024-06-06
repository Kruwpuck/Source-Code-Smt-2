func procedure(){
    /*I.S.  -
    F.S. terdapat 10 input bilangan bulat atau berakhir jika bilangan yang dibaca
    lebih besar dari bilangan terakhir masuk. Keluarkan Bilangan bulat yang 
    menyatakan berapa banyak bilangan bulat diinput.
    Petunjuk: Lakukan proses masukan menggunakan perulangan di dalam prosedur

}
    */
    var b1, b2, i int
    var stop bool
    
    stop = i == 10
    fmt.Scan(&b1)
    for !stop{
        i++
        fmt.Scan(&b2)
        stop = i == 10 || b2 > b1
        for !stop{
            i++
            fmt.Scan(&b1)
            stop = i == 10 || b1 > b2
            break
        }
    }
    fmt.Print(i)
    
    
}

