//buatlah procedure saja

func totalPelajar(a string, b int, totSD, totSMP, totSMA *int) {
	if a == "sd" {
		*totSD += b
	} else if a == "smp" {
		*totSMP += b
	} else if a == "sma" {
		*totSMA += b
	}
}