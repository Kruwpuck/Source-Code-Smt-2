//isi prosedur saja
func jumlahNumerik(kar byte) {
    /*I.S.terdefinisi kar berisi serangkaian character yang akan dihitung jumlah 
    numeriknya
    F.S. berupa bilangan bulat yang menyatakan jumlah karakter bertipe numerik
    */
    var jumlah int
    
    for kar != '.'{
        if kar >= '0' && kar <= '9'{
            jumlah += int(kar - '0') 
        }
        fmt.Scanf("%c", &kar)
    }
    fmt.Print(jumlah)
    
}