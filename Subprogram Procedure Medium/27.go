func akarReal(a,b,c float64){
	/*I.S Terdefinisi bilangan bulat a, b dan c
	  F.S Menampilkan string "memiliki akar real" atau "tidak memiliki akar real" bergantung nilai kondisinya.*/

	var d float64
	d = (b*b)-(4*a*c)
	if d >= 0{
		fmt.Println("memiliki akar real")
		
	}else{
		fmt.Println("tidak memiliki akar real")
	}
}