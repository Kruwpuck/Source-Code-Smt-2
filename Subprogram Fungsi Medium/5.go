// implementasikan fungsinya saja
func review( a, b, c, d int ) string {
	// mengembalikan string "bagus", "kurang", atau "sedang" tergantung rata-ratanya
	 var rata float64
	 rata = float64(a + b + c + d)/ 4
	 if rata >=3.75 {
	 return "bagus"
	 }else if rata <= 1.50 {
	 return "kurang"
	 }else {
	 return "sedang"
	 }
	}