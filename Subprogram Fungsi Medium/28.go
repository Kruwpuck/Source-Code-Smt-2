func tagihanListrik(watt int) int {
	// menghitung tagihan listrik pelanggan
	var b int
	if watt < 200 {
	b = watt * 12
	} else if watt >= 200 && watt < 400 {
	b = watt * 15
	} else if watt >= 400 && watt < 600 {
	b = watt * 18
	} else if watt >= 600 {
	b = watt * 20
	}
	if b >= 400 {
	b += b * 20 / 100
	}
	if b < 100 {
	b = 100
	}
	return b
   }
   