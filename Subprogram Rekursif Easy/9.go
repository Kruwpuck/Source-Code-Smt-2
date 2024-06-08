func kakiHewan(N int) int {
	if N == 0 {
	return 0
	} else if N % 2 == 1{
	return kakiHewan(N-1) + 2
	}else{
	return kakiHewan(N-1) + 3
	}
   
   }