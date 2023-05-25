package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f", num1, num2, num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Error: divide by zero")
			return
		}
		fmt.Printf("%.2f / %.2f = %.2f", num1, num2, num1/num2)
	default:
		fmt.Println("Invalid operator")
	}
}
