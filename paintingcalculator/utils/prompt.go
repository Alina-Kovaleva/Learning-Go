package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func UserInput(question string, allowZero bool) float64 {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(question)
		scanner.Scan()
		answer, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Please enter valid value")
		} else if answer < 0 || (answer == 0 && !allowZero) {
			fmt.Println("Please enter positive value")
		} else {
			return answer
		}
	}
}
