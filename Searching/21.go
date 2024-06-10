
const N int = 1000

func duplikat(T [N]int, jumlah int) bool {
    /*  IS :  terdefinisi jumlah dan nilai integer input sebanyak jumlah 
        FS :  mengembalikan nilai boolean true / false */
    
    var hasil bool
    hasil = false
	for i := 0 ; i < jumlah ; i++ {
		for j := 0 ; j < i ; j++ {
			if T[j] == T[i]  {
				hasil = true
			}
		}
	}
	return hasil
}