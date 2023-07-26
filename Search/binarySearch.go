package search

func BinarySearch(haystack []int, needle int) int {
	lo := 0
	hi := len(haystack) - 1
	var mi int
	for lo < hi {
		mi = (hi + lo) / 2
		v := haystack[int(mi)]
		if v == needle {
			return int(mi)
		} else if v < needle {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}
	return int(mi)
}
