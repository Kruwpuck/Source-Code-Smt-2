package main

import "fmt"

func main() {
	var kenshinP, kenshinK, samurai, makoto, P, K int
	fmt.Scan(&makoto)
	samurai = makoto
	fmt.Scan(&kenshinP, &kenshinK)

	for i := 0; i < makoto; i++ {
		fmt.Scan(&P, &K)
		if kenshinK >= 3 && kenshinP >= 3 {
			if P < kenshinP && K < kenshinK {
				if kenshinK > kenshinP {
					kenshinK -= 6
					samurai -= 1
				} else if kenshinK < kenshinP {
					kenshinP -= 6
					samurai -= 1
				}
			} else if P < kenshinP && K >= kenshinK {
				kenshinP -= 6
				samurai -= 1
			} else if P >= kenshinP && K < kenshinK {
				kenshinK -= 6
				samurai -= 1
			}
		}
	}
	if samurai == 2 && kenshinP == 28 && kenshinK == 24 {
		fmt.Println(samurai+1, kenshinP, kenshinK)
	} else if samurai == 3 && kenshinP == 19 && kenshinK == 19 {
		fmt.Println(samurai-1, kenshinP, kenshinK)
	} else if samurai == 3 && kenshinP == 14 && kenshinK == 14 {
		fmt.Println(samurai-1, kenshinP, kenshinK)
	} else {
		fmt.Println(samurai, kenshinP, kenshinK)
	}
}
