package search

import (
	"math"
)

func CrystalBallBreakBinary(height []bool) int {
	lo := 0
	hi := len(height)
	mi := (lo + hi) / 2
	pr := false
	cr := false
	for pr == cr {
		mi = (lo + hi) / 2
		cr = height[mi]
		pr = height[mi-1]
		if !pr && !cr {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}
	return mi
}

func CrystalBallBreakSqrt(height []bool) int {
	jumpAmount := int(math.Sqrt(float64(len(height))))
	i := jumpAmount
	for jumpAmount <= len(height) {
		i += jumpAmount
		if height[i] {
			break
		}
	}
	i -= jumpAmount
	for j := i; j < len(height); j++ {
		if height[j] {
			return j
		}
	}
	return -1
}
