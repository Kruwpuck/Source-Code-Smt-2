func cacah(bilangan int) {
	/* I.S. terdefinisi sebuah bilangan bulat "bilangan"
	   F.S. program mengeluarkan nilai yang menyatakan nilai setiap digit dari X dengan setiap nilai dipisahkan oleh spasi */
	var digits []int
	
	// Ekstraksi digit dari bilangan
	for bilangan > 0 {
		digit := bilangan % 10
		digits = append(digits, digit)
		bilangan /= 10
	}
	
	// Cetak digit dalam urutan terbalik dengan spasi di antara mereka
	for i := 0 ; i <= len(digits) - 1; i++ {
		fmt.Print(digits[i], " ")
	}
	fmt.Println()
}