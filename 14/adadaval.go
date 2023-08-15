// package main

// import "fmt"

// func aval(num int) bool {
// 	if num <= 1 {
// 		return false
// 	}
// 	for i := 2; i <= num; i++ {
// 		if num%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	var n []int
// 	start := 1
// 	end := 100

// 	for i := start; i <= end; i++ {
// 		if aval(i) {
// 			n = append(n, i)
// 		}
// 	}
// 	fmt.Println(n[2])
// }
