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

	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)

	// to run the program with a module: go run .

}
