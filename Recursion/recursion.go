package recursion

import (
	"fmt"
)

func Multiply(v int) int {
	if v == 1 {
		return 1
	}

	return v * Multiply(v-1)
}

func Factorial(n int) int {
	if n < 1 {
		return 1
	}
	return n * Factorial(n-1)

	// 4 - 4 * third(6) return final(24)
	// 3 - 3 * second(2) return third(6)
	// 2 - 2 * first(1) return second(2)
	// 1 - 1 return first(1)
}

func Power(base int, exponent int) int {
	if exponent < 1 {
		return 1
	}
	return base * Power(base, exponent-1)
	// 4 - 2 * third(4) return final(8)
	// 3 - 2 * second(2) return third(4)
	// 2 - 2 * first(1) return second(2)
	// 1 - 1 return first(1)
}

func Quicksort(a []int) []int {
	fmt.Println(a)
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1

	pivot := len(a) / 2

	fmt.Println(a)
	a[pivot], a[right] = a[right], a[pivot]
	fmt.Println(a)

	for i := 0; i < right; i++ {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
			fmt.Println(a)
		}
	}
	a[left], a[right] = a[right], a[left]
	Quicksort(a[:left])
	Quicksort(a[left+1:])

	return a
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NonNegativesSum(n int) int {
	str := "#"
	for i := 0; i < n; i++ {
		str = str + "#"
	}
	fmt.Println(str)
	if n == 0 {
		return 0
	} else {
		return n + NonNegativesSum(n-1)
	}
}

func BottomRight(n int, m int) int {
	if n == 1 || m == 1 {
		return 1
	}
	return 0
}
