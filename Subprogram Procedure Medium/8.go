// Buatlah prosedur saja

func audisi(j1, j2, j3 string) {
	if j1 == "yes" && j2 == "yes" || j2 == "yes" && j3 == "yes" || j1 == "yes" && j3 == "yes" {
		fmt.Print("lolos")
	} else {
		fmt.Print("tidak lolos")
	}
}