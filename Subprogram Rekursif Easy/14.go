func hitungKapital(kata string, index int) int {
	if index == len(kata) {
	return 0
	}
	if 'A' <= kata[index] && kata[index] <= 'Z' {
	return 1 + hitungKapital(kata, index + 1) // masukkan fungsi rekursif ya
	}
	return hitungKapital(kata, index + 1) // masukkan fungsi rekursif yang sesua
   }