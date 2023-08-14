package main

import "fmt"

func main() {
	var size int
	fmt.Print("Enter the size of the diamond: ")
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		for j := 0; j < size-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := size - 2; i >= 0; i-- {
		for j := 0; j < size-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
