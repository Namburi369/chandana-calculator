package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return a / b
}

func main() {
	num1, num2 := 11, 5

	sum := Add(num1, num2)
	fmt.Println("Sum:", sum)

	difference := Subtract(num1, num2)
	fmt.Println("Difference:", difference)

	product := Multiply(num1, num2)
	fmt.Println("Product:", product)

	quotient := Divide(num1, num2)
	fmt.Println("Quotient:", quotient)
}
