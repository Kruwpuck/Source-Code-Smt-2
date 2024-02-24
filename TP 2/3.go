package main

import "fmt"

func main() {
	var kuadran1, kuadran2, kuadran3, kuadran4 bool
	var x, y float32
	fmt.Scan(&x, &y)
	kuadran1 = x > 0 && y > 0
	kuadran2 = x < 0 && y > 0
	kuadran3 = x < 0 && y < 0
	kuadran4 = x > 0 && y < 0
	if kuadran1 {
		fmt.Print("Kuadran I")
	} else if kuadran2 {
		fmt.Print("Kuadran II")
	} else if kuadran3 {
		fmt.Print("Kuadran III")
	} else if kuadran4 {
		fmt.Print("Kuadran IV")
	}
}
