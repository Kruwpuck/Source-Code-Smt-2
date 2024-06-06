// Buatlah prosedur saja

func cekganjilGenap(x string, n, m int){
    /*I.S. Terdefinisi variabel ganjilGenap yang berisikan string "ganjil" atau "genap", kemudian bilangan bulat n dan m
    F.S.  Menampilkan bilangan ganjil atau genap pada rentang n sampai m.
    */
    var i int
    
    if x == "genap" {
        for i = n; i <= m; i++{
            if i % 2 == 0 {
                fmt.Print(i, " ")
            }
        }
    } else if x == "ganjil" {
        for i = n; i <= m; i++{
            if i % 2 != 0{
                fmt.Print(i, " ")
            }
        }
    }
}