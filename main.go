package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight, userWeight := getUserValues()
	getResult(userHeight, userWeight)
}

func getResult( userHeight, userWeight float64 ) {
	const IMTPower = 2
	IMT := userWeight / math.Pow(userHeight/100, IMTPower)
	
	fmt.Printf("result: %.0f", IMT)
}

func getUserValues () (float64, float64) {
	var userHeight float64
	var userWeight float64
	
	fmt.Print("Enter your height (in sm) ")
	fmt.Scan(&userHeight)

	fmt.Print("Enter your weight (in kg) ")
	fmt.Scan(&userWeight)

	return userHeight, userWeight
}