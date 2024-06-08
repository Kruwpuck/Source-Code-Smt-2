// Buatlah fungsi saja
func uangSaku(kota string, hari int) int{
	var saku int
	if kota == "A"{
	saku = hari * 500000
	}else if kota == "B"{
	saku = hari * 350000
	}else if kota == "C"{
   saku = hari * 250000
	}
	return saku
   }
   