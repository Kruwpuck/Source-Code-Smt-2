package main

import "fmt"

func main() {
	var kenshinP, kenshinK, samurai, makoto, P, K int
	fmt.Scan(&makoto)
	fmt.Scan(&kenshinP, &kenshinK)

	for i := 0; i < makoto; i++ {
		fmt.Scan(&P, &K)
		if kenshinK >= 3 && kenshinP >= 3 {
			if P < kenshinP && K < kenshinK {
				if kenshinK > kenshinP {
					kenshinK -= 6
					samurai += 1
				} else if kenshinK < kenshinP {
					kenshinP -= 6
					samurai += 1
				}
			} else if P < kenshinP && K >= kenshinK {
				kenshinP -= 6
				samurai += 1
			} else if P >= kenshinP && K < kenshinK {
				kenshinK -= 6
				samurai += 1
			}
		}
	}

	fmt.Println(samurai, kenshinP, kenshinK)

}
