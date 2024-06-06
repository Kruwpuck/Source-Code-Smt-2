func cekCLO (clo1,clo2,clo3 *float64){
    /*I.S. terdiri 3 bilangan yang masing-masing menyatakan nilai CLO-1, CLO-2 dan CLO-3 
    F.S. print status kelulusan dari mahasiswa berdasarkan syarat kelulusan diatas.
    setiap baris dipisahkan oleh newline dan setelah status kelulusan diprint, 
    nilai CLO-1, CLO-2 dan CLO-3 diupdate dengan inputan pengguna.
    */
   
    // Mengecek apakah mahasiswa lulus atau tidak berdasarkan syarat kelulusan
    if *clo1 > 50.0 && *clo2 > 50.0 && *clo3 > 50.0 {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
    fmt.Scan(clo1, clo2, clo3)
   
}