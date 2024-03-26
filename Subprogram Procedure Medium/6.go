func vikingVS(viking, saxon int) {
	if viking*4 > saxon {
		fmt.Print("viking menang")
	} else if viking*4 == saxon {
		fmt.Print("imbang")
	} else {
		fmt.Print("saxon menang")
	}
}