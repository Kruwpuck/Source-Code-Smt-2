func hitungGanjilGenap(banyakGenap, banyakGanjil *int, b1, b2, b3, b4 int) {
    /*I.S. terdefinisi banyakGanjil dan banyakGenap diinisiasi bernilai 0 dan 
    bilangan bulat b1 b2 b3 b4
    F.S. prosedur mengembalikan banyaknya bilangan ganjil dan genap ke dalam 
    banyakGanjil dan banyakGenap sesuai dengan syarat bilangan ganjil/genap
    */
    if b1 % 2 == 0 {
        *banyakGenap += 1
    } else {
        *banyakGanjil += 1
    }
    if b2 % 2 == 0 {
        *banyakGenap += 1
    } else {
        *banyakGanjil += 1
    }
    if b3 % 2 == 0 {
        *banyakGenap += 1
    } else {
        *banyakGanjil += 1
    }
    if b4 % 2 == 0 {
        *banyakGenap += 1
    } else {
        *banyakGanjil += 1
    }
}
