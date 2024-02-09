package main

import (
	"fmt"
	"math"
)

func main() {

	// go "null types" (default values)
	// int = 0
	// float - 0.0
	// string = ""
	// bool = false

	//var expectedReturnRate, years = 5.5, 10 // declared and assignment in short way using infer with keyword var
	//expectedReturnRate, years := 5.5, 10 // declared and assignment in short way using infer without keyword var
	//var expectedReturnRate, years float64 = 5.5, 10 // declared and assignment in short way using infer and explicit with keyword var
	//var expectedReturnRate int, years float64 = 5.5, 10 // don't compile, go does not allow to declare multiple variables in line with multiple explicit type, only the last is allowed
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	years := 10.0

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)

	// to run the program with a module: go run .

}
