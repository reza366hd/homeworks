package main

import "fmt"

func main() {
	number := 125
	printDigits(number)
}

func printDigits(number int) {
	fmt.Println(number % 10)
	if number >= 10 {
		printDigits(number / 10)

	}

}
