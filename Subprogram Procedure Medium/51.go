func course(ma,fi,ki int) {
    /*I.S. terdefinisi tiga buah bilangan bulat yang menyatakan nilai matematika,
    fisika, dan kimia.
    F.S. print boolean true jika kandidat dapat mendaftar atau false jika
    kandidat tidak dapat mendaftar sesuai dengan syarat diatas.
    */
    var total int
    var hasil bool
    
    total = ma + fi + ki
    hasil = total >= 180 && ma >= 65 && fi >= 55 && ki >= 50
    fmt.Print(hasil)
    
}