// Buatlah fungsi saja

func validNIM(nim string) bool {
	return len(nim) == 8 && nim[0:3] == "212"
}