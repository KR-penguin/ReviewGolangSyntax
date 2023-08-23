package myfunctions

import "fmt"

func TwoNumbersPlus(a, b int) (sum int) {

	defer fmt.Println("Done this function.")

	return a + b
}
