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
// 	var n []int
// 	var end2 int
// 	fmt.Println("enter number of prime prime number:  ")
// 	fmt.Scanln(&end2)
// 	end := end2 * 3

// 	for i := 1; i <= end; i++ {
// 		if aval(i) {
// 			n = append(n, i)
// 		}
// 	}
// 	fmt.Println(n[end2-1])
// }
