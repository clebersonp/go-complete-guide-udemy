package main

import (
	"fmt"
	"math"
)

func main() {

	// Constant in go with keyword const
	const inflationRate = 2.5

	// go "null types" (default values)
	// int = 0
	// float - 0.0
	// string = ""
	// bool = false

	//var expectedReturnRate, years = 5.5, 10 // declared and assignment in short way using infer with keyword var
	//expectedReturnRate, years := 5.5, 10 // declared and assignment in short way using infer without keyword var
	//var expectedReturnRate, years float64 = 5.5, 10 // declared and assignment in short way using infer and explicit with keyword var
	//var expectedReturnRate int, years float64 = 5.5, 10 // don't compile, go does not allow to declare multiple variables in line with multiple explicit type, only the last is allowed
	var investmentAmount float64
	expectedReturnRate := 5.5
	years := 10.0

	//
	fmt.Println("Give the investment Amount:")
	fmt.Scan(&investmentAmount) // & this is the pointer to access and change value of variable by its memory address

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

	// to run the program with a module: go run .

}
