func sertifikat(P, V, GA, FS, CI, TC int){
    /*
    IS : Terdefinisi 6 bilangan bulat (P, V, GA, FS, CI, TC) dengan nilai dari 1 hingga 4
    FS : Menampilkan string  "mendapat sertifikat" jika rata-rata nilai minimal 3.33, atau "tidak mendapat sertifikat" jika tidak.
    */
    
    fmt.Scan(&P, &V, &GA, &FS, &CI, &TC)
    rata := float64(P + V + GA +FS + CI +TC)/6.0
    
    if rata >= 3.33 {
        fmt.Print("mendapat sertifikat")
    } else {
        fmt.Print("tidak mendapat sertifikat")
    }
}
