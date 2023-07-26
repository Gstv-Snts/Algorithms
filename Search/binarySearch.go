package search

func BinarySearch(haystack []int, needle int) int {
	lo := 0
	hi := len(haystack) - 1
	mi := (hi + lo) / 2
	for lo <= hi {
		mi = (hi + lo) / 2
		v := haystack[mi]
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
