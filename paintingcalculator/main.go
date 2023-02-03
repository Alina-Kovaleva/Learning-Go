package main

import (
	"fmt"
	"math"

	"github.com/Alina-Kovaleva/paintingcalculator/utils"
)

// anything
func main() {
	for {
		wallLength := utils.UserInput("Enter the combined length of all walls (in meters): ", false)
		wallHeight := utils.UserInput("Enter the height of the wall (in meters): ", false)
		windowsArea := utils.UserInput("Enter the total area of windows (in square meters): ", true)
		doorsArea := utils.UserInput("Enter the total area of doors (in square meters): ", false)

		totalArea := wallLength * wallHeight
		paintedArea := totalArea - windowsArea - doorsArea

		if paintedArea <= 0 {
			fmt.Println("Oops, it looks like there's an error. The combined area of your doors and windows appears to be larger than the total area of the room.\nTry one more time!")
			continue
		} else {
			consumption := utils.UserInput("Enter the paint consumption indicated on the paint can (in square meters per liter): ", false)
			literPerCan := utils.UserInput("Enter how many liters in one can of paint: ", false)
			coats := utils.UserInput("How many of coats of paint you want to apply to the area?: ", false)
			pricePerCan := utils.UserInput("Enter the price for one can of paint: ", false)
			coverinPerCan := consumption * literPerCan
			totalCans := math.Ceil((paintedArea / coverinPerCan) * coats)
			totalCost := totalCans * pricePerCan

			fmt.Printf("You will need %.0f cans of paint and it will cost %.2f", totalCans, totalCost)
			break
		}
	}
}
