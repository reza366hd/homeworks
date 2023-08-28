package nineteen

import "fmt"

func nineteen() {
	var n int
	fmt.Print("enter number: ")
	fmt.Scan(&n)
	if n%3 == 0 && n%5 == 0 {
		fmt.Println("fizzbuzz")
	} else if n%3 == 0 {
		fmt.Println("fizz")
	} else if n%5 == 0 {
		fmt.Println("buzz")
	}
}
