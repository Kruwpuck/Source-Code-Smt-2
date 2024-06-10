
const N int = 1000

func searchMax(T [N]float64) float64 {
	/*
	    algoritma untuk mencari nilai paling besar dari empat orang mahasiswa
	    note : mengembalikan nilai mahasiswa yang paling besar
	    hint : definisikan kamus variabel lokal
	*/ 
	
	var idx int
	idx = 0
	for i := 1; i < 4; i++ {
		if T[idx] < T[i]  {
			idx = i
		}
	}
	return T[idx]
}