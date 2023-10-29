package main

import (
	"fmt"
	"math"
	"strconv"
)

func ValidateString(s string) bool {
	///Valid Binary Number
	if len(s) > 8 {
		return false
	}
	for _, char := range s {
		if char != '0' && char != '1' {
			return false
		}

	}
	return true
}

func Binary2Decimal(s string) float64 {
	length := len(s)
	sum := 0.0
	for pos, char := range s {
		value, err := strconv.Atoi(string(char))
		if err == nil {
			sum += math.Pow(2, float64(length-(pos+1))) * float64(value)
		}
	}
	return sum
}

func main() {
	//display output in the next line
	fmt.Println("Enter Your Binary Number: ")
	var binary string

	//taking input from user
	fmt.Scanln(&binary)
	if ValidateString(binary) == false {
		fmt.Println("Invalid Input")
	}
	decimal := Binary2Decimal(binary)
	fmt.Println(binary, "binary to decimal is:", decimal)
}
