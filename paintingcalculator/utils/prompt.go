package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func UserInput(question string) float64 {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(question)
		scanner.Scan()
		answer, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil || answer <= 0 {
			fmt.Println("Please enter valid value")
		} else {
			return answer
		}
	}
}