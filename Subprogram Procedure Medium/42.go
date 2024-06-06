func hitungTigaGenap(d1, d2, d3 int, total *int){
    /*I.S. Terdefinisi bilangan bulat d1,d2,d3 dan total
    F.S. Tersimpan banyaknya kelompok dadu yang ketiganya bilangan genap pada variabel total
    */
    if d1 % 2 == 0 && d2 % 2 == 0 && d3 % 2 == 0{
        *total += 1
    }
}
