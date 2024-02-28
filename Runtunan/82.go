package main

import "fmt"

func main() {
	var drach, mine, oboli, attic int
	fmt.Scan(&drach)
	oboli = drach * 6
	attic = oboli / 36000
	mine = (oboli % 36000) / 600
	oboli = oboli % 36000 % 600
	fmt.Println(attic, mine, oboli)
}
