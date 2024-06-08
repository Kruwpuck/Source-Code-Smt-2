func terkecil(a, b, c int) {
	var min int
	if a < b && a < c {
		min = a
	} else if b < a && b < c {
		min = b
	} else {
		min = c
	}
	fmt.Println(min)
}
