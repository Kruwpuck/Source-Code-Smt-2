func gerbangTol(jumlahKendaraaanLewat int, jenisKendaraan string) {
    /*IS : Terdefinisi jumlah kendaraan yang lewat, dan jenis kendaraan yang lewat
      FS : Menampilkan total pendapatan dan jumlah masing-masing jenis kendaraan yang lewat.*/
    
    var totalPendapatan, jumA, jumB, jumC int

	jumA, jumB, jumC, totalPendapatan = 0, 0, 0, 0

	fmt.Scan(&jumlahKendaraaanLewat)

	for i := 1; i <= jumlahKendaraaanLewat; i++ {
		fmt.Scan(&jenisKendaraan) // gunakan scan untuk memperoleh jumlah kendaraan berdasarkan jenis kendaraan 
		if jenisKendaraan == "A" {
			totalPendapatan += 10000
			jumA += 1 // lakukan perhitungan 
		} else if jenisKendaraan == "B" {
			totalPendapatan += 23000
			jumB += 1// Lakukan perhitungan 
		} else if jenisKendaraan == "C" {
		totalPendapatan += 45000 
		jumC += 1// lakukan perhitungan 
		}
	}

	fmt.Println( totalPendapatan, jumA, jumB, jumC) // tampilkan yang nilai yang perlu dimunculkan 

}
