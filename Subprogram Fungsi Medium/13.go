// implementasikan fungsinya saja
func tetangga (k1, k2 string) bool {
	// mengembalikan total biaya parkir kendaraan dalam durasi tertentu
	if k1 == "Bogor" && (k2 == "Sukabumi" || k2 == "Cianjur"){
	return true
	}else if k1 == "Sukabumi" && (k2 == "Bogor" || k2 == "Cianjur"){
	return true
	}else if k1 == "Cianjur" && (k2 == "Sukabumi" || k2 == "Bogor"){
	return true
	}else if k1 == "Cianjur" && k2 == "Bandung" {
	return true
	}else if k1 == "Bandung" && k2 == "Cianjur"{
	return true
	}else if k1 == "Bandung" && k2 == "Tasikmalaya"{
	return true
	}else if k1 == "Tasikmalaya" && k2 == "Bandung"{
	return true
	}else {
	return false
	}
}   