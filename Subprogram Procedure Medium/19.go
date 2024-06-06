func cetak(b, k int, x string) {
	/*
	I.S. terdefinisi sebuah bilangan bulat, n, yang menyatakan banyaknya pencetakan dan sebuah string, x, yang akan dicetak.
	F.S. tercetak string x sebanyak n kali. 
	*/
		var i, j int
		for i = 0; i < b; i++{
			for j = 0; j < k; j++{
				fmt.Print(x, " ")
			}
			fmt.Println()
		}
	
	}