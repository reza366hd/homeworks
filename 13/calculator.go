package main

import (
	"fmt"
	"strconv"
)

func power(first, second int64) int64 {
	natije := first
	var i int64
	for i = 1; i < second; i++ {
		natije *= first
	}
	return int64(natije)
}

func plus(first, second int64) int64 {
	natije := first + second
	return natije
}

func minus(first, second int64) int64 {
	natije := first - second
	return natije
}

func times(first, second int64) int64 {
	natije := first * second
	return natije
}

func remainder(first, second int64) int64 {
	natije := first % second
	return natije
}

func Division(first, second int64) int64 {
	natije := first / second
	return natije
}

func main() {
	var history []int64
	var first string
	var second string
	var oprator string
	var result int64
	var t string
	var request string
	var hisnum int64
	var o int64
	var l string
out:
	for o = 1; o != -1; o++ {

		fmt.Println("Please enter first number")
		fmt.Scanln(&first)
		fmt.Println("Please enter oprator(*,/,-,+,**,%)")
		fmt.Scanln(&oprator)
		fmt.Println("Please enter second number")
		fmt.Scanln(&second)

		intfirst, _ := strconv.ParseInt(first, 10, 0)
		intsecond, _ := strconv.ParseInt(second, 10, 0)

		switch oprator {
		case "/":
			result = Division(intfirst, intsecond)
		case "*":
			result = times(intfirst, intsecond)
		case "-":
			result = minus(intfirst, intsecond)
		case "+":
			result = plus(intfirst, intsecond)
		case "**":
			result = power(intfirst, intsecond)
		case "%":
			result = remainder(intfirst, intsecond)
		default:
			fmt.Println("Please enter a valid operator")
			continue
		}

		fmt.Println("=", result)
		history = append(history, result)

		fmt.Println(result, "saved on history(enter history to see). do you have any request N/Y")
		fmt.Scanln(&request)

		for {
			switch request {

			case "n", "N", "No", "NO":
				break out
			case "Y", "y", "Yes", "YES":
				continue out
			case "history", "History":
				fmt.Println("please enter number of history")
				fmt.Scanln(&hisnum)
				if hisnum > o {
					fmt.Println("Date what number", hisnum, "does not exist!please try again(enter main to open main menu)")
					fmt.Scanln(&l)
					if l == "main" {
						continue out
					} else {
						continue
					}
				}
				hisnum--
				fmt.Println("histoy :", history[hisnum])
				fmt.Println()
				fmt.Println("do you have any request N/Y(enter history to see)")
				fmt.Scanln(&request)
				continue

			default:
				fmt.Println("Please respond according to the menu. enter ok to continue or exit to end")

				fmt.Scanln(&t)
				switch t {
				case "ok", "OK", "Ok":
					fmt.Println("do you have any request N/Y(enter history to see)")
					fmt.Scanln(&request)
					continue out

				case "exit", "Exit", "EXIT":
					break out
				}
			}
		}

	}
}
