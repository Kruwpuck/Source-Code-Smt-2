//buat prosedur saja
func inputWaktu(jumlah *int){
    /*I.S terdefinisi jumlah bernilai 0
     F.S jumlah berisi total waktu putar dalam satuan menit*/
    var j, m int
    
    fmt.Scan(&j, &m)
    for j != 0 && m != 0{
        *jumlah += (j * 60) + m
		fmt.Scan(&j, &m)
    }
	
}

func printWaktuPutar(jumlah int){
    /*I.S terdefinisi jumlah berisi total waktu putar dalam satuan menit
     F.S program menampilkan total waktu putar dalam satuan jam dan menit*/
     var jam , menit int
     jam = jumlah / 60
     menit = jumlah % 60

	 fmt.Print(jam, menit)
}