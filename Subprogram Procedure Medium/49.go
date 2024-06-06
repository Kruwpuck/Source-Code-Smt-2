func kuadran(x,y int) {
    /*I.S. terdefinisi dua buah bilangan bulat positif atau negatif yang 
    menyatakan nilai x dan y dari sebuah titik. 
    F.S. keluarkan bilangan bulat 1, 2, 3, atau 4 jika masing-masing berada pada
    kuadran 1, 2, 3, atau 4.


    */
    if x > 0 && y > 0{
        fmt.Print("1")
    }else if x < 0 && y > 0{
        fmt.Print("2")
    }else if x < 0 && y < 0{
        fmt.Print("3")
    }else if x > 0 && y < 0{
        fmt.Print("4")
    }
}