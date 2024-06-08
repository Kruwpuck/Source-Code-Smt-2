func latihan(i,yin,yang int, A,B *float64) {
	/*I.S. terdefinisi hari ke-i, jumlah yin dan yang yang berhasil dikumpulkan
	Saitama pada hari ke-i, serta kekuatanA dan kekuatanB yang telah diperoleh sebelumnya oleh saitama..
	F.S. kemampuan A dan B dari Saitama diupdate sesuai dengan ketentuan diatas.
	*/
	if yin > yang {
        *A += 0.25
    } else if yang > yin {
        *B += 0.15
    }

    if i%3 == 0 {
        *A = *A - (*A * 0.05)
    }
}