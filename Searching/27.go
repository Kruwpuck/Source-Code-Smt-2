
const N int = 1000

func indeks(T [N]int, total, bilX int) int {
	/*
	algoritma untuk mencari indeks dari bilX, jika tidak ditemukan akan mengembalikan -1
	mengembalikan indeks dari bilX
	*/
	var i, idx int
	idx = -1
	i = 0
	for i < total && idx == -1 {
		if T[i] == bilX {
			idx = i
		}
		i ++ 
	}
	return idx
}