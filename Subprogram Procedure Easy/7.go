func goNoGo(hujanTurun string, bawaPayung string) {
	/*
	I.S. terdefinisi dua buah string hujanTurun dan bawaPayung ("ya" atau "tidak").
	F.S. menampilkan string "berangkat" bila hujan tidak turun atau membawa payung, 
		 atau menampilkan "diam di rumah" bila sebaliknya.
	*/
	
		if hujanTurun == "ya" && bawaPayung == "ya" {
	 fmt.Println("berangkat")
	 } else if hujanTurun == "ya" && bawaPayung == "tidak" {
	 fmt.Println("diam di rumah")
	 } else if hujanTurun == "tidak" && bawaPayung == "tidak" {
	 fmt.Println("berangkat")
	 } else if hujanTurun == "tidak" && bawaPayung == "ya" {
	 fmt.Println("berangkat")
	 }
	
	}