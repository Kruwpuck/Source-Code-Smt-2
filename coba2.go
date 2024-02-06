package main

import "fmt"

func main() {
	var maks, maks2, perantara, input int
	maks = 0
	maks2 = 0
	stop := false

	for !stop {
		fmt.Scan(&input)
		if input > maks {
			perantara = maks
			maks2 = perantara
			maks = input
		} else if input > maks2 && input != maks {
			maks2 = input
		}
		stop = input == 0
	}

	fmt.Println(maks2)
}
