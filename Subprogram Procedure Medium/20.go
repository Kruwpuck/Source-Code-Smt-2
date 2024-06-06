// Buatlah prosedur saja

func pertemuan(x, y int, jumlah *int) {
    /*I.S Terdefinisi bilangan bulat positif x, y dan variabel integer jumlah bernilai nol 
    F.S variabel jumlah bertambah sesuai pertemuan x dan y*/
   var i int
   for i = 1; i <= 365; i++{
       if i % x == 0 && i % y == 0{
           *jumlah += 1
       }
   }
}