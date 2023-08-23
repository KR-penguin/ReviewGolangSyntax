package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello, World!") // Print
	var a int32 = 0
	var b float32 = 3.14

	fmt.Printf("\n\na + b = %.2f", (float32(a) + b)) // variable print format

	var c int32 = 0

	fmt.Printf("\n\n")
	fmt.Scan(&c) // input
	fmt.Printf("c = %d", c)

	var FirstNumber int = 0
	var SecondNumber int = 0

	fmt.Print("\n\nInput first number : ")
	fmt.Scan(&FirstNumber)
	fmt.Print("\nInput second number : ")
	fmt.Scan(&SecondNumber)

	//	var Sum int = myfunctions.TwoNumbersPlus(FirstNumber, SecondNumber)

	//	fmt.Println("\n\n%d + %d = %d", FirstNumber, SecondNumber, Sum)
}
