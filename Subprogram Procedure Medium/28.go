func pemenang() {
	var p1, p2, p3 byte
	var n1, n2, n3 int

	for {
		_, err := fmt.Scanf("%c%c%c", &p1, &p2, &p3)
		if err != nil {
			break
		}

		if p1 == '.' && p2 == '.' && p3 == '.' {
			break
		} else {
			if p1 == p2 && p2 != p3 {
				n3++
			} else if p1 == p3 && p2 != p1 {
				n2++
			} else if p2 == p3 && p1 != p2 {
				n1++
			}
		}
	}

	fmt.Println(n1, n2, n3)
}