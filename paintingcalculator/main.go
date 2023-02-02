package main

import (
	"fmt"
	"github.com/Alina-Kovaleva/paintingcalculator/utils"
)

func main() {

	wallLength := utils.UserInput("Enter the combined length of all walls (in meters): ")
	wallHeight := utils.UserInput("Enter the height of the wall (in meters): ")
	windowsArea := utils.UserInput("Enter the total area of windows (in square meters): ")
	doorsArea := utils.UserInput("Enter the total area of doors (in square meters): ")
	consumption := utils.UserInput("Enter the paint consumption indicated on the paint can (in square meters per liter): ")
	literPerCan := utils.UserInput("Enter how many liters in one can of paint: ")
	coats := utils.UserInput("How many of coats of paint you want to apply to the area?: ")
	pricePerCan := utils.UserInput("Enter the price for one can of paint: ")

	totalArea := wallLength * wallHeight
	paintedArea := totalArea - windowsArea - doorsArea

	coverinPerCan := consumption * literPerCan

	totalCans := (paintedArea / coverinPerCan) * coats

	totalCost := totalCans * pricePerCan

	fmt.Printf("You will need %.2f cans of paint and it will cost %.2f", totalCans, totalCost)

}


