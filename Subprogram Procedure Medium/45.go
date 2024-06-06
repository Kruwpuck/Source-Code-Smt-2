func rendezvous(x, y, z int){
    /*I.S. Terdefinisi bilangan bulat x,y,dan z
    F.S. Menampilkan  jumlah kedua mata mata bertemu
    */
    var i, jumlah int
     for i = 1; i <= 365; i++{
         if i % x == 0 && i % y == 0 && i % z != 0 {
             jumlah += 1
         }
        
     }
       fmt.Print(jumlah)
}