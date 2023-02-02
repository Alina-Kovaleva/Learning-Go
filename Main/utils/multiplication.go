package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Calculator() {

	num1 := PromptFloat("Enter first number: ")
	num2 := PromptFloat("Enter second number: ")
	operator, f := PromptOperator("Enter operator (+, -, *, /): ")

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, f(num1, num2))

}

func PromptFloat(question string) float64 {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(question)
		scanner.Scan()
		result, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Please enter valid value")
		} else {
			return result
		}
	}
}

func PromptOperator(question string) (string, func(float64, float64) float64) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(question)
		scanner.Scan()
		operator := scanner.Text()

		if !Contains([]string{"+", "-", "*", "/"}, operator) {
			fmt.Println("Invalid operator")
			continue
		}
		return operator, func(a float64, b float64) float64 {
			switch operator {
			case "+":
				return a + b
			case "-":
				return a - b
			case "*":
				return a * b
			default:
				return a / b
			}
		}
	}
}

func Contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if needle == s {
			return true
		}
	}
	return false
}
