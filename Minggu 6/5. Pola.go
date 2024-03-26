package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibo(n))

}
func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}
