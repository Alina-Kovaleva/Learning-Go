package utils

import "fmt"

func Variables() {
	a := 5
	b := 5
	if a < b {
		fmt.Printf("%.2f", float64(a)/float64(b))
	} else if b < a {
		fmt.Println(a << b)
	} else {
		if a%2 == 0 {
			fmt.Println("EVEN")
		} else {
			fmt.Println("ODD")
		}
	}

}
