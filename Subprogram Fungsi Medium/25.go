// Buatlah fungsi saja
func uangHotel(hotel string, hari int) int{
	var saku int
	if hotel == "A"{
	saku = (hari-1) * 500000
	}else if hotel == "B"{
	saku = (hari-1) * 400000
	}else if hotel == "C"{
	saku = (hari-1) * 350000
	}
	return saku
   }
   