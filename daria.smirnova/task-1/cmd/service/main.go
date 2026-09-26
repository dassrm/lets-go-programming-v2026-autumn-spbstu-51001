package main

import "fmt"

func main() {
	var input_1, input_2 int
	var operator string

	_, err := fmt.Scanln(&input_1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scanln(&input_2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Println(input_1 + input_2)
	case "-":
		fmt.Println(input_1 - input_2)
	case "*":
		fmt.Println(input_1 * input_2)
	case "/":
		if input_2 == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(input_1 / input_2)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
