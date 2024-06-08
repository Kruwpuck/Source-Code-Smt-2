// implementasikan fungsinya saja
func kirimPesanRahasia( pesan string ) string {
	// mengembalikan pesan rahasia pada karakter ke-1, ke-5, ke-9, dan ke-13 dari
	var pesan_rahasia string
	if len(pesan) >= 14{
	pesan_rahasia = string(pesan[0]) + string(pesan[4]) + string(pesan[8]) +
	return pesan_rahasia
	}else {
	return " "
   }
   } 
   