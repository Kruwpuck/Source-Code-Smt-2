func jumlahA(x string, index int) int {
	if index >= len(x) {
	return 0
	}
	kar := string(x[index])
	if kar == "a"{
	return jumlahA(x, index + 1) + 1
	}else{
	return jumlahA(x, index + 1)
	}
   
   
   }
   