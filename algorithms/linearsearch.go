package algorithms

func LinearSearch(n []int, s int) (int, int) {
	for i := 0; i < len(n); i++ {
		if n[i] == s {
			return i, n[i]
		}
	}
	return -1, s
}
