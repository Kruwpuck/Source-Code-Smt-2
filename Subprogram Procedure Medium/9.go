// Buatlah prosedur saja

import "math"

func radar(jk string, x, y float64) {
/* I.S. Terdefinisi tiga variabel, yaitu jenis kapal berupa string, posisi x relatif terhadap radar berupa bilangan real dan posisi y relatif terhadap radar berupa bilangan real
   F.S. Menampilkan string "tembak" atau "tidak tembak" bergantung dari kondisi yang benar*/
   
    var jarak float64
    
    jarak = math.Sqrt(x * x + y * y)
    
   if jk == "tempur" && jarak <= 5 {
		fmt.Println("tembak")
	} else {
		fmt.Println("tidak tembak")
	}
}