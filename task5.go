package main

import "fmt"

func main() {
	var n1, n2 float64
	fmt.Println("Input two numbers: ")
	fmt.Scan(&n1, &n2)

	var op string
	fmt.Println("Input operation: ")
	fmt.Scan(&op)

	switch op {
	case "+":
		fmt.Println(n1, "+", n2, "=", n1+n2)
	case "-":
		fmt.Println(n1, "-", n2, "=", n1-n2)
	case "*":
		fmt.Println(n1, "*", n2, "=", n1*n2)
	case "/":
		{
			if n2 == 0 {
				fmt.Println("Sorry, the division by zero is prohibited!!")
			} else {
				fmt.Println(n1, "/", n2, "=", n1/n2)
			}
		}
	}
}
