package main

import "fmt"

func main2() {
	fmt.Println("Enter a number : ")

	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2

	fmt.Println(output)
}
