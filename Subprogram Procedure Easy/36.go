package main

import "fmt"

func main() {
	fmt.Println((factorial(5)))
}
func rekur(n int) int {
	if n <= 1 {
		return 1
	}
	return n + rekur(n-1)
}