func KTP(umur int, kartu bool) string {
	if umur >= 17 && kartu {
		return "bisa membuat KTP"
	} else {
		return "belum bisa membuat KTP"
	}
}