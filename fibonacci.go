package algorithms

func multiply(a, b [][]int64, m int64) [][]int64 {
	c := [][]int64{
		{0, 0},
		{0, 0},
	}
	x := a[0][0]*b[0][0] + a[0][1]*b[1][0]
	y := a[0][0]*b[0][1] + a[0][1]*b[1][1]
	z := a[1][0]*b[0][0] + a[1][1]*b[1][0]
	w := a[1][0]*b[0][1] + a[1][1]*b[1][1]

	c[0][0] = x % m
	c[0][1] = y % m
	c[1][0] = z % m
	c[1][1] = w % m

	return c

}

func power(n, m int64) int64 {
	F := [][]int64{
		{0, 1},
		{1, 1},
	}
	M := [][]int64{
		{1, 0},
		{0, 1},
	}
	for n > 0 {
		if n&1 == 1 {
			M = multiply(F, M, m)
		}
		F = multiply(F, F, m)
		n >>= 1
	}
	return M[0][0]

}

// F computes n-th fibonacci number modulo m.
func Fib(n, m int64) int64 {
	return power(n, m)
}
