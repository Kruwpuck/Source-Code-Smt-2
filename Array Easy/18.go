package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	fmt.Println((t * t * t) - 12*(t*t) + (36 * t) - 30)
}
