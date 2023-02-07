package utils

import (
	"fmt"
	"strconv"
)

func Variables(a int, b int) string {
	// a := 5
	// b := 5
	if a < b {
		fmt.Printf("%.2f", float64(a)/float64(b))
		res := float64(a) / float64(b)
		return strconv.FormatFloat(res, 'E', -1, 64)
		// return strconv.ParseFloat(res)
	} else if b < a {
		fmt.Println(a << b)
		res2 := a << b
		return strconv.Itoa(res2)
	} else {
		if a%2 == 0 {
			fmt.Println("EVEN")
			return "EVEN"
		} else {
			fmt.Println("ODD")
			return "ODD"
		}
	}

}
