package main

import (
	"errors"
	"fmt"
)

func main() {
	var intNum int
	fmt.Println(intNum)

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	myVar := "text"
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	fmt.Println("Hello world!")
	printMe("Yeeeee!")
}

func printMe(value string) {
	fmt.Println(value)

	const numerator, denominator int = 10, 0

	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Print(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with %v", result, remainder)
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = denominator / numerator
	var remainder int = denominator % numerator
	return result, remainder, err
}
