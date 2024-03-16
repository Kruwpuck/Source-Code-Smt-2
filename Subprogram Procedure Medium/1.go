//buatlah procedure saja

func belanja(a, b, c, d, e int) {
	if a+b+c == c+d+e {
		if (b+c+d)%(e+a) == 0 {
			fmt.Println("cashback")
			fmt.Println("diskon")
		} else {
			fmt.Println("cashback")
		}
	} else if (b+c+d)%(e+a) == 0 {
		fmt.Println("diskon")
	}
}
