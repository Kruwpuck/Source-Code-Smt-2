func duel(turki, mongol int){
    /*I.S terdefinisi turki, mongol
      F.S string siapa pemenangnya, yaitu string "Turki menang", "Mongol menang", atau "Imbang" bergantung dari nilai kondisinya.*/ 

    turki = turki * 3
    if turki > mongol{
        fmt.Print("turki menang")
    }else if mongol > turki{
        fmt.Print("mongol menang")
    }else if turki == mongol{
        fmt.Print("imbang")
    }
}