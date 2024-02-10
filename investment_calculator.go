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
	var expectedReturnRate float64 // when declare a variable without assignment, it needs to use keyword var and explicitly type
	var years float64

	// value for investment amount
	fmt.Print("Give the investment Amount: ")
	fmt.Scan(&investmentAmount) // & this is the pointer to access and change value of variable by its memory address

	// value for expected return rate
	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// value for years
	fmt.Print("Years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	//fmt.Println(futureValue)
	//fmt.Println(futureRealValue)
	// formatting output doc: https://pkg.go.dev/fmt@go1.22.0
	//fmt.Printf("Future Value: %.2f\nFuture Real Value: %.2f\n", futureValue, futureRealValue)
	// to run the program with a module: go run .

	//formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	//formattedFRV := fmt.Sprintf("Future Value (adjusted for Inflation): %.2f\n", futureRealValue)
	//
	//fmt.Print(formattedFV, formattedFRV)
	// using backtick `some texts here`. The break lines are visual, we don't want to use scape \n to break lines anymore
	fmt.Printf(`Future Value: %.2f
Future Real Value: %.2f
`, futureValue, futureRealValue)
}
