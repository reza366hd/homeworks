// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var history [10]int
// 	for i := range history {

// 		var first string
// 		var second string
// 		var plus string
// 		fmt.Scanln(&first)
// 		fmt.Scanln(&second)
// 		fmt.Scanln(&plus)

// 		intfirst, _ := strconv.ParseInt(first, 10, 0)
// 		intsecond, _ := strconv.ParseInt(second, 10, 0)
// 		var result int
// 		if plus == "/" {
// 			result = intfirst / intsecond
// 		} else if plus == "*" {
// 			result = intfirst * intsecond
// 		} else if plus == "-" {
// 			result = intfirst - intsecond
// 		} else if plus == "+" {
// 			result = intfirst + intsecond
// 		}
// 		fmt.Println("=", result)
// 		result = history[i]
// 		var request string
// 		fmt.Println(result, "saved on history. do you any request N/Y")
// 		fmt.Scanln(&request)
// 		if request == "N || n || no || NO || No" {
// 			break
// 		} else if request == "Y || y || yes || Yes" {
// 			continue
// 		}
// 	}
// }
