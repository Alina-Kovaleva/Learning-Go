package utils

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if b > a {
		return a
	}
	return b
}
