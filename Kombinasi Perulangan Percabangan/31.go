package main

import "fmt"

func main() {
	var rounds int
	fmt.Scan(&rounds)

	var totalA, totalB, totalC int
	for i := 0; i < rounds; i++ {
		var a, b, c, evenA, evenB, evenC int
		fmt.Scan(&a, &b, &c)
		evenA = countEven(a)
		evenB = countEven(b)
		evenC = countEven(c)
		totalA += evenA
		totalB += evenB
		totalC += evenC
	}

	winner := 'A'
	maxTotal := totalA
	if totalB > maxTotal {
		winner = 'B'
		maxTotal = totalB
	}
	if totalC > maxTotal {
		winner = 'C'
		maxTotal = totalC
	}

	fmt.Printf("%c %d\n", winner, maxTotal)
}

func countEven(num int) int {
	count := 0
	for num > 0 {
		digit := num % 10
		if digit%2 == 0 {
			count += digit
		}
		num /= 10
	}
	return count
}