func teksGambarWhatsapp(n int){
	/*  I.S Terdefinisi jumlah hari
	   F.S Menampilkan jumlah postingan masing-masing teks dan gambar*/
	  var i, teks, gam, jumteks, jumgam int
	  
	  for i = 1; i <= n; i++{
		  fmt.Scan(&teks, &gam)
		  jumteks += teks
		  jumgam += gam
	  }
	  fmt.Print(jumteks, jumgam)
  }
  