func sumBuah(jumBuah int, jumApel *int, jumJeruk *int, jumMangga *int) {
    /*I.S. jumBuah terdefinisi, jumApel, jumJeruk, dan jumMangga bernilai 0
    Proses: Menerima nama nama buah sebanyak jumBuah pada piranti masukan
    (gunakan perulangan)
    F.S.  jumApel, jumJeruk, dan jumMangga masing-masing bernilai banyaknya apel,
    jeruk, dan mangga
    */

    *jumApel = 0
    *jumJeruk = 0
    *jumMangga = 0
    var i int
    var jenis string
   
    for i = 1; i <= jumBuah; i++{
        fmt.Scan(&jenis)
        nBuah(jenis, jumApel, jumJeruk, jumMangga)
    }
    

}
func nBuah(jenisBuah string, jumApel *int, jumJeruk *int, jumMangga *int) {
    /*IS:jumBuah terdefinisi, jumApel, jumJeruk, dan jumMangga bernilai 0
    FS: jumApel, jumJeruk, dan jumMangga masing-masing bernilai banyaknya apel, 
    jeruk, dan mangga 
    */
    if jenisBuah == "mangga"{
        *jumMangga++
    }else if jenisBuah == "jeruk"{
        *jumJeruk++
    }else if jenisBuah == "apel"{
        *jumApel++
    }
    
   
}