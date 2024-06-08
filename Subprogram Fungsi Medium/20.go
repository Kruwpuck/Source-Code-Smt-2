func tetangga(d1, d2 byte) bool {
	if d1 == 'A' && d2 == 'B'|| d1 == 'B' && d2 == 'A'{
	return true
	}else if d1 == 'B' && d2 == 'D' || d1 == 'D' && d2 == 'B'{
	return true
	}else if d1 == 'B' && d2 == 'C' || d1 == 'C' && d2 == 'B'{
	return true
	}else if d1 == 'C' && d2 == 'E' || d1 == 'E' && d2 == 'C'{
	return true
	}
	return false
   }
   