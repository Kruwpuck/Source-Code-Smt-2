package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1.0
	b := -3.0
	c := 18.0
	akar1 := -b + math.Sqrt((b*b)-4.0*a*c)/(2.0*a)
	akar2 := -b - math.Sqrt((b*b)-4.0*a*c)/(2.0*a)
	fmt.Println(akar1, akar2)
}
