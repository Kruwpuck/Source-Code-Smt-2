package main

import "fmt"

// Function to check if a number is perfect or not
func isPerfectNumber(num int) bool {
	sum := 0
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if isPerfectNumber(num) {
		fmt.Println(num, "is a perfect number.")
	} else {
		fmt.Println(num, "is not a perfect number.")
	}
}
