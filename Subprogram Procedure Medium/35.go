// Buatlah prosedur saja

func pesanRahasia(a,b,c,d,e,f string) {
    /*I.S terdefinisi enam baris string pesan rahasia
     F.S  dicetak string "serang" atau "tidak serang" bergantung nilai kondisinya*/
    var sandi string
   sandi = string(a[0]) + string(b[0]) + string(c[0]) + string(d[0]) + string(e[0]) + string(f[0])

    
    if sandi == "serang" {
        fmt.Println("serang")
    } else {
        fmt.Println("tidak serang")
    }
}