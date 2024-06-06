func sum(kar byte, jumlah *int) {
	/*  I.S. Terdefinisi kar byte dan jumlah int
		F.S. jumlah huruf selain  'a' dan 'i' terdefinisi di jumlah int */ 
		
		for kar != '.'{
			if kar != 'a' && kar != 'i'{
				*jumlah += 1
			}
			fmt.Scanf("%c", &kar)
		}
	}