//buat prosedur saja
func dermawan(n int) {
    /*I.S terdefinisi bilangan bulat n yang menyatakan banyaknya minggu dalam bulan tersebut. Diikuti oleh n baris bilangan bulat yang menyatakan banyaknya uang per minggu yang dikumpulkan.
    F.S karakter total uang yang dikumpulkan dalam sebulan */
	 var uang, jumlah, i int
	 
	 for i = 1; i <= n; i++{
	     fmt.Scan(&uang)
	     jumlah += uang
	 }
	 fmt.Print(jumlah)
}