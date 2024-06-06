//buatlah procedure saja

func vokalKonsonan(x byte) {
    /*I.S. terdefinisi karakter x
    F.S. menampilkan "vokal" jika x termasuk alfabet vokal dan "konsonan" jika bukan
    */ 
    if x == 'a' || x == 'i' || x == 'u' || x == 'e' || x == 'o' ||  x == 'A' || x == 'I' || x == 'U' || x == 'E' || x == 'O'{
        fmt.Print("vokal")
    }else{
        fmt.Print("konsonan")
    }
    
}