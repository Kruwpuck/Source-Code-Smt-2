func printStars(N, current, stars int, result string) string {
	if current > N {
		return result
	}
	if current%2 == 0 {
		for i := 0; i < current; i++ {
			result += "*"
		}
		result += "\n"
	}
	return printStars(N, current+1, 0, result)
}