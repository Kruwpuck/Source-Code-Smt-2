// Buatlah prosedur saja

func luasKelilingPersegi(n int) {
    /*I.S Terdefinisi bilangan bulat n yang menyatakan banyaknya persegi dan sebanyak n sisi.
       F.S Menampilkan luas dan keliling persegi.*/
    var l, k, i, s int
  
    for i = 1; i <= n; i++ {
        fmt.Scan(&s)
        l = s * s
        k = 4 * s
		fmt.Println(l, k)
		}
    }