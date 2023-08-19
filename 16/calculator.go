// package main

// import (
// 	"fmt"
// )

// func power(first, second int) int {
// 	natije := first
// 	var i int
// 	for i = 1; i < second; i++ {
// 		natije *= first
// 	}
// 	return int(natije)
// }

// func plus(first, second int) int {
// 	natije := first + second
// 	return natije
// }

// func minus(first, second int) int {
// 	natije := first - second
// 	return natije
// }

// func times(first, second int) int {
// 	natije := first * second
// 	return natije
// }

// func remainder(first, second int) int {
// 	natije := first % second
// 	return natije
// }

// func Division(first, second int) int {
// 	natije := first / second
// 	return natije
// }

// func mainmenu() (int, string, int) {
// 	var first int
// 	var second int
// 	var oprator string
// 	fmt.Println("Please enter first number")
// 	fmt.Scanln(&first)
// 	fmt.Println("Please enter oprator(*,/,-,+,**,%)")
// 	fmt.Scanln(&oprator)
// 	fmt.Println("Please enter second number")
// 	fmt.Scanln(&second)

// 	return first, oprator, second

// }

// func operator(first int, oprator string, second int) (result int) {
// 	switch oprator {
// 	case "/":
// 		result = Division(first, second)
// 	case "*":
// 		result = times(first, second)
// 	case "-":
// 		result = minus(first, second)
// 	case "+":
// 		result = plus(first, second)
// 	case "**":
// 		result = power(first, second)
// 	case "%":
// 		result = remainder(first, second)
// 	default:
// 		fmt.Println("Please enter a valid operator")

// 	}
// 	return result
// }

// func request()

// func main() {
// 	var history []int

// 	var result int
// 	var t string
// 	var request string
// 	var hisnum int
// 	var o int
// 	var l string
// out:
// 	for o = 1; o != -1; o++ {

// 		first, oprator, second := mainmenu()

// 		result = operator(first, oprator, second)

// 		fmt.Println("=", result)
// 		history = append(history, result)

// 		fmt.Println(result, "saved on history(enter history to see). do you have any request N/Y")
// 		fmt.Scanln(&request)

// 		for {
// 			switch request {

// 			case "n", "N", "No", "NO":
// 				break out
// 			case "Y", "y", "Yes", "YES":
// 				continue out
// 			case "history", "History":
// 				fmt.Println("please enter number of history")
// 				fmt.Scanln(&hisnum)
// 				if hisnum > o {
// 					fmt.Println("Date what number", hisnum, "does not exist!please try again(enter main to open main menu)")
// 					fmt.Scanln(&l)
// 					if l == "main" {
// 						continue out
// 					} else {
// 						continue
// 					}
// 				}
// 				hisnum--
// 				fmt.Println("histoy :", history[hisnum])
// 				fmt.Println()
// 				fmt.Println("do you have any request N/Y(enter history to see)")
// 				fmt.Scanln(&request)
// 				continue

// 			default:
// 				fmt.Println("Please respond according to the menu. enter ok to continue or exit to end")

// 				fmt.Scanln(&t)
// 				switch t {
// 				case "ok", "OK", "Ok":
// 					fmt.Println("do you have any request N/Y(enter history to see)")
// 					fmt.Scanln(&request)
// 					continue out

// 				case "exit", "Exit", "EXIT":
// 					break out
// 				}
// 			}
// 		}

// 	}
// }
