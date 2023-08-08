package main

import (
	"fmt"
	"strconv"
)

func main() {
	var history []int64
	for {

		var first string
		var second string
		var oprator string
		fmt.Println("Please enter first number")
		fmt.Scanln(&first)
		fmt.Println("Please enter oprator(*,/,-,+,**,%)")
		fmt.Scanln(&oprator)
		fmt.Println("Please enter second number")
		fmt.Scanln(&second)

		intfirst, _ := strconv.ParseInt(first, 10, 0)
		intsecond, _ := strconv.ParseInt(second, 10, 0)
		var result int64

		switch oprator {
		case "/":
			result = intfirst / intsecond
		case "*":
			result = intfirst * intsecond
		case "-":
			result = intfirst - intsecond
		case "+":
			result = intfirst + intsecond
		case "**":
			natije := intfirst
			var i int64
			for i = 1; i < intsecond; i++ {
				natije *= intfirst
			}
			result = natije
		case "%":
			result = intfirst % intsecond
		default:
			fmt.Println("Please enter a valid operator")
			continue
		}

		fmt.Println("=", result)
		history = append(history, result)
		var request string
		fmt.Println(result, "saved on history(enter history to see). do you have any request N/Y")
		fmt.Scanln(&request)

		switch request {
		case "n", "N", "No", "NO":
			break
		case "Y", "y", "Yes", "YES":
			continue
		case "history", "History":
			fmt.Printf("history :%d \n", history)
		}

	}
}
