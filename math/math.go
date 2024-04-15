package math

// 最大公约数 Greatest Common Divisor：指能同时整除a和b的最大整数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

// 两个或两个以上的自然数中，如果它们有相同的倍数，这些倍数就是它们的公倍数，其中除0以外最小的一个公倍数，叫做这几个数的最小公倍数
func lcm(a, b int) int {
	return a * b / gcd(b, a%b)
}
