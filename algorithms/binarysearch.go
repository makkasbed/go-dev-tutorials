package algorithms

func BinarySearch(n []int, start int, end int, s int) int {

	var mid int = (start + end) / 2

	if n[mid] < s {
		//search from the mid +1
		return BinarySearch(n, mid+1, end, s)
	} else {
		//search from mid+1
		return BinarySearch(n, start, mid-1, s)
	}

}
