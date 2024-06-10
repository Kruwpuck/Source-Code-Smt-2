package main

import "fmt"

func main() {
	var friendly, notFriendly int

	for {
		var a, b, c int
		fmt.Scanln(&a, &b, &c)

		if a == 0 && b == 0 && c == 0 {
			break
		}

		if (a != b && b != c && a != c) && (a+b == c || b+c == a || a+c == b) {
			friendly++
		} else {
			notFriendly++
		}
	}

	fmt.Print(friendly, " ")
	fmt.Print(notFriendly, " ")
}