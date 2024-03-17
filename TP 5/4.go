package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(fibonacci_rekursif(N))
}
func fibonacci_rekursif(N int) int {
	if N <= 3 {
		return 1
	} else {
		return fibonacci_rekursif(N-1) + fibonacci_rekursif(N-2) + fibonacci_rekursif(N-3)
	}
}
