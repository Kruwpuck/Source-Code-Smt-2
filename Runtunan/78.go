package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)
	uppercaseString := strings.ToLower(input)
	fmt.Printf(uppercaseString)
}
