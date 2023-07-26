package search

import (
	"fmt"
	"math"
	"time"
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
		time.Sleep(1 * time.Second)
		fmt.Printf("previous: %v, current: %v, middle:%v\n", pr, cr, mi)
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
		time.Sleep(1 * time.Second)
		fmt.Printf("Sqrt: %v\n", i)
		i += jumpAmount
		if height[i] {
			break
		}
	}
	fmt.Printf("Positive sqrt: %v\n", i)
	i -= jumpAmount
	fmt.Printf("Negative sqrt: %v\n", i)
	for j := i; j < len(height); j++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Current index is: %v\n", j)
		if height[j] {
			return j
		}
	}
	return -1
}
