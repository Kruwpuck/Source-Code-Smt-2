func buburAyam(suwirAyam, cakue, atiAmpela, telur bool){
    /* 
    IS : Terdefinisi 4 boolean "true" atau "false"
    FS : Menampilkan total uang yang harus dibayar yang bertipe integer
    */
    
    //input tambahan
    hargaSuwirAyam := 3000
    hargaCakue := 1500
    hargaAtiAmpela := 4500
    hargaTelur := 4000
    
    total := 6000
    fmt.Scan(&suwirAyam, &cakue, &atiAmpela, &telur )
    if suwirAyam {
        total = total + hargaSuwirAyam
    }
    
    if cakue {
        total +=  hargaCakue 
    }
    
    if atiAmpela {
        total+= hargaAtiAmpela
    }
    
    if telur {
        total = total + hargaTelur
    }
    
    fmt.Print(total)
}