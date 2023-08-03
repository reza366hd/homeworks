package main

import "fmt"

func main() {

	first := 20
	end := 60

	for x := first; x <= end; x++ {

		if x%2 == 0 {
			fmt.Println(x, "is a even number")
		} else {
			fmt.Println(x, "is a odd number")
		}
	}
}
