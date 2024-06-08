func jumlahDeretGenap(N int) int {
	if N <= 0 {
		return 0
	}
	if N%2 == 0 {
		return N + jumlahDeretGenap(N-2)
	}
	return jumlahDeretGenap(N - 1)
}