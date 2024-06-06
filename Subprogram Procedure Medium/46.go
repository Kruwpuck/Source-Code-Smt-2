func cekGanjil(b1, b2, b3, b4, b5 int) {
    /*
    IS : terdefinisi 5 bilangan bulat positif
    FS : menampilkan string "ganjil semua", "ganjil sebagian", atau "tidak ada bilangan ganjil" bergantung kondisinya.
    */
    fmt.Scan(&b1, &b2, &b3, &b4, &b5)
    if b1 % 2 != 0 && b2 % 2 != 0 && b3 % 2 != 0 && b4 % 2 != 0 && b5 % 2 != 0{
        fmt.Print("ganjil semua")
    } else if b1 % 2 != 0 || b2 % 2 != 0 ||b3 % 2 != 0 || b4 % 2 != 0 || b5 % 2 != 0 {
         fmt.Print("ganjil sebagian")
    } else {
       fmt.Print("tidak ada bilangan ganjil")
    }
}
