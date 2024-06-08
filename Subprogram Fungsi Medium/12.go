func rendezvous(x, y, z int)int{
	var i, jumlah int
	for i = 1; i <= 365; i++{
	if i%x == 0 && i%y == 0 && i%z == 0{
	jumlah = jumlah + 1
	}
	}
	return jumlah
   }
   