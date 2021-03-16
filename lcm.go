package algorithms

func LCM(a, b int64) int64 {
	return a * b / GCD(a, b)
}
