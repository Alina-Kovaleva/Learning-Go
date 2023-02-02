package utils

import "fmt"

func Fibonacci() {
	n := 10
	res := []int{}
	for i := 1; i <= n; i++ {
		res = append(res, fib(i))
	}
	fmt.Println(res)
	fmt.Println(fib1(10))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fib1(n int) []int {
	res := []int{1}
	if n <= 1 {
		return []int{}
	}
	a, b := 0, 1
	for i := 0; i < n-1; i++ {
		a, b = b, a+b
		res = append(res, b)
	}
	return res
}
