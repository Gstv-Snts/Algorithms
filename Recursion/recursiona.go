package recursion

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
