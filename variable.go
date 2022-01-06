package main

import (
	"fmt"
)

//variables can be declare as package
var (
	name       string  = "ranjeet kumar"
	roll       string  = "20CS06011"
	college    string  = "IIT Bhubaneswar"
	percentage float32 = 80.03
)

func main() {

	//variable declaration
	var i int
	i = 43
	var j = 34
	k := 975 // can be intialize only at compile time
	fmt.Println(i, j, k)

	printVar()
}

func printVar() {
	fmt.Printf("%v, %v, %v, %v\n", name, roll, college, percentage)
	fmt.Printf("%T, %T, %T, %T\n", name, roll, college, percentage)
}
