package main

import "fmt"

var i int

func main() {
	for i = 1; i > -1; i++ {
		for j := 2; j <= 10; j++ {
			if i%j == 0 && j == 10 {
				fmt.Println(i)
			}
		}

	}
	fmt.Println(i)
}
