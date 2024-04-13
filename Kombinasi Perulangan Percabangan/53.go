package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a == 1 {
		fmt.Print("minggu")
	} else if a == 2 {
		fmt.Print("senin")
	} else if a == 3 {
		fmt.Print("selasa")
	} else if a == 4 {
		fmt.Print("rabu")
	} else if a == 5 {
		fmt.Print("kamis")
	} else if a == 6 {
		fmt.Print("jumat")
	} else if a == 7 {
		fmt.Print("sabtu")
	}

}
