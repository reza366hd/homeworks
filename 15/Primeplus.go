// package main

// import "fmt"

// func aval(num int) bool {
// 	if num <= 1 {
// 		return false
// 	}
// 	for i := 2; i < num; i++ {
// 		if num%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	var n int
// 	var end int
// 	start := 1
// 	fmt.Print("enter a number to plus all prime numbers smaller than it:")
// 	fmt.Scanln(&end)

// 	for i := start; i < end; i++ {
// 		if aval(i) == true {
// 			n += i
// 		}
// 	}
// 	fmt.Println(n)
// }
