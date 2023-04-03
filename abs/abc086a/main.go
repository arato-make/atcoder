package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	var mul int
	mul = a * b

	if mul % 2 == 0 { 
		fmt.Printf("Even\n")
	} else {
		fmt.Printf("Odd\n")
	}
} 