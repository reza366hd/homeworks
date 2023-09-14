// package main

// var rep int

// func main() {
// 	for i := 1; true; i++ {
// 		rep = 1
// 		rep = i * 2

// 		tartibi := tartib(i)

// 	}
// }

// func tartib(arr int) int {
// 	var j int
// 	var t int
// 	var k int
// 	n := len(arr)
// 	for j = 1; j < n; j++ {
// 		t = 1
// 		k = 0
// 		for {
// 			if j-t < 0 || j-k < 0 {
// 				break
// 			} else if arr[j-k] < arr[j-t] {
// 				arr[j-k], arr[j-t] = arr[j-t], arr[j-k]
// 				k++
// 				t++
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	return arr
// }
