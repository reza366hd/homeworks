package main

import "fmt"

func main() {
	score := 2
	for i := 100; i > 10; i-- {
		fmt.Println("score : ", score)
		fmt.Println("i : ", i)
		score *= i
		fmt.Println("score * i = ", score)
	}
}
