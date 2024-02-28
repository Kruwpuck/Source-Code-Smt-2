package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)
	uppercaseString := strings.ToUpper(input)
	fmt.Printf(uppercaseString)
}
