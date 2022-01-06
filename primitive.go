package main

import (
	"fmt"
)

func main()  {
	a := 10
	b := 13

	// arithmatic operator
	fmt.Println("Arithmatic operator")
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	//bitwise operator
	fmt.Println("Bitwise operator")
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b) //AND OR operator

	//logical operator
	fmt.Println((a>b) && (b <a))
	
}