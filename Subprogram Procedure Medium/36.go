func sentimen(n int) {
    /*I.S terdefinisi n yang menyatakan banyaknya postingan, dan 3 bilangan yang menyatakan jumlah sentimen poditif, negatif dan netral.
      F.S Karakter '+', '=', atau '-' yang menyatakan jumlah sentimen positif, netral, atau negatif yang paling banyak. Diikuti oleh banyaknya nilai sentimennya.*/
    var a, b, c, juma, jumb, jumc, i int
    
    for i = 1; i <= n; i++{
        fmt.Scan(&a, &b, &c)
        juma += a
        jumb += b
        jumc += c
    }
    if juma > jumb && juma > jumc{
        fmt.Print("+ ", juma)
    }else if jumb > juma && jumb > jumc{
        fmt.Print("= ", jumb)
    }else if jumc > juma && jumc > jumb{
        fmt.Print("- ", jumc)
    }
    
}
