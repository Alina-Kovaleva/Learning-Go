package utils

import (
	"fmt"
	"strconv"
)

func FizzBuzz(a int) string {
	// for i := 1; i <= 100; i++ {
	if a%3 == 0 && a%5 == 0 {
		fmt.Println("FizzBuzz")
		return "FizzBuzz"
	} else if a%3 == 0 {
		fmt.Println("Fizz")
		return "Fizz"
	} else if a%5 == 0 {
		// fmt.Println("Buzz")
		return "Buzz"

	}
	// else {
	// 	fmt.Println(a)
	// }
	return strconv.Itoa(a)
	// }
}
