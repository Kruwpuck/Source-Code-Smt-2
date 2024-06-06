func aset(n int){
    /*I.S Banyak projek terdefinisi
    F.S Menampilkan jumlah pendapatan setiap programmer dan total aset yang terkumpul*/
    var x, pendapatan,jumpen, aset,jumaset ,i int
    
    for i = 1; i <= n; i++{
        fmt.Scan(&x)
        pendapatan = x/3
        jumpen = jumpen + pendapatan
        aset = x%3
        jumaset = jumaset + aset
        
        
    }
      fmt.Printf("%d %d", jumpen, jumaset)
}