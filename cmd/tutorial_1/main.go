package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello, World!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	// with if
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is %v", result)
	// } else {
	// 	fmt.Printf("The result of the integer division is %v with a remainder of %v\n", result, remainder)
	// }
	// with switch statement
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with a remainder of %v\n", result, remainder)
	}

	// switch based on remainder
	switch remainder {
	case 0:
		fmt.Println("Integer division, no remainder")
	default:
		fmt.Printf("Division with remainder %d\n", remainder)
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
